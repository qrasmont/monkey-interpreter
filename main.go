package main

import (
	"fmt"
	"os"

	"github.com/qrasmont/monkey-interpreter/repl"
)

func main() {
	fmt.Println("Starting the Monkey REPL")
	repl.Start(os.Stdin, os.Stdout)
}
