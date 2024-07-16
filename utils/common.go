package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"
	"time"
)

func writeLogFile(level string, log string, callerLevel int) {
	// log to file
	// write now
	_, filename, line, _ := runtime.Caller(callerLevel)
	_, err := writer.WriteString(fmt.Sprintf("%s, %s\tfile:///%s:%d\t %s\n", level,
		time.Now().Format("2006-01-02 15:04:05"),
		filename, line,
		log))
	if err != nil {
		println(err.Error())
	}
}

// 如何接受变长
func LogInfo(log string) {
	writeLogFile("INFO", log, 2)
}

func LogWarn(log string) {
	writeLogFile("WARN", log, 2)
}

func LogError(log string) {
	writeLogFile("ERROR", log, 2)
}

func LogErrorWithCallerLevel(log string, callerLevel int) {
	writeLogFile("ERROR", log, callerLevel)
}

// GetFirstAndLogError 有错误打印日志并拿默认值继续
// 这个两个方法不适合有错return,continue,break的情况
// 完全不会有错误的情况当然也不需要使用这两个方法
func GetFirstAndLogError[T any](fn func() (T, error), defaultValue ...T) T {
	result, err := fn()
	if err != nil {
		//LogError(err.Error())
		//这里不能使用LogError，因为caller等级不同
		LogErrorWithCallerLevel(err.Error(), 3)
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return result
}

// GetFirstWithoutLogError 有错误不打印日志(已知大概率会出现错误的情况)
// 这个方法使用场景更少,大多数的情况可以用_忽略错误的方式,只有不想提前赋值,想写在一行的时候才会用到
func GetFirstWithoutLogError[T any](fn func() (T, error), defaultValue ...T) T {
	result, _ := fn()
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return result
}

func ExtractFileStat(file string) (size uint64, accessTime, modifyTime string) {
	fi, err := os.Stat(file)
	if err != nil {
		return
	}
	size = uint64(fi.Size())
	accessTime = FormatTime(fi.ModTime())
	modifyTime = FormatTime(fi.ModTime())
	return
}

func FormatTimeByTimestamp(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func FormatTime(timestamp time.Time) string {
	return timestamp.Format("2006-01-02 15:04:05")
}
func FormatTimeForFileName(timestamp time.Time) string {
	return timestamp.Format("2006-01-02_15-04-05")
}

func ReadFile(fileName string) (content string, err error) {
	readBytes, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	content = string(readBytes)
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

func SplitStringAndGetIndexSafelyBySelfDefineSeq(s, sep string, index int) string {
	splitString := strings.Split(s, sep)
	if len(splitString) <= index {
		return ""
	}
	return splitString[index]
}

func SplitStringAndGetIndexSafelyByDefault(s string, index int) string {

	re := regexp.MustCompile(`\s+`)
	splitString := re.Split(s, -1)
	if len(splitString) <= index {
		return ""
	}
	return splitString[index]
}

func FindMatchedFiles(matchPattern string) (files []string, err error) {
	//split matchPattern
	splitString := strings.Split(matchPattern, string(os.PathSeparator))
	matchPattern = splitString[len(splitString)-1]
	matchPattern = strings.ReplaceAll(matchPattern, "*", ".*")
	pathName := strings.Join(splitString[:len(splitString)-1], string(os.PathSeparator))
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

// InterpretSourceExportToGoMap 简单字符串判断，没有能力条件判断
// linux 独有方法
func InterpretSourceExportToGoMap(content string, in map[string]string) (out map[string]string) {
	//深拷贝,浪费了点空间,没有污染原始数据
	out = DeepCopyMapGeneric(in)

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.Trim(line, "\t ")

		//如果是注释跳过
		if strings.HasPrefix(line, "#") {
			continue
		}
		//必须开始于export
		if !strings.HasPrefix(line, "export") {
			//去掉line开头的export和空格
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
	re := regexp.MustCompile(`\${\w+}`)

	matches := re.FindAll([]byte(content), -1)
	for _, match := range matches {
		key := string(match[2 : len(match)-1])
		if value, ok := envMap[key]; ok {
			content = strings.ReplaceAll(content, string(match), value)
		} else {
			content = strings.ReplaceAll(content, string(match), "")
		}
	}

	re = regexp.MustCompile(`"\$\w+"`)
	matches = re.FindAll([]byte(content), -1)
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

// GBKToUTF8 convert gbk to utf8 成熟的函数不做单元测试
func GBKToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}

	return d, nil
}

func FileExists(path string) bool {
	cleanedPath := filepath.Clean(path)
	fs, err := os.Stat(cleanedPath)
	if err != nil {
		return false
	}
	return !fs.IsDir()
}

//func DirExists(path string) bool {
//	fs, err := os.Stat(path)
//	if err != nil {
//		return false
//	}
//	return fs.IsDir()
//}

func DeepCopyMapGeneric[K comparable, V any](originalMap map[K]V) map[K]V {
	newMap := make(map[K]V)
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}
