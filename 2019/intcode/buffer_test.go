package intcode

import "testing"

func TestBuffer(t *testing.T) {
	b := newBuffer()

	if want, got := 0, b.len(); got != want {
		t.Fatalf("b.len() = %#v; want %#v", got, want)
	}

	{
		_, ok := b.read()
		if ok {
			t.Fatalf("b.read() = _, %#v; want _, true", ok)
		}
	}

	v, w := 0, 0
	for i := 0; i < 700; i++ {
		b.write(v)
		v++
	}

	if want, got := 700, b.len(); got != want {
		t.Fatalf("b.len() = %#v; want %#v", got, want)
	}

	for i := 0; i < 1000; i++ {
		b.write(v)
		v++
	}

	if want, got := 1700, b.len(); got != want {
		t.Fatalf("b.len() = %#v; want %#v", got, want)
	}

	for i := 0; i < 1600; i++ {
		got, ok := b.read()
		if !ok || got != w {
			t.Fatalf("b.read() = %#v, %#v; want %#v, true", got, ok, w)
		}
		w++
	}

	if want, got := 100, b.len(); got != want {
		t.Fatalf("b.len() = %#v; want %#v", got, want)
	}

	for i := 0; i < 2000; i++ {
		b.write(v)
		v++
	}

	if want, got := 2100, b.len(); got != want {
		t.Fatalf("b.len() = %#v; want %#v", got, want)
	}

	for i := 0; i < 2100; i++ {
		got, ok := b.read()
		if !ok || got != w {
			t.Fatalf("b.read() = %#v, %#v; want %#v, true", got, ok, w)
		}
		w++
	}

	if want, got := 0, b.len(); got != want {
		t.Fatalf("b.len() = %#v; want %#v", got, want)
	}

	{
		_, ok := b.read()
		if ok {
			t.Fatalf("b.read() = _, %#v; want _, true", ok)
		}
	}
}
