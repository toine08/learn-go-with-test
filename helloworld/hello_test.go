package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Toine")
	want := "Hello, Toine"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}