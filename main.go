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

	data, err := os.ReadFile(args[1])
	if err != nil {
		panic(err.Error())
	}

	mylang.Run(string(data))
}
