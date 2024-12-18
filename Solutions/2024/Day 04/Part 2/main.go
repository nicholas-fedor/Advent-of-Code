package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	SampleFile = "sample.txt"
	InputFile  = "input.txt"
)

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func parseFile(file *os.File) []string {
	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	return data
}

func main() {
	// Open the file
	file, err := openFile(SampleFile)
	// file, err := openFile(InputFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Copy the file's content to memory
	data := parseFile(file)
	for _, line := range data {
		fmt.Println(line)
	}

	// TODO: Day 4 - Part 2 Business Logic
}
