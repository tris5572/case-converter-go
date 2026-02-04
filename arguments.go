package main

// コマンドライン引数を処理する
// フラグ（例: --camel）を検出して CaseStyle を決定し、残りの引数を入力リストとして返す
// 第3戻り値はヘルプ表示を要求するかどうか
func processArgs(args []string) (CaseStyle, []string, bool) {
	// デフォルトは lower
	style := CaseLower
	inputs := []string{}
	help := false
	for _, arg := range args {
		// ヘルプフラグを優先的に検出
		if arg == "-h" || arg == "--help" || arg == "help" {
			help = true
			continue
		}
		if s, ok := CaseStyleFromFlag(arg); ok {
			style = s
			continue
		}
		// フラグでなければ入力として扱う
		inputs = append(inputs, arg)
	}
	return style, inputs, help
}
