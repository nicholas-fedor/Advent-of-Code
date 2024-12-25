package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func filterData(data []string) [][]string {
	var filteredData [][]string

	// Converting the data to a single string
	var dataString string
	for _, line := range data {
		dataString += line
	}

	// Regex that looks for the mul(), do(), and don't() strings.
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(dataString, -1)

	// Filtering through the regex matches
	enabled := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				filteredData = append(filteredData, match)
			}
		}
	}

	return filteredData
}

func getDataString(data []string) string {
	var dataString string
	for _, line := range data {
		dataString += line
	}
	return dataString
}

func main() {
	// Open the file
	// file, err := openFile(SampleFile)
	file, err := openFile(InputFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Copy the file's content to memory
	data := parseFile(file)

	// Apply conditional filtering to return number pairs in a 2D int array
	matches := filterData(data)
	for _, match := range matches {
		fmt.Println(match)
	}

	// Calculate the sum of products
	var output int
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		output += num1 * num2
	}

	fmt.Println("Output:", output)
}
