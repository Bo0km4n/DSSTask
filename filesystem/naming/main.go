package main

import (
	"fmt"
	"os"

	"github.com/Bo0km4n/DSSTask/filesystem/naming/repl"
)

func main() {
	fmt.Println("Hello this is dummy file system!")
	fmt.Println("I'm supporting 'cd' or 'ls' command only... sry ><")

	repl.Start(os.Stdin, os.Stdout)
}
