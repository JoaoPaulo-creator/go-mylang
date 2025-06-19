package main

import (
	"fmt"
	"my-lang/mylang"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: my file.my")
		os.Exit(-1)
	}

	mylang.Run(args[1])
}
