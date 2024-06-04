package prettyPrints

import (
	"github.com/fatih/color"
)

func Success(value string) {
	checkNewLine(&value)
	success := color.New(color.Bold, color.FgGreen)
	success.Print(value)
}

func Warning(value string) {
	checkNewLine(&value)
	warning := color.New(color.Bold, color.FgYellow)
	warning.Print(value)
}

func Fatal(value string) {
	checkNewLine(&value)
	fatal := color.New(color.Bold, color.FgRed)
	fatal.Print(value)
}

func Info(value string) {
	checkNewLine(&value)
	info := color.New()
	info.Print(value)
}

func Print(value string) {
	checkNewLine(&value)
	print := color.New(color.Faint)
	print.Print(value)
}
