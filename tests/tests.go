package main

import (
	"fmt"

	"github.com/zovenor/logging/logging"
)

func main() {
	logging.Info("info message")
	logging.Success("success message")
	logging.Fatal(fmt.Errorf("%v", "error message"))
	logging.Warning(fmt.Errorf("%v", "warning message"))
	logging.Println("it is just println function")
	logging.Printf("and it is just printf function with value = %v", 10)

}
