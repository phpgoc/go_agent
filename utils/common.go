package utils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"
)

var (
	defaultLogFileNameAtTempDir = "xxx/agent.log"
	useFileToLog                = flag.Bool("use_file", false, fmt.Sprintf("Use file to log, if not set is Std Out,if set default is %s", defaultLogFileNameAtTempDir))
	logFileNameAtTempDir        = flag.String("log_file", defaultLogFileNameAtTempDir, "Log file name,, Must use use_file to take effect,it is about relative paths to temporary folders")
)
var logFile *os.File

// default std out
var writer = os.Stdout

func Init() (err error) {
	osInitBefore()
	defer osInitAfter()
	//use LogFileName get path
	if *useFileToLog {
		fullFileName := filepath.Join(os.TempDir(), *logFileNameAtTempDir)
		dirName := filepath.Dir(fullFileName)
		//make dir
		//q: windows 为什么没权利创建文件夹
		err = os.MkdirAll(dirName, 0755)
		if err != nil {
			println(1)
			return err
		}
		println(fullFileName)
		logFile, err = os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		// 适当的条件下设置writer = logFile,默认是os.Stdout
		writer = logFile
	}

	return err
}

func writeLogFile(log string, level string) {
	// log to file
	// write now
	_, err := writer.WriteString(fmt.Sprintf("%s, %s, %s\n", level,
		time.Now().Format("2006-01-02 15:04:05"), log))
	if err != nil {
		println(err.Error())
	}
}

func LogInfo(log string) {
	writeLogFile(log, "INFO")
}

func LogWarn(log string) {
	writeLogFile(log, "WARN")
}

func LogError(log string) {
	writeLogFile(log, "ERROR")
}

func ExtractFileStat(file string) (size uint64, accessTime, modifyTime string) {
	fi, err := os.Stat(file)
	if err != nil {
		return
	}
	size = uint64(fi.Size())
	accessTime = fi.ModTime().String()
	modifyTime = fi.ModTime().String()
	return
}

func FormatTimeByTimestamp(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}
func FormatTime(timestamp time.Time) string {
	return timestamp.Format("2006-01-02 15:04:05")
}

func ReadFile(fileName string) (content string, err error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	content = string(bytes)
	return
}
func FormatDuration(uptime time.Duration) string {
	var result string
	for _, unit := range []struct {
		duration time.Duration
		unit     string
	}{
		{time.Hour * 24 * 365, "year"},
		{time.Hour * 24 * 30, "month"},
		{time.Hour * 24, "day"},
		{time.Hour, "hour"},
		{time.Minute, "minute"},
		{time.Second, "second"},
	} {

		if uptime >= unit.duration {
			count := uptime / unit.duration
			result += fmt.Sprintf("%d %s ", count, unit.unit)
			uptime -= count * unit.duration
		}
	}
	return result
}

func SplitStringAndGetIndexSafely(s, sep string, index int) string {
	splitString := strings.Split(s, sep)
	if len(splitString) <= index {
		return ""
	}
	return splitString[index]
}

func FindMatchedFiles(matchPattern string) (files []string, err error) {

	//split matchPattern
	splitString := strings.Split(matchPattern, "/")
	matchPattern = splitString[len(splitString)-1]
	pathName := strings.Join(splitString[:len(splitString)-1], "/")
	dirEntries, err := os.ReadDir(pathName)
	if err != nil {
		return
	}
	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}
		if matched, _ := regexp.MatchString(matchPattern, entry.Name()); matched {
			files = append(files, filepath.Join(pathName, entry.Name()))
		}
	}
	return
}

func findFromInAndOut(key string, in map[string]string, out map[string]string) string {
	//优先out
	if v, ok := out[key]; ok {
		return v
	}
	if v, ok := in[key]; ok {
		return v
	}
	return ""
}

// 简单字符串判断，没有能力条件判断
func InterpretSourceExportToGoMap(content string, in map[string]string) (out map[string]string) {
	out = make(map[string]string)

	//for key, value := range in {
	//	re := regexp.MustCompile(key)
	//	content = re.ReplaceAllString(content, value)
	//}
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.Trim(line, "\t ")

		//如果是注释跳过
		if strings.HasPrefix(line, "#") {
			continue
		}
		//必须开始于export
		if !strings.HasPrefix(line, "export") {
			continue
		}
		if strings.Contains(line, "=") {
			split := strings.Split(line, "=")
			key := strings.Trim(split[0][6:], "\t ")
			forOut := strings.Trim(split[1], " \"")
			var trueOut string
			//不一定穷举，发现bug，如果还有别的字符也可以中断变量输入再继续加入
			var canStopVariableChar = []int32{' ', '\t', '\n', ';', '/', '\\', ',', '('}
			if strings.Contains(forOut, "$") {
				var v string
				var hasQuotationMark bool = false
				var hasBrace bool = false
				var started = false
				// find every variable

				for _, c := range forOut {
					if c == '$' {
						started = true
					} else if v == "" && c == '"' {
						hasQuotationMark = true
					} else if v == "" && c == '{' {
						hasBrace = true
					} else if (hasQuotationMark && c == '"') ||
						(hasBrace && c == '}') ||
						(!hasQuotationMark && !hasBrace && slices.Contains(canStopVariableChar, c)) {
						//结束符号
						//search in and out replace it in for
						//都找不到也没问题，替换空字符串
						trueOut += findFromInAndOut(v, in, out)
						//非“，}的字符直接加入
						if !hasQuotationMark && !hasBrace {
							trueOut += string(c)
						}
						started = false
						hasQuotationMark = false
						hasBrace = false
						v = ""
					} else if started {
						v += string(c)
					} else {
						trueOut += string(c)
					}
				}
				if v != "" {
					trueOut += findFromInAndOut(v, in, out)
				}
			} else {
				trueOut = forOut
			}
			out[key] = trueOut
		}
	}
	return
}

func ReplaceStrUseEnvMapStrictWithBrace(content string, envMap map[string]string) string {
	re := regexp.MustCompile(`\${(\w+)}`)

	matches := re.FindAll([]byte(content), -1)
	for _, match := range matches {
		key := string(match[2 : len(match)-1])
		if value, ok := envMap[key]; ok {
			content = strings.ReplaceAll(content, string(match), value)
		} else {
			content = strings.ReplaceAll(content, string(match), "")
		}
	}

	return content
}

func FormatBytes(total uint64) string {
	for _, unit := range []struct {
		unit string
		size uint64
	}{
		{"TB", 1024 * 1024 * 1024 * 1024},
		{"GB", 1024 * 1024 * 1024},
		{"MB", 1024 * 1024},
		{"KB", 1024},
	} {
		if total >= unit.size {
			return fmt.Sprintf("%.2f %s", float64(total)/float64(unit.size), unit.unit)
		}
	}
	return fmt.Sprintf("%d B", total)
}
