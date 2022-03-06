package main

import "fmt"

func wrap(blocks []Block) []string {
	var output []string

	for _, block := range blocks {

		// Pre
		if block.Type == Code {
			output = append(output, "<code>")
		}

		// Content
		for index, line := range block.Content {

			if block.Type == Normal {
				output = append(output, "<p>"+line+"</p>")
			}

			if block.Type == Code {
				output = append(output, fmt.Sprintf("<samp>%d</samp>", index+1)+line)
			}

		}

		// Post
		if block.Type == Code {
			output = append(output, "</code>")
		}

	}

	return output
}
