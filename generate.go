package main

func Generate(source []string) (string, error) {
	var final string

	// Collect the source of a page into blocks based on content.
	// E.g. paragraphs, code segments, or images
	collected := collect(source)

	// Wrap the blocks of the page in their respective tags.
	// Additionally, add hints which describe the content of each block.
	wrapped := wrap(collected)

	// Format the source blocks to match arbitrary standards.
	// These standards may include ideal line lengths per paragraph,
	// automatic spell checking, and escaping of character sequences for
	// friendlier alternatives.
	formatted := format(wrapped)

	// Embed the source blocks into an existing template, which may include
	// headers, footers, style-sheets, and scripts.
	embedded := embed(formatted)

	for _, line := range embedded {
		final += line
		final += "\n"
	}

	return final, nil
}
