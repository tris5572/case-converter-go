package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadInputs_NoArgs_ReadsStdin(t *testing.T) {
	stdin := bytes.NewBufferString("a\nb\n\n c \n")
	got, err := readInputs([]string{}, stdin)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func TestReadInputs_Dash_ReplacedWithStdinLines(t *testing.T) {
	stdin := bytes.NewBufferString("x\ny\n")
	got, err := readInputs([]string{"before", "-", "after"}, stdin)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"before", "x", "y", "after"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func TestReadInputs_NoDash_NoStdinUsed(t *testing.T) {
	stdin := bytes.NewBufferString("should not be read\n")
	got, err := readInputs([]string{"one", "two"}, stdin)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"one", "two"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}
