package main

import "testing"

func TestCaseStyleString(t *testing.T) {
	tests := []struct {
		style CaseStyle
		want  string
	}{
		{CaseCamel, "camel"},
		{CasePascal, "pascal"},
		{CaseSnake, "snake"},
		{CaseKebab, "kebab"},
		{CaseLower, "lower"},
		{CaseUpper, "upper"},
		{CaseUpperSnake, "upper-snake"},
		{CaseUnknown, "unknown"},
	}

	for _, tt := range tests {
		if got := tt.style.String(); got != tt.want {
			t.Fatalf("%v.String() = %q; want %q", tt.style, got, tt.want)
		}
	}
}

func TestCaseStyleFromFlag(t *testing.T) {
	tests := []struct {
		flag string
		want CaseStyle
		ok   bool
	}{
		{"--camel", CaseCamel, true},
		{"camel", CaseCamel, true},
		{"--pascal", CasePascal, true},
		{"snake", CaseSnake, true},
		{"--kebab", CaseKebab, true},
		{"kebab-case", CaseKebab, true},
		{"--lower", CaseLower, true},
		{"uppercase", CaseUpper, true},
		{"--upper-snake", CaseUpperSnake, true},
		{"upper_snake_case", CaseUpperSnake, true},
		{"--unknown", CaseUnknown, false},
		{"", CaseUnknown, false},
	}

	for _, tt := range tests {
		if got, ok := CaseStyleFromFlag(tt.flag); ok != tt.ok || got != tt.want {
			t.Fatalf("CaseStyleFromFlag(%q) = (%v, %v); want (%v, %v)", tt.flag, got, ok, tt.want, tt.ok)
		}
	}
}
