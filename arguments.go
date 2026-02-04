package main

// コマンドライン引数を処理する
// フラグ（例: --camel）を検出して CaseStyle を決定し、残りの引数を入力リストとして返す
func processArgs(args []string) (CaseStyle, []string) {
	// デフォルトは lower
	style := CaseLower
	var inputs []string
	for _, arg := range args {
		if s, ok := CaseStyleFromFlag(arg); ok {
			style = s
			continue
		}
		// フラグでなければ入力として扱う
		inputs = append(inputs, arg)
	}
	return style, inputs
}
