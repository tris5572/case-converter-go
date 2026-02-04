package main

// splitString のテスト
import "testing"

func TestSplitString(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"HelloWorld", []string{"Hello", "World"}},
		{"hello_world", []string{"hello", "world"}},
		{"kebab-case-string", []string{"kebab", "case", "string"}},
		{"Mixed_Case-String Example", []string{"Mixed", "Case", "String", "Example"}},
		{"word", []string{"word"}},
		{"theACRONYMWord", []string{"the", "ACRONYM", "Word"}},
		{"version1Number2", []string{"version", "1", "Number", "2"}},
		{"JSON2XML", []string{"JSON", "2", "XML"}},
		{"123abc", []string{"123", "abc"}},
		{"hello世界", []string{"hello", "世界"}},
		{"こんにちは世界", []string{"こんにちは世界"}},
		{"", []string{}},
	}

	for _, test := range tests {
		result := splitString(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("splitString(%q) = %v; want %v", test.input, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("splitString(%q) = %v; want %v", test.input, result, test.expected)
				break
			}
		}
	}
}
