package main

import "fmt"

func wrap(blocks []Block) []string {
	var output []string

	for _, block := range blocks {

		// Pre
		if block.Type == Code {
			output = append(output, "<code>")
		}

		if block.Type == Image {
			output = append(output, "<figure>")
		}

		if block.Type == List {
			output = append(output, "<ul>")
		}

		// Content
		for index, line := range block.Content {

			if block.Type == Normal {
				output = append(output, fmt.Sprintf("<p>%s</p>", line))
			}

			if block.Type == Code {
				output = append(output, fmt.Sprintf("<samp>%d</samp>%s<br/>", index+1, line))
			}

			if block.Type == Image {
				filePath := line
				caption := block.Content[index+1]
				output = append(output, fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", filePath, caption))
				output = append(output, fmt.Sprintf("<figcaption>%s</figcaption>", caption))
				// There exists a case for images; since all image blocks should be
				// of the same format (two elements in the content slice), we can extract
				// the content we need then exit out of the image block. Since we already
				// extracted the content we needed beforehand and the exit prevents further
				// access of content, we should avoid all errors related to incorrect accesses.
				break
			}

			if block.Type == List {
				output = append(output, fmt.Sprintf("<li>%s</li>", line))
			}

			if block.Type == Header {
				title := line
				output = append(output, fmt.Sprintf("<h2>%s</h2>", title))
			}

		}

		// Post
		if block.Type == Code {
			output = append(output, "</code>")
		}

		if block.Type == Image {
			output = append(output, "</figure>")
		}

		if block.Type == List {
			output = append(output, "</ul>")
		}

	}

	return output
}
