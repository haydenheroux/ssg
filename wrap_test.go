package main

import "testing"

func TestWrapCode(t *testing.T) {
	expected := []string{"<code>", "<samp>1</samp>Line one<br/>", "<samp>2</samp>Line two<br/>", "<samp>3</samp>Line three<br/>", "</code>"}
	blocks := make([]Block, 1)
	blocks = append(blocks, Block{Type: Code, Content: []string{"Line one", "Line two", "Line three"}})
	result := wrap(blocks)
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestWrapImage(t *testing.T) {
	expected := []string{"<figure>", "<img src=\"path/to/image\" alt=\"Caption for the image\">", "<figcaption>Caption for the image</figcaption>", "</figure>"}
	blocks, err := collect([]string{"@img path/to/image Caption for the image"})
	if err != nil {
		t.Fatal(err)
	}
	result := wrap(blocks)
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestWrapList(t *testing.T) {
	expected := []string{"<ul>", "<li>List element one</li>", "<li>List element two</li>", "<li>List element three</li>", "</ul>"}
	blocks := make([]Block, 1)
	blocks = append(blocks, Block{Type: List, Content: []string{"List element one", "List element two", "List element three"}})
	result := wrap(blocks)
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}
