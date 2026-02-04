package main

import (
	"bufio"
	"io"
	"strings"
)

// readInputs は与えられた引数配列を処理して最終的な入力リストを返します。
// - 引数が空の場合、stdin から読み取った各行（空行は無視）を返す
// - 引数に "-" が含まれる場合、その位置を stdin の行で置き換える
func readInputs(args []string, stdin io.Reader) ([]string, error) {
	// stdin が nil の場合は空とみなす
	if len(args) == 0 {
		return readLines(stdin)
	}

	out := make([]string, 0, len(args))
	for _, a := range args {
		if a == "-" {
			lines, err := readLines(stdin)
			if err != nil {
				return nil, err
			}
			out = append(out, lines...)
		} else {
			out = append(out, a)
		}
	}
	return out, nil
}

// readLines は r から行を読み、トリムして空行を除外した文字列スライスを返します
func readLines(r io.Reader) ([]string, error) {
	if r == nil {
		return []string{}, nil
	}
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if l == "" {
			continue
		}
		lines = append(lines, l)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
