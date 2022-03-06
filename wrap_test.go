package main

import "testing"

func TestWrapCode(t *testing.T) {
	expected := []string{"<code>", "<samp>1</samp>Line one", "<samp>2</samp>Line two", "<samp>3</samp>Line three", "</code>"}
	blocks := make([]Block, 1)
	blocks = append(blocks, Block{Type: Code, Content: []string{"Line one", "Line two", "Line three"}})
	result := wrap(blocks)
	if !equal(result, expected) {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}
