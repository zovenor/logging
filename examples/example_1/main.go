package main

import (
	"github.com/zovenor/logging/v2"
	"github.com/zovenor/logging/v2/types"
)

func main() {
	logging.Success("success message")
	logging.SuccessPrint("success message")
	logging.SuccessSave("success message")

	logging.Fatal("fatal message")
	logging.FatalPrint("fatal message")
	logging.FatalSave("fatal message")

	logging.Warning("warning message")
	logging.WarningPrint("warning message")
	logging.WarningSave("warning message")

	logging.Info("info message")
	logging.InfoPrint("info message")
	logging.InfoSave("info message")

	logging.InfoPrint("info message: %v", 10000)
	logging.InfoPrint("info message", types.WithChanneling)
	logging.InfoPrint("info message: %v %v", 999, types.WithChanneling, true)
}
