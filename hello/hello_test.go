package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kelvin")
	want := "Hello, Kelvin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
