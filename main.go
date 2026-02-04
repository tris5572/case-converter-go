package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	style, inputs := processArgs(args)
	// 今は解析結果を表示します。次のステップで実際の変換を適用します。
	fmt.Printf("Style: %s\n", style.String())
	for _, in := range inputs {
		fmt.Println(in)
	}
}
