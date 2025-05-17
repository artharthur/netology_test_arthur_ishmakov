package main

import "testing"

func TestExample(t *testing.T) {
    want := "Привет, Jenkins!"
    got := "Привет, Jenkins!"
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
