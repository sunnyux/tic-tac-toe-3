package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	if got != "Hello, world!" {
		t.Errorf("HelloWorld() = %s; want \"Hello, world!\"", got)
	}
}