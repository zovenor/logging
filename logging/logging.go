package logging

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/zovenor/logging/prettyPrints"
)

const formatterTime string = "15:04:05"
const formatterTimeForFile string = "2006-01-02-15h"

func Warning(e error, opts ...func(*LogConfigs)) error {
	if e == nil {
		e = fmt.Errorf("")
	}

	_, err := UpdateChannel(e.Error(), &LogConfigs{msgType: WarningMessageType}, opts...)
	if err != nil {
		return err
	}
	timeNow := time.Now()
	value := getFormattedValue(timeNow, e.Error())
	prettyPrints.Warning(value)
	err = saveLogs(timeNow, "[warning]", e.Error())
	if err != nil {
		return err
	}
	return nil
}

func WarningPrint(e error, opts ...func(*LogConfigs)) error {
	if e == nil {
		e = fmt.Errorf("")
	}
	_, err := UpdateChannel(e.Error(), &LogConfigs{msgType: WarningMessageType}, opts...)
	if err != nil {
		return err
	}
	timeNow := time.Now()
	value := getFormattedValue(timeNow, e.Error())
	prettyPrints.Warning(value)
	return nil
}

func WarningSave(e error, opts ...func(*LogConfigs)) error {
	if e == nil {
		e = fmt.Errorf("")
	}
	_, err := UpdateChannel(e.Error(), &LogConfigs{msgType: WarningMessageType}, opts...)
	if err != nil {
		return err
	}
	timeNow := time.Now()
	err = saveLogs(timeNow, "[warning]", e.Error())
	if err != nil {
		return err
	}
	return nil
}

func Warningf(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: WarningMessageType}, opts...)
	if err != nil {
		return err
	}
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Warning(value)
	err = saveLogs(timeNow, "[warning]", valueBase)
	if err != nil {
		return err
	}
	return nil
}

func WarningfPrint(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: WarningMessageType}, opts...)
	if err != nil {
		return err
	}
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Warning(value)
	return nil
}

func WarningfSave(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: WarningMessageType}, opts...)
	if err != nil {
		return err
	}
	err = saveLogs(timeNow, "[warning]", valueBase)
	if err != nil {
		return err
	}
	return nil
}

func Fatal(e error, opts ...func(*LogConfigs)) error {
	if e == nil {
		e = fmt.Errorf("")
	}
	_, err := UpdateChannel(e.Error(), &LogConfigs{msgType: FatalMessageType}, opts...)
	if err != nil {
		return err
	}
	timeNow := time.Now()
	v := e.Error()
	packagePath, err := getBinPath()
	if err != nil {
		return err
	}
	_, filename, line, _ := runtime.Caller(1)
	v = fmt.Sprintf("%v:%v: %v", filename[len(packagePath):], line, v)
	value := getFormattedValue(timeNow, v)
	prettyPrints.Fatal(value)
	err = saveLogs(timeNow, "[fatal]", v)
	if err != nil {
		return err
	}
	return nil
}

func FatalPrint(e error, opts ...func(*LogConfigs)) error {
	if e == nil {
		e = fmt.Errorf("")
	}
	_, err := UpdateChannel(e.Error(), &LogConfigs{msgType: FatalMessageType}, opts...)
	if err != nil {
		return err
	}
	timeNow := time.Now()
	value := getFormattedValue(timeNow, e.Error())
	prettyPrints.Fatal(value)
	return nil
}

func FatalSave(e error, opts ...func(*LogConfigs)) error {
	if e == nil {
		e = fmt.Errorf("")
	}
	_, err := UpdateChannel(e.Error(), &LogConfigs{msgType: FatalMessageType}, opts...)
	if err != nil {
		return err
	}
	timeNow := time.Now()
	err = saveLogs(timeNow, "[fatal]", e.Error())
	if err != nil {
		return err
	}
	return nil
}

func Fatalf(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: FatalMessageType}, opts...)
	if err != nil {
		return err
	}
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Fatal(value)
	err = saveLogs(timeNow, "[fatal]", valueBase)
	if err != nil {
		return err
	}
	return nil
}

func FatalfPrint(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: FatalMessageType}, opts...)
	if err != nil {
		return err
	}
	value := getFormattedValue(timeNow, valueBase)
	prettyPrints.Fatal(value)
	return nil
}

func FatalfSave(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: FatalMessageType}, opts...)
	if err != nil {
		return err
	}
	err = saveLogs(timeNow, "[fatal]", valueBase)
	if err != nil {
		return err
	}
	return nil
}

