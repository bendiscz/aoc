package aoc

import (
	"reflect"
	"testing"
)

func TestTask_Reset(t *testing.T) {
	in := Example("xyz")
	if want, got := "xyz", in.ReadAll(); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}
	if want, got := "", in.ReadAll(); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}

	in.Reset()

	if want, got := "xyz", in.ReadAll(); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}
}

func TestTask_ReadLine(t *testing.T) {
	in := Example(`
one
two
three
`)

	if !in.NextLine() {
		t.Errorf("want: <next line>, got: EOF")
	}
	if want, got := "one", in.Line(); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}

	in.NextLine()

	if !in.NextLine() {
		t.Errorf("want: <next line>, got: EOF")
	}
	if want, got := "three", in.Line(); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}

	if in.NextLine() {
		t.Errorf("want: EOF, got: <next line>")
	}
	if want, got := "", in.Line(); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}
}

func TestTask_ParseLine(t *testing.T) {
	in := Example(`
1,2
69,420
`)

	if !in.NextLine() {
		t.Errorf("want: <next line>, got: EOF")
	}
	if want, got := []string{"1,2", "1", "2"}, in.Parse(`^(\d+),(\d+)$`); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %#v, got: %#v", want, got)
	}

	if !in.NextLine() {
		t.Errorf("want: <next line>, got: EOF")
	}
	if want, got := []string{"69,420", "69", "420"}, in.Parse(`^(\d+),(\d+)$`); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %#v, got: %#v", want, got)
	}

	if in.NextLine() {
		t.Errorf("want: EOF, got: <next line>")
	}
	if want, got := []string(nil), in.Parse(`^(\d+),(\d+)$`); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %#v, got: %#v", want, got)
	}
}

func TestExample_Trim(t *testing.T) {
	in := Example(`

 a
bcd
 e

`)

	if !in.NextLine() {
		t.Errorf("want: <next line>, got: EOF")
	}
	if want, got := " a", in.Line(); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}

	if !in.NextLine() {
		t.Errorf("want: <next line>, got: EOF")
	}
	if want, got := "bcd", in.Line(); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}

	if !in.NextLine() {
		t.Errorf("want: <next line>, got: EOF")
	}
	if want, got := " e", in.Line(); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}

	if in.NextLine() {
		t.Errorf("want: EOF, got: <next line>")
	}
}
