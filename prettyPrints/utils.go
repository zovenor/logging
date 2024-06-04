package prettyPrints

import (
	"fmt"
	"strings"
)

func checkNewLine(value *string) {
	if !strings.HasSuffix(*value, "\n") {
		*value = *value + "\n"
	}
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}
