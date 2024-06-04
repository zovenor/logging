package types

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

func getOptsFromInterface(args ...any) (filteredArgs []any, opts []func(*Message)) {
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

func saveLogs(message *Message, prefix string) error {
	fileName := fmt.Sprintf("%v.log", message.time.Format(timeFileFormatter))
	file, err := openLogFileOrCreate(loggerConfigs.logsDirPath, fileName)
	if err != nil {
		defer file.Close()
		return err
	}
	defer file.Close()
	logger := log.New(file, prefix, log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
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
