package main

import "fmt"

// printHelp はコマンドの使い方を表示します
func printHelp() {
	fmt.Println("Usage: case-converter-go [options] [strings...]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help        Show this help message")
	fmt.Println("  - (single dash)    Read input from stdin (or use no strings)")
	fmt.Println("  --camel, camel    Output camelCase")
	fmt.Println("  --pascal, pascal  Output PascalCase")
	fmt.Println("  --snake, snake    Output snake_case")
	fmt.Println("  --upper-snake, upper_snake_case    Output UPPER_SNAKE_CASE")
	fmt.Println("  --kebab, kebab    Output kebab-case")
	fmt.Println("  --train, train-case    Output Train-Case")
	fmt.Println("  --lower, lower    Output lowercase")
	fmt.Println("  --upper, upper    Output UPPERCASE")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  case-converter-go --camel \"Hello World\"")
	fmt.Println("  case-converter-go --snake someString")
}
