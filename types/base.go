package types

import (
	"fmt"

	"github.com/zovenor/logging/v2/prettyPrints"
)

func PrintAndSave(message *Message) error {
	PrintOnly(message)
	err := SaveOnly(message)
	if err != nil {
		return err
	}
	return nil
}

func PrintOnly(message *Message) {
	formattedValue := getFormattedValue(message.time, message.value)
	switch message.msgType {
	case WarningMessageType:
		prettyPrints.Warning(formattedValue)
	case FatalMessageType:
		prettyPrints.Fatal(formattedValue)
	case SuccessMessageType:
		prettyPrints.Success(formattedValue)
	case InfoMessageType:
		prettyPrints.Info(formattedValue)
	}
}

func SaveOnly(message *Message) error {
	prefix := "unknown"
	switch message.msgType {
	case WarningMessageType:
		prefix = warningPrefix
	case FatalMessageType:
		prefix = fatalPrefix
	case SuccessMessageType:
		prefix = successPrefix
	case InfoMessageType:
		prefix = infoPrefix
	}
	prefix = fmt.Sprintf("[%v]", prefix)
	err := saveLogs(message, prefix)
	if err != nil {
		return err
	}
	return nil
}

type HandlerAction uint8

const (
	PrintAndSaveAction HandlerAction = iota
	PrintAction
	SaveAction
)

func LogHandler(value any, msgType MessageType, action HandlerAction, args ...any) error {
	args, opts := getOptsFromInterface(args...)
	if len(args) > 0 {
		value = fmt.Sprintf(fmt.Sprint(value), args...)
	} else {
		value = fmt.Sprint(value)
	}
	msg := NewMessage(value, msgType, opts...)
	msg.SetAction(action)
	switch action {
	case PrintAndSaveAction:
		return PrintAndSave(msg)
	case PrintAction:
		PrintOnly(msg)
		return nil
	case SaveAction:
		return SaveOnly(msg)
	}
	return nil
}
