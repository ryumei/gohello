package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := "Hello world"
	actual := greeting("world")
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
