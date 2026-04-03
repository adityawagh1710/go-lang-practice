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

func TestAdd2(t *testing.T) {
	// Setup (runs for all subtests)
	t.Log("Starting TestAdd")

	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Positive", 2, 3, 5},
		{"Negative", -1, 1, 0},
		{"Zero", 0, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}

	// Teardown (runs after all subtests)
	t.Log("Finished TestAdd")
}
