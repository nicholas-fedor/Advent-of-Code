// TODO:
// Iterate through the first set of numbers before the "|" and compare them
// against the second set.
// Count the number of matching numbers.
// If just one number, then output of 1 point for that card.
// For each additional match, multiply by 2.
// Summarize the total of all points.
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Input filepath
	filePath := "sample.txt"

	// Opens the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Instantiates the scanner object with the file
	scanner := bufio.NewScanner(file)

	// Enumerates through the file and maps it accordingly
	for scanner.Scan() {
		card := scanner.Text()
		// fmt.Println(card)

		// Enumerate through the line to parse for the winning numbers
		// var winningNumbers []int
		// var cardNumbers []int
		winningNumbersRegex := regexp.MustCompile(`\d+:\s*([\d\s]+?)\s*\|`)
		match := winningNumbersRegex.FindStringSubmatch(card)
		if len(match) > 1 {
			winningNumbers := match[1]
			fmt.Println(winningNumbers)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
