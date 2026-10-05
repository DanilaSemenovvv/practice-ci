package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 2) != 4 {
		t.Error("Expected 4")
	}
}

func TestSub(t *testing.T) {
	if Sub(2, 2) != 0 {
		t.Error("Expected 0")
	}
}
