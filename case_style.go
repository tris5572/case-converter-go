package main

import "strings"

// CaseStyle は出力フォーマットの列挙を表します
type CaseStyle int

const (
	CaseUnknown CaseStyle = iota
	CaseCamel
	CasePascal
	CaseSnake
	CaseKebab
	CaseLower
	CaseUpper
	CaseUpperSnake
	CaseTrain
)

// String は CaseStyle を人間可読な名前で返します
func (c CaseStyle) String() string {
	switch c {
	case CaseCamel:
		return "camel"
	case CasePascal:
		return "pascal"
	case CaseSnake:
		return "snake"
	case CaseKebab:
		return "kebab"
	case CaseLower:
		return "lower"
	case CaseUpper:
		return "upper"
	case CaseUpperSnake:
		return "upper-snake"
	case CaseTrain:
		return "train"
	default:
		return "unknown"
	}
}

// CaseStyleFromFlag はコマンドライン引数（例: "--camel" や "camel"）から CaseStyle を返します
// 受け取れない場合は第二戻り値に false を返します
func CaseStyleFromFlag(s string) (CaseStyle, bool) {
	if s == "" {
		return CaseUnknown, false
	}
	// 前方のハイフンを取り除く
	s = strings.TrimLeft(s, "-")
	s = strings.ToLower(s)

	switch s {
	case "camel":
		return CaseCamel, true
	case "pascal":
		return CasePascal, true
	case "snake", "snake_case":
		return CaseSnake, true
	case "kebab", "kebab-case":
		return CaseKebab, true
	case "lower", "lowercase":
		return CaseLower, true
	case "upper", "uppercase":
		return CaseUpper, true
	case "upper-snake", "upper_snake_case":
		return CaseUpperSnake, true
	case "train", "train-case", "train_case":
		return CaseTrain, true
	default:
		return CaseUnknown, false
	}
}
