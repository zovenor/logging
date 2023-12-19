package logging

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/zovenor/logging/prettyPrints"
)

const formatterTime string = "2006/01/02 15:04:05"
const formatterTimeForFile string = "2006-01-02-15h"

// With saving
func Warning(err error) {
	timeNow := time.Now()
	value := getFormattedValue(timeNow, err.Error())
	prettyPrints.Warning(value)
	saveLogs(timeNow, "[warning]", err.Error())
}

func Warningf(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Warning(value)
	saveLogs(timeNow, "[warning]", valueBase)
}

func Fatal(err error) {
	timeNow := time.Now()
	value := getFormattedValue(timeNow, err.Error())
	prettyPrints.Fatal(value)
	saveLogs(timeNow, "[fatal]", err.Error())
}

func Fatalf(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Fatal(value)
	saveLogs(timeNow, "[fatal]", valueBase)
}

func Info(values ...any) {
	timeNow := time.Now()
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Info(valueString)
	saveLogs(timeNow, "[info]", values...)
}

func Infof(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Info(valueString)
	saveLogs(timeNow, "[info]", valueBase)
}

func Success(values ...any) {
	timeNow := time.Now()
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Success(valueString)
	saveLogs(timeNow, "[success]", values...)
}

func Successf(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Success(valueString)
	saveLogs(timeNow, "[success]", valueBase)
}

// Without saving
func Println(values ...any) {
	timeNow := time.Now()
	value := getFormattedValue(timeNow, values...)
	fmt.Print(value)
}

func Printf(format string, args ...interface{}) {
	timeNow := time.Now()
	value := fmt.Sprintf(format, args...)
	value = getFormattedValue(timeNow, value)
	fmt.Print(value)
}

// Additional functions
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

func saveLogs(timeNow time.Time, prefix string, values ...any) error {
	projectPath, err := os.Getwd()
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%v.log", timeNow.Format(formatterTimeForFile))
	logsDir := path.Join(projectPath, "logs")
	file, err := openLogFileOrCreate(logsDir, fileName)
	if err != nil {
		defer file.Close()
		return err
	}
	defer file.Close()
	logger := log.New(file, prefix, log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	logger.Println(values...)
	return nil
}

func getFormattedValue(timeNow time.Time, values ...any) string {
	return fmt.Sprintf("[%v] %v", timeNow.Format(formatterTime), fmt.Sprintln(values...))
}
