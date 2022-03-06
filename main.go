package main

import (
	"io/ioutil"
	"os"
)

func main() {

	filePath := os.Args[1]

	lines, err := Read(filePath)
	if err != nil {
		panic(err)
	}

	mutated, err := Generate(lines)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filePath+".html", []byte(mutated), 0666); err != nil {
		panic(err)
	}
}
