package main

func Generate(source []string) (string, error) {
	var output string
	for _, line := range source {
		output += line
		output += "\n"
	}
	return output, nil
}
