package logging

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/zovenor/logging/prettyPrints"
)

const formatterTime string = "15:04:05"
const formatterTimeForFile string = "2006-01-02-15h"
var lastCheckOldFiles time.Time

func Warning(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}
	timeNow := time.Now()
	value := getFormattedValue(timeNow, err.Error())
	prettyPrints.Warning(value)
	saveLogs(timeNow, "[warning]", err.Error())
}

func WarningPrint(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}
	timeNow := time.Now()
	value := getFormattedValue(timeNow, err.Error())
	prettyPrints.Warning(value)
}

func WarningSave(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}
	timeNow := time.Now()
	saveLogs(timeNow, "[warning]", err.Error())
}

func Warningf(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Warning(value)
	saveLogs(timeNow, "[warning]", valueBase)
}

func WarningfPrint(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Warning(value)
}

func WarningfSave(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	saveLogs(timeNow, "[warning]", valueBase)
}

func Fatal(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}
	timeNow := time.Now()
	v := err.Error()
	packagePath, err := getBinPath()
	if err != nil {
		log.Fatal(err)
	}
	_, filename, line, _ := runtime.Caller(1)
	v = fmt.Sprintf("%v:%v: %v", filename[len(packagePath):], line, v)
	value := getFormattedValue(timeNow, v)
	prettyPrints.Fatal(value)
	saveLogs(timeNow, "[fatal]", v)
}

func FatalPrint(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}
	timeNow := time.Now()
	value := getFormattedValue(timeNow, err.Error())
	prettyPrints.Fatal(value)
}

func FatalSave(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}
	timeNow := time.Now()
	saveLogs(timeNow, "[fatal]", err.Error())
}

func Fatalf(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Fatal(value)
	saveLogs(timeNow, "[fatal]", valueBase)
}

func FatalfPrint(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Fatal(value)
}

func FatalfSave(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	saveLogs(timeNow, "[fatal]", valueBase)
}

func Info(values ...any) {
	timeNow := time.Now()
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Info(valueString)
	saveLogs(timeNow, "[info]", values...)
}

func InfoPrint(values ...any) {
	timeNow := time.Now()
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Info(valueString)
}

func InfoSave(values ...any) {
	timeNow := time.Now()
	saveLogs(timeNow, "[info]", values...)
}

func Infof(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Info(valueString)
	saveLogs(timeNow, "[info]", valueBase)
}

func InfofPrint(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Info(valueString)
}

func InfofSave(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	saveLogs(timeNow, "[info]", valueBase)
}

func Success(values ...any) {
	timeNow := time.Now()
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Success(valueString)
	saveLogs(timeNow, "[success]", values...)
}

func SuccessPrint(values ...any) {
	timeNow := time.Now()
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Success(valueString)
}

func SuccessSave(values ...any) {
	timeNow := time.Now()
	saveLogs(timeNow, "[success]", values...)
}

func Successf(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Success(valueString)
	saveLogs(timeNow, "[success]", valueBase)
}

func SuccessfPrint(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Success(valueString)
}

func SuccessfSave(format string, args ...interface{}) {
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	saveLogs(timeNow, "[success]", valueBase)
}

// Without saving
func Println(values ...any) {
	timeNow := time.Now()
	value := getFormattedValue(timeNow, values...)
	prettyPrints.Print(value)
}

func Printf(format string, args ...interface{}) {
	timeNow := time.Now()
	value := fmt.Sprintf(format, args...)
	value = getFormattedValue(timeNow, value)
	prettyPrints.Print(value)
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

func getBinPath() (string, error) {
	projectPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return projectPath, nil
}

func checkOldLogs(dir string) {
    if time.Since(lastCheckOldFiles) < 1*time.Hour {
        return
    }
    lastCheckOldFiles = time.Now()

	now := time.Now()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
        return
	}
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())

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

func saveLogs(timeNow time.Time, prefix string, values ...any) error {
	projectPath, err := getBinPath()
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%v.log", timeNow.Format(formatterTimeForFile))
	logsDir := path.Join(projectPath, "logs")
    checkOldLogs(logsDir)
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
