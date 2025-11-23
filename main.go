package main

import (
	"github.com/soadmized/idconv/cmd"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("some input is required")
	}

	input := os.Args[1]

	if err := cmd.Convert(input); err != nil {
		println(err.Error())

		os.Exit(1)
	}
}
