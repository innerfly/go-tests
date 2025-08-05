package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "Hi there"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
