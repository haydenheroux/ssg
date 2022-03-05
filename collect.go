package main

import "strings"

type BlockType int64

const (
	Error BlockType = iota
	Normal
	Code
	Image
)

type Block struct {
	Type    BlockType
	Content []string
}

func collect(source []string) ([]Block, error) {

	var blocks []Block

	for _, line := range source {

		var _type BlockType
		var content []string

		indicator := line[0]
		if indicator == []byte("@")[0] {
			_type = determineType(line)
			content = determineContent(line, _type)
		} else {
			_type = Normal
			content = []string{line}
		}

		block := Block{Type: _type, Content: content}
		blocks = append(blocks, block)

	}

	return blocks, nil
}

func determineType(line string) BlockType {
	if line[0] != []byte("@")[0] {
		// However this case would end up happening, it ended up happening
		// and it should not have, so return an error
		return Error
	}

	content := line[1:]
	segments := strings.Split(content, " ")
	token := segments[0]

	switch token {
	case "code":
		return Code
	case "img", "image":
		return Image
	default:
		return Error
	}
}

func determineContent(line string, _type BlockType) []string {
	return []string{line}
}
