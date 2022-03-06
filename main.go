package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {

	var outputFilePath string

	flag.StringVar(&outputFilePath, "o", "", "The path to the output file. Defaults to <inputFilePath>.html if left blank.")
	flag.Parse()

	inputFilePath := flag.Arg(0)
	if len(inputFilePath) == 0 {
		log.Fatalf("error while loading file: no file specified")
	}

	lines, err := Read(inputFilePath)
	if err != nil {
		log.Fatalf("error while loading file: %v", err)
	}

	mutated, err := Generate(lines)
	if err != nil {
		log.Fatalf("error while generating content: %v", err)
	}

	if outputFilePath == "" {
		outputFilePath = inputFilePath + ".html"
	}

	if err := ioutil.WriteFile(outputFilePath, []byte(mutated), 0666); err != nil {
		log.Fatalf("error while writing file: %v", err)
	}
}
