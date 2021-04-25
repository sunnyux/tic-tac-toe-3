package main

import "testing"

func TestMySanity(t *testing.T) {
	if true != true {
		t.Errorf("Oh no ;(")
	}
}