package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}

	if Add(0, 0) != 0 {
		t.Error("0+0 should be 0")
	}

	if Add(-1, 1) != 0 {
		t.Error("-1+1 should be 0")
	}

	if Add(100, -100) != 0 {
		t.Error("100-100 should be 0")
	}
}
