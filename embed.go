package main

func embed(lines []string, headerFilePath string, footerFilePath string) ([]string, error) {
	var output []string

	header, err := Read(headerFilePath)
	if err != nil {
		return nil, err
	}

	footer, err := Read(footerFilePath)
	if err != nil {
		return nil, err
	}

	for _, line := range header {
		output = append(output, line)
	}

	for _, line := range lines {
		output = append(output, line)
	}

	for _, line := range footer {
		output = append(output, line)
	}

	return output, nil
}
