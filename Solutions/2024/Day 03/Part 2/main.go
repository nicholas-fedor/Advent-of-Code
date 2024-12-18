package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func filterData(data []string) [][]int {
	var filteredData [][]int
	// First, convert the []string array into a single string
	dataString := getDataString(data)

	// Regex matching for mul()
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := reMul.FindAllStringSubmatch(dataString, -1)
	matchesIndex := reMul.FindAllStringSubmatchIndex(dataString, -1)
	// Output: [[mul(2,4) 2 4] [mul(5,5) 5 5] [mul(11,8) 11 8] [mul(8,5) 8 5]]
	fmt.Println(matches)
	// fmt.Println(matchesIndex)

	// Sanitize matchesIndex
	var matchesIndexSanitized [][]int
	for _, match := range matchesIndex {
		matchesIndexSanitized = append(matchesIndexSanitized, []int{match[3], match[5]})
	}
	fmt.Println("Matches Index:", matchesIndexSanitized)

	// Iterate through matches to return digit pairs
	// var digitPairs [][]string
	// for _, match := range matches {
	// 	digitPairs = append(digitPairs, []string{match[1], match[2]})
	// }
	// fmt.Println(digitPairs)

	reDo := regexp.MustCompile("do\\(\\)")
	reDont := regexp.MustCompile("don't\\(\\)")
	doIndex := reDo.FindAllStringSubmatchIndex(dataString, -1)
	dontIndex := reDont.FindAllStringSubmatchIndex(dataString, -1)
	fmt.Println("doIndex:", doIndex)
	fmt.Println("dontIndex:", dontIndex)

	for _, value := range matchesIndexSanitized {
		fmt.Println("value:", value)
		
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

func getSumProduct(filteredData any) any {
	panic("unimplemented")
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

	// Apply conditional filtering to return number pairs in a 2D int array
	filteredData := filterData(data)
	fmt.Println(filteredData)

	// sumProductFiltered := getSumProduct(filteredData)

	// fmt.Println(sumProductFiltered)
}
