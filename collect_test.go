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

func TestDetermineTypeList(t *testing.T) {
	result := determineType("@list path/to/list")
	if result != List {
		t.Errorf("got: %v, want: %v", result, List)
	}
}

func TestDetermineContentCode(t *testing.T) {
	expected, _ := Read("./tests/read_test")
	result, err := determineContent("@code ./tests/read_test", Code)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestDetermineContentImage(t *testing.T) {
	expected := []string{"path/to/image", "Caption for the image"}
	result, err := determineContent("@img path/to/image Caption for the image", Image)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestDetermineContentList(t *testing.T) {
	expected, _ := Read("./tests/list_test")
	result, err := determineContent("@list ./tests/list_test", List)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestCode(t *testing.T) {
	expected, _ := Read("./tests/read_test")
	input := "@code ./tests/read_test"
	_type := determineType(input)
	if _type != Code {
		t.Errorf("got: %v, want: %v", _type, Code)
	}
	result, err := determineContent(input, _type)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestImage(t *testing.T) {
	expected := []string{"path/to/image", "Caption for the image"}
	input := "@img path/to/image Caption for the image"
	_type := determineType(input)
	if _type != Image {
		t.Errorf("got: %v, want: %v", _type, Image)
	}
	result, err := determineContent(input, _type)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestList(t *testing.T) {
	expected, _ := Read("./tests/list_test")
	input := "@list ./tests/list_test"
	_type := determineType(input)
	if _type != List {
		t.Errorf("got: %v, want: %v", _type, List)
	}
	result, err := determineContent(input, _type)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}
