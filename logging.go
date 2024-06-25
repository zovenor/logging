package logging

import (
	types "github.com/zovenor/logging/v2/types"
)

// Warning

func Warning(value any, args ...any) error {
	return types.LogHandler(value, types.WarningMessageType, types.PrintAndSaveAction, args...)
}

func WarningPrint(value any, args ...any) error {
	return types.LogHandler(value, types.WarningMessageType, types.PrintAction, args...)
}

func WarningSave(value any, args ...any) error {
	return types.LogHandler(value, types.WarningMessageType, types.SaveAction, args...)
}

// Fatal

func Fatal(value any, args ...any) error {
	return types.LogHandler(value, types.FatalMessageType, types.PrintAndSaveAction, args...)
}

func FatalPrint(value any, args ...any) error {
	return types.LogHandler(value, types.FatalMessageType, types.PrintAction, args...)
}

func FatalSave(value any, args ...any) error {
	return types.LogHandler(value, types.FatalMessageType, types.SaveAction, args...)
}

// Info

func Info(value any, args ...any) error {
	return types.LogHandler(value, types.InfoMessageType, types.PrintAndSaveAction, args...)
}

func InfoPrint(value any, args ...any) error {
	return types.LogHandler(value, types.InfoMessageType, types.PrintAction, args...)
}

func InfoSave(value any, args ...any) error {
	return types.LogHandler(value, types.InfoMessageType, types.SaveAction, args...)
}

// Success

func Success(value any, args ...any) error {
	return types.LogHandler(value, types.SuccessMessageType, types.PrintAndSaveAction, args...)
}

func SuccessPrint(value any, args ...any) error {
	return types.LogHandler(value, types.SuccessMessageType, types.PrintAction, args...)
}

func SuccessSave(value any, args ...any) error {
	return types.LogHandler(value, types.SuccessMessageType, types.SaveAction, args...)
}

// Channeling
func SendToChan(msg *types.Message) error {
	return types.SendToChan(msg)
}
