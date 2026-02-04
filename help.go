package main

import "fmt"

// printHelp はコマンドの使い方を表示します
func printHelp() {
	fmt.Println("Usage: case-converter-go [options] [strings...]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help        Show this help message")
	fmt.Println("  --camel, camel    Output camelCase")
	fmt.Println("  --pascal, pascal  Output PascalCase")
	fmt.Println("  --snake, snake    Output snake_case")
	fmt.Println("  --kebab, kebab    Output kebab-case")
	fmt.Println("  --lower, lower    Output lowercase")
	fmt.Println("  --upper, upper    Output UPPERCASE")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  case-converter-go --camel \"Hello World\"")
	fmt.Println("  case-converter-go --snake someString")
}
