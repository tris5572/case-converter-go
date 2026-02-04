package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	style, inputs := processArgs(args)
	fmt.Printf("Style: %s\n", style.String())
	for _, in := range inputs {
		words := splitString(in)
		converted := ConvertWords(words, style)
		fmt.Println(converted)
	}
}
