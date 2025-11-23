package main

import (
	"github.com/soadmized/idconv/cmd"
	"os"
)

func main() {
	input := os.Args[1]

	if err := cmd.Convert(input); err != nil {
		println(err.Error())

		os.Exit(1)
	}
}
