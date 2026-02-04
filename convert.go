package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// capitalize は文字列の先頭ルーンを大文字化して返します（残りはそのまま）
func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

// ConvertWords は分割済みの単語列を指定された CaseStyle に変換して返します
func ConvertWords(words []string, style CaseStyle) string {
	switch style {
	case CaseCamel:
		if len(words) == 0 {
			return ""
		}
		var b strings.Builder
		b.WriteString(strings.ToLower(words[0]))
		for _, w := range words[1:] {
			b.WriteString(capitalize(w))
		}
		return b.String()
	case CasePascal:
		var b strings.Builder
		for _, w := range words {
			b.WriteString(capitalize(w))
		}
		return b.String()
	case CaseSnake:
		{
			lowered := make([]string, 0, len(words))
			for _, w := range words {
				lowered = append(lowered, strings.ToLower(w))
			}
			return strings.Join(lowered, "_")
		}
	case CaseKebab:
		{
			lowered := make([]string, 0, len(words))
			for _, w := range words {
				lowered = append(lowered, strings.ToLower(w))
			}
			return strings.Join(lowered, "-")
		}
	case CaseLower:
		return strings.ToLower(strings.Join(words, ""))
	case CaseUpper:
		return strings.ToUpper(strings.Join(words, ""))
	default:
		return strings.Join(words, " ")
	}
}
