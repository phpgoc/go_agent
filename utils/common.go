package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"
)

var logFile *os.File

func Init() (err error) {
	//use LogFileName get path
	dirName := path.Dir(logFileName)
	//make dir
	err = os.MkdirAll(dirName, 0755)
	if err != nil {
		return err
	}

	logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	osInit()

	return err
}

func writeLogFile(log string, level string) (err error) {
	// log to file
	_, err = logFile.WriteString(fmt.Sprintf("%s %s\n", level, log))
	return err
}

func LogInfo(log string) (err error) {
	return writeLogFile(log, "INFO")
}

func LogWarn(log string) (err error) {
	return writeLogFile(log, "WARN")
}

func LogError(log string) (err error) {
	return writeLogFile(log, "ERROR")
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

func FormatTime(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func ReadFile(fileName string) (content string, err error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	content = string(bytes)
	return
}

func SplitStringAndGetIndexSafe(s, sep string, index int) string {
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
				var hasDoubleQuote bool = false
				var hasBigParentheses bool = false
				var started = false
				// find every variable

				for _, c := range forOut {
					if c == '$' {
						started = true
					} else if v == "" && c == '"' {
						hasDoubleQuote = true
					} else if v == "" && c == '{' {
						hasBigParentheses = true
					} else if (hasDoubleQuote && c == '"') ||
						(hasBigParentheses && c == '}') ||
						(!hasDoubleQuote && !hasBigParentheses && slices.Contains(canStopVariableChar, c)) {
						//结束符号
						//search in and out replace it in for
						//都找不到也没问题，替换空字符串
						trueOut += findFromInAndOut(v, in, out)
						//非“，}的字符直接加入
						if !hasDoubleQuote && !hasBigParentheses {
							trueOut += string(c)
						}
						started = false
						hasDoubleQuote = false
						hasBigParentheses = false
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
