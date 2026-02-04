package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	style, inputs, help := processArgs(args)
	if help {
		printHelp()
		return
	}
	fmt.Printf("Style: %s\n", style.String())
	for _, in := range inputs {
		words := splitString(in)
		converted := ConvertWords(words, style)
		fmt.Println(converted)
	}
}
