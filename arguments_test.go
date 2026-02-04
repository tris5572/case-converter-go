package main

import (
	"reflect"
	"testing"
)

func TestProcessArgs_ParseFlags(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		wantS  CaseStyle
		wantIn []string
	}{
		{"flag first", []string{"--camel", "hello", "world"}, CaseCamel, []string{"hello", "world"}},
		{"no flag", []string{"hello", "world"}, CaseLower, []string{"hello", "world"}},
		{"multiple flags", []string{"--snake", "--camel", "x"}, CaseCamel, []string{"x"}},
		{"unknown flag treated as input", []string{"--unknown", "x"}, CaseLower, []string{"--unknown", "x"}},
		{"flag without hyphens", []string{"pascal", "A"}, CasePascal, []string{"A"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotIn := processArgs(tt.args)
			if gotS != tt.wantS {
				t.Fatalf("style = %v; want %v", gotS, tt.wantS)
			}
			if !reflect.DeepEqual(gotIn, tt.wantIn) {
				t.Fatalf("inputs = %v; want %v", gotIn, tt.wantIn)
			}
		})
	}
}
