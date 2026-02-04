package main

import (
	"slices"
	"unicode"
)

// 文字列を区切り文字で分割する
//
// - キャメルケース・スネークケース・ケバブケースなどに対応
// - 英字、数字、その他の文字種の遷移で分割
// - 大文字連続の後に小文字が来る場合で分割（例: ACRONYMWord -> ACRONYM, Word）
// - Latin文字と非Latin文字（例: 漢字、ひらがな、カタカナ、キリル文字など）の遷移で分割
func splitString(s string) []string {
	delimiters := []rune{' ', '_', '-'}
	var result []string
	current := []rune{}

	rs := []rune(s)
	isDelimiter := func(r rune) bool {
		return slices.Contains(delimiters, r)
	}

	for i := range rs {
		r := rs[i]
		if isDelimiter(r) {
			if 0 < len(current) {
				result = append(result, string(current))
				current = []rune{}
			}
			continue
		}

		// 文字種（Letter, Digit, Other）の遷移とキャメルケース・アクロニム境界を検出
		if 0 < i {
			prev := rs[i-1]
			var next rune
			if len(rs) > i+1 {
				next = rs[i+1]
			}

			isLetter := unicode.IsLetter(r)
			prevIsLetter := unicode.IsLetter(prev)
			isDigit := unicode.IsDigit(r)
			prevIsDigit := unicode.IsDigit(prev)

			// Letter の内部: 小文字->大文字 の遷移や、ACRONYMWord のような大文字連続の後に小文字が来る場合で分割
			if isLetter && prevIsLetter {
				// Latin と非-Latin（例: latin <-> 漢字）でスクリプトが異なる場合は分割
				prevLatin := unicode.In(prev, unicode.Latin)
				currLatin := unicode.In(r, unicode.Latin)
				if prevLatin != currLatin {
					if 0 < len(current) {
						result = append(result, string(current))
						current = []rune{}
					}
				} else if unicode.IsUpper(r) && unicode.IsLower(prev) {
					if 0 < len(current) {
						result = append(result, string(current))
						current = []rune{}
					}
				} else if unicode.IsUpper(r) && unicode.IsUpper(prev) && i+1 < len(rs) && unicode.IsLower(next) {
					if 0 < len(current) {
						result = append(result, string(current))
						current = []rune{}
					}
				}
			} else if isDigit && !prevIsDigit {
				// 英字/その他 -> 数字 の遷移
				if 0 < len(current) {
					result = append(result, string(current))
					current = []rune{}
				}
			} else if !isDigit && prevIsDigit {
				// 数字 -> 英字/その他 の遷移
				if 0 < len(current) {
					result = append(result, string(current))
					current = []rune{}
				}
			} else if isLetter != prevIsLetter {
				// 英字 <-> その他（例: latin <-> 日本語）の遷移
				if 0 < len(current) {
					result = append(result, string(current))
					current = []rune{}
				}
			}
		}

		current = append(current, r)
	}

	if 0 < len(current) {
		result = append(result, string(current))
	}

	return result
}
