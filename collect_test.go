package main

import "testing"

func TestDetermineTypeImage(t *testing.T) {
	result := determineType("@img path/to/image")
	if result != Image {
		t.Logf("got: %v, want: %v", result, Image)
		t.Fail()
	}
}

func TestDetermineTypeCode(t *testing.T) {
	result := determineType("@code path/to/code")
	if result != Code {
		t.Logf("got: %v, want: %v", result, Code)
		t.Fail()
	}
}
