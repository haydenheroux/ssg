package main

import "testing"

func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestRead(t *testing.T) {
	filePath := "./tests/read_test"
	expected := []string{"Line one", "Line two", "Line three"}
	lines, err := Read(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(lines, expected) {
		t.Errorf("Got: %v, want: %v\n", lines, expected)
	}
}
