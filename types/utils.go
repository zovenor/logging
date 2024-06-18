package types

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

var lastCheckOldFiles time.Time

func getOptsFromInterface(args ...any) ([]any, []func(*Message)) {
	filteredArgs := make([]any, 0)
	opts := make([]func(*Message), 0)
	for _, arg := range args {
		if opt, ok := arg.(func(*Message)); ok {
			opts = append(opts, opt)
		} else {
			filteredArgs = append(filteredArgs, arg)
		}
	}
	return filteredArgs, opts
}

func getBinPath() (string, error) {
	projectPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return projectPath, nil
}

func getFormattedValue(timeNow time.Time, values ...any) string {
	return fmt.Sprintf("[%v] %v", timeNow.Format(timeFormatter), fmt.Sprint(values...))
}

func setOpts(message *Message, opts ...func(*Message)) {
	for _, opt := range opts {
		opt(message)
	}
	if message.sendToChan {
		loggerConfigs.sendToChan(message)
	}
}

func checkOldLogs() {
    if time.Since(lastCheckOldFiles) < 1*time.Hour {
        fmt.Println("return old logs")
        return
    }
    fmt.Println("clear old logs")
    lastCheckOldFiles = time.Now()

	now := time.Now()
	files, err := ioutil.ReadDir(loggerConfigs.logsDirPath)
	if err != nil {
        return
	}
	for _, file := range files {
		filePath := filepath.Join(loggerConfigs.logsDirPath, file.Name())

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			continue
		}

		if now.Sub(fileInfo.ModTime()) > 24*time.Hour {
			err := os.Remove(filePath)
			if err != nil {
                continue
			}	
        }
	}
}

func saveLogs(message *Message, prefix string) error {
    checkOldLogs()
	if message.msgType == FatalMessageType {
		psList := make([]string, 0)
		skip := 4
		if message.action == PrintAndSaveAction {
			skip = 5
		}
		for i := range 100 {
			_, filename, line, _ := runtime.Caller(i + skip)
			if filename == "" {
				break
			}
			psList = append(psList, fmt.Sprintf("%v:%v", filename, line))
		}
		psList = psList[:len(psList)-2]
		message.value = "	" + message.value
		for _, psElem := range psList {
			message.value = fmt.Sprintf("	> %v\n%v", psElem[len(loggerConfigs.binPath):], message.value)
		}
		message.value = "\n" + message.value
	}
	fileName := fmt.Sprintf("%v.log", message.time.Format(timeFileFormatter))
	file, err := openLogFileOrCreate(loggerConfigs.logsDirPath, fileName)
	if err != nil {
		defer file.Close()
		return err
	}
	defer file.Close()
	logger := log.New(file, prefix, log.LstdFlags)
	logger.Println(message.value)
	return nil
}

func openLogFileOrCreate(dirPath string, fileName string) (*os.File, error) {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		os.Mkdir(dirPath, 0755)
	}
	filePath := path.Join(dirPath, fileName)
	logFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logFile, err = os.Create(filePath)
		if err != nil {
			defer logFile.Close()
			return nil, err
		}
	}
	return logFile, nil
}
