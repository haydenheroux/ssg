package main

import "testing"

func TestDetermineTypeCode(t *testing.T) {
	result := determineType("@code path/to/code")
	if result != Code {
		t.Errorf("got: %v, want: %v", result, Code)
	}
}

func TestDetermineTypeImage(t *testing.T) {
	result := determineType("@img path/to/image")
	if result != Image {
		t.Errorf("got: %v, want: %v", result, Image)
	}
}

func TestDetermineContentCode(t *testing.T) {
	expected, _ := Read("./tests/read_test")
	t.Log(expected)
	result, err := determineContent("@code ./tests/read_test", Code)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}
