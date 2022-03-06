package main

import (
	"strings"
)

type BlockType int64

const (
	Error BlockType = iota
	Normal
	Code
	Image
	Header
	List
)

type Block struct {
	Type    BlockType
	Content []string
}

func collect(source []string) ([]Block, error) {

	var blocks []Block

	for _, line := range source {

		var _type BlockType

		// Skip blank lines
		if len(line) == 0 {
			continue
		}

		indicator := line[0]
		if indicator == []byte("@")[0] {
			_type = determineType(line)
		} else {
			_type = Normal
		}
		content, err := determineContent(line, _type)
		if err != nil {
			return blocks, err
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
	case "code", "include":
		return Code
	case "img", "image":
		return Image
	case "section", "segment":
		return Header
	case "list":
		return List
	default:
		return Error
	}
}

func determineContent(line string, _type BlockType) ([]string, error) {
	if _type == Normal {
		return []string{line}, nil
	}

	segments := strings.Split(line, " ")
	data := segments[1:]

	switch _type {
	case Code:
		copied, err := Read(data[0])
		return copied, err
	case Image:
		filePath := data[0]
		text := data[1:]
		caption := strings.Join(text, " ")
		return []string{filePath, caption}, nil
	case Header:
		text := data
		title := strings.Join(text, " ")
		return []string{title}, nil
	case List:
		copied, err := Read(data[0])
		return copied, err
	default:
		return []string{"ERROR: UNIMPLEMENTED"}, nil
	}
}