func Info(values ...any) error {
	values, opts := getOptsFromInterface(values)
	timeNow := time.Now()
	_, err := UpdateChannel(fmt.Sprint(values), &LogConfigs{msgType: InfoMessageType}, opts...)
	if err != nil {
		return err
	}
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Info(valueString)
	err = saveLogs(timeNow, "[info]", values...)
	if err != nil {
		return err
	}
	return nil
}

func InfoPrint(values ...any) error {
	values, opts := getOptsFromInterface(values)
	timeNow := time.Now()
	_, err := UpdateChannel(fmt.Sprint(values), &LogConfigs{msgType: InfoMessageType}, opts...)
	if err != nil {
		return err
	}
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Info(valueString)
	return nil
}

func InfoSave(values ...any) error {
	values, opts := getOptsFromInterface(values)
	timeNow := time.Now()
	_, err := UpdateChannel(fmt.Sprint(values), &LogConfigs{msgType: InfoMessageType}, opts...)
	if err != nil {
		return err
	}
	err = saveLogs(timeNow, "[info]", values...)
	if err != nil {
		return err
	}
	return nil
}

func Infof(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: InfoMessageType}, opts...)
	if err != nil {
		return err
	}
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Info(valueString)
	err = saveLogs(timeNow, "[info]", valueBase)
	if err != nil {
		return err
	}
	return nil
}

func InfofPrint(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: InfoMessageType}, opts...)
	if err != nil {
		return err
	}
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Info(valueString)
	return nil
}

func InfofSave(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: InfoMessageType}, opts...)
	if err != nil {
		return err
	}
	err = saveLogs(timeNow, "[info]", valueBase)
	if err != nil {
		return err
	}
	return nil
}

func Success(values ...any) error {
	values, opts := getOptsFromInterface(values)
	timeNow := time.Now()
	_, err := UpdateChannel(fmt.Sprint(values), &LogConfigs{msgType: SuccessMessageType}, opts...)
	if err != nil {
		return err
	}
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Success(valueString)
	err = saveLogs(timeNow, "[success]", values...)
	if err != nil {
		return err
	}
	return nil
}

func SuccessPrint(values ...any) error {
	values, opts := getOptsFromInterface(values)
	timeNow := time.Now()
	_, err := UpdateChannel(fmt.Sprint(values), &LogConfigs{msgType: SuccessMessageType}, opts...)
	if err != nil {
		return err
	}
	valueString := getFormattedValue(timeNow, values...)
	prettyPrints.Success(valueString)
	return nil
}

func SuccessSave(values ...any) error {
	values, opts := getOptsFromInterface(values)
	timeNow := time.Now()
	_, err := UpdateChannel(fmt.Sprint(values), &LogConfigs{msgType: SuccessMessageType}, opts...)
	if err != nil {
		return err
	}
	err = saveLogs(timeNow, "[success]", values...)
	if err != nil {
		return err
	}
	return nil
}

func Successf(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: SuccessMessageType}, opts...)
	if err != nil {
		return err
	}
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Success(valueString)
	err = saveLogs(timeNow, "[success]", valueBase)
	if err != nil {
		return err
	}
	return nil
}

func SuccessfPrint(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: SuccessMessageType}, opts...)
	if err != nil {
		return err
	}
	valueString := getFormattedValue(timeNow, valueBase)
	prettyPrints.Success(valueString)
	return nil
}

func SuccessfSave(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	valueBase := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(valueBase, &LogConfigs{msgType: SuccessMessageType}, opts...)
	if err != nil {
		return err
	}
	err = saveLogs(timeNow, "[success]", valueBase)
	if err != nil {
		return err
	}
	return nil
}

// Without saving
func Println(values ...any) error {
	values, opts := getOptsFromInterface(values)
	timeNow := time.Now()
	_, err := UpdateChannel(fmt.Sprint(values), &LogConfigs{msgType: PrintMessageType}, opts...)
	if err != nil {
		return err
	}
	value := getFormattedValue(timeNow, values...)
	prettyPrints.Print(value)
	return nil
}

func Printf(format string, args ...interface{}) error {
	args, opts := getOptsFromInterface(args)
	timeNow := time.Now()
	value := fmt.Sprintf(format, args...)
	_, err := UpdateChannel(value, &LogConfigs{msgType: PrintMessageType}, opts...)
	if err != nil {
		return err
	}
	value = getFormattedValue(timeNow, value)
	prettyPrints.Print(value)
	return nil
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

func saveLogs(timeNow time.Time, prefix string, values ...any) error {
	projectPath, err := getBinPath()
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
