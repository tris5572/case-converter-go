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
	// 入力を展開（"-" を stdin の行で置換、引数が空なら stdin 全体を読み取る）
	inputs, err := readInputs(inputs, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading stdin: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Style: %s\n", style.String())
	for _, in := range inputs {
		words := splitString(in)
		converted := ConvertWords(words, style)
		fmt.Println(converted)
	}
}
