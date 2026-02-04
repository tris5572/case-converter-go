package main

import (
	"reflect"
	"testing"
)

func TestProcessArgs_ParseFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantS    CaseStyle
		wantIn   []string
		wantHelp bool
	}{
		{"flag first", []string{"--camel", "hello", "world"}, CaseCamel, []string{"hello", "world"}, false},
		{"no flag", []string{"hello", "world"}, CaseLower, []string{"hello", "world"}, false},
		{"multiple flags", []string{"--snake", "--camel", "x"}, CaseCamel, []string{"x"}, false},
		{"unknown flag treated as input", []string{"--unknown", "x"}, CaseLower, []string{"--unknown", "x"}, false},
		{"flag without hyphens", []string{"pascal", "A"}, CasePascal, []string{"A"}, false},
		{"help flag", []string{"--help"}, CaseLower, []string{}, true},
		{"help with other args", []string{"--camel", "--help", "x"}, CaseCamel, []string{"x"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotIn, gotHelp := processArgs(tt.args)
			if gotS != tt.wantS {
				t.Fatalf("style = %v; want %v", gotS, tt.wantS)
			}
			if !reflect.DeepEqual(gotIn, tt.wantIn) {
				t.Fatalf("inputs = %v; want %v", gotIn, tt.wantIn)
			}
			if gotHelp != tt.wantHelp {
				t.Fatalf("help = %v; want %v", gotHelp, tt.wantHelp)
			}
		})
	}
}
