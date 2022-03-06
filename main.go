package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	filePath := os.Args[1]

	lines, err := Read(filePath)
	if err != nil {
		log.Fatalf("error while loading file: %v", err)
	}

	mutated, err := Generate(lines)
	if err != nil {
		log.Fatalf("error while generating content: %v", err)
	}

	if err := ioutil.WriteFile(filePath+".html", []byte(mutated), 0666); err != nil {
		log.Fatalf("error while writing file: %v", err)
	}
}
