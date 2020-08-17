package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 2)
	fmt.Println(repeated)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestStringsCompare(t *testing.T) {
	a := "a"
	b := "b"
	got := strings.Compare(a, b)
	want := -1

	if got != want {
		t.Errorf("expected %q but got %q", got, want)

	}
}
