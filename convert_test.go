package main

import "testing"

func TestConvertWords(t *testing.T) {
	tests := []struct {
		input string
		style CaseStyle
		want  string
	}{
		{"HelloWorld", CaseCamel, "helloWorld"},
		{"HelloWorld", CasePascal, "HelloWorld"},
		{"HelloWorld", CaseSnake, "hello_world"},
		{"HelloWorld", CaseKebab, "hello-world"},
		{"HelloWorld", CaseLower, "helloworld"},
		{"HelloWorld", CaseUpper, "HELLOWORLD"},

		{"theACRONYMWord", CaseCamel, "theACRONYMWord"},
		{"theACRONYMWord", CasePascal, "TheACRONYMWord"},
		{"theACRONYMWord", CaseSnake, "the_acronym_word"},
		{"theACRONYMWord", CaseKebab, "the-acronym-word"},
		{"theACRONYMWord", CaseLower, "theacronymword"},
		{"theACRONYMWord", CaseUpper, "THEACRONYMWORD"},

		{"JSON2XML", CaseCamel, "json2XML"},
		{"JSON2XML", CasePascal, "JSON2XML"},
		{"JSON2XML", CaseSnake, "json_2_xml"},
		{"JSON2XML", CaseKebab, "json-2-xml"},

		{"hello世界", CaseCamel, "hello世界"},
		{"hello世界", CasePascal, "Hello世界"},
		{"hello世界", CaseSnake, "hello_世界"},
		{"hello世界", CaseKebab, "hello-世界"},
		{"hello世界", CaseLower, "hello世界"},
		{"hello世界", CaseUpper, "HELLO世界"},

		{"HelloWorld", CaseUpperSnake, "HELLO_WORLD"},
		{"theACRONYMWord", CaseUpperSnake, "THE_ACRONYM_WORD"},
		{"JSON2XML", CaseUpperSnake, "JSON_2_XML"},
	}

	for _, tt := range tests {
		words := splitString(tt.input)
		got := ConvertWords(words, tt.style)
		if got != tt.want {
			t.Fatalf("ConvertWords(%q, %v) = %q; want %q", tt.input, tt.style, got, tt.want)
		}
	}
}
