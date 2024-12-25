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
	fmt.Println("Matches:", matches)
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

	// Get regex matches to do() and don't()
	reDo := regexp.MustCompile("do\\(\\)")
	reDont := regexp.MustCompile("don't\\(\\)")
	doIndex := reDo.FindAllStringSubmatchIndex(dataString, -1)
	dontIndex := reDont.FindAllStringSubmatchIndex(dataString, -1)

	// Returns the last index value of each match of reDo and reDont
	var doIndexSanitized []int
	for _, i := range doIndex {
		doIndexSanitized = append(doIndexSanitized, i[1])
	}
	var dontIndexSanitized []int
	for _, i := range dontIndex {
		dontIndexSanitized = append(dontIndexSanitized, i[1])
	}

	fmt.Println("doIndex:", doIndexSanitized)
	fmt.Println("dontIndex:", dontIndexSanitized)

	for _, value := range matchesIndexSanitized {
		fmt.Println("value:", value)

	}

	// Filtering logic
	// Iterate through matches index
	// If first index value in a match is less than the first index value in don't matches, then no need to filter out.
	// i.e. First match's first index value is 6.
	// The first index value in the don't() matches is 27.
	// Ok to proceed with business logic of adding first match to "validMatches" array.
	// Next match's first index value is 33.
	// The first don't() match's index value is 27.
	// Since 33 > 27, this means that value comes after,
	// and therefore we need to check if there is a do() match index value that is between 27 and 33.
	// The first (and only) do() match index value is 63, so a false is returned, and this is therefore not a "valid" match.
	// etc...

	var validMatches [][]int
	for _, matchIndex := range matchesIndexSanitized {
		for _, dontMatchIndex := range dontIndexSanitized {
			fmt.Println("matchIndex:", matchIndex[0], "dontMatchIndex:", dontMatchIndex)
			if matchIndex[0] < dontMatchIndex {
				validMatches = append(validMatches, matchIndex)
			} else if matchIndex[0] > dontMatchIndex {
				for _, doMatchIndex := range doIndexSanitized {
					if doMatchIndex < matchIndex[0] {
						validMatches = append(validMatches, matchIndex)
					}
				}
			}
		}
	}
	fmt.Println("Valid Matches:", validMatches)

	for _, validMatch := range validMatches {
		// filteredData = append(filteredData, []byte{dataString[validmatch[0]], dataString[validMatch[1]]})
		fmt.Println(string([]byte{dataString[validMatch[0]-1], dataString[validMatch[1]-1]}))
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
