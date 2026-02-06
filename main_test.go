package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hellow Go"

	got := hello()

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
