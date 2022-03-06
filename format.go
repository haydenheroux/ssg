package main

func format(lines []string) []string {

	var output []string

	for _, line := range lines {
		if len(line) == 0 || (len(line) == 1 && line[0] == []byte("\n")[0]) {
			continue
		}
		output = append(output, line)
	}

	return output
}
