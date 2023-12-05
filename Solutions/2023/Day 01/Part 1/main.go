package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func findFirstAndLastDigits(line string) (rune, rune) {
	var firstDigit, lastDigit rune

	// Iterate over the line's characters to find digits.
	for _, char := range line {
		if unicode.IsDigit(char) {
			if firstDigit == 0 {
				firstDigit = char // Set the first digit
			}
			lastDigit = char // Update the last digit through iteration
		}
	}
	return firstDigit, lastDigit
}

func main() {
	var totalSum int

	// Opens input file.
	input, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(input)

	// Iterate over each line scanned.
	for scanner.Scan() {
		// Get the text of the scanned line.
		line := scanner.Text()

		// Extract first and last digits from each line.
		firstRune, lastRune := findFirstAndLastDigits(line)
		firstNumber := int(firstRune - '0')
		lastNumber := int(lastRune - '0')

		// Create the two-digit number based on first and last digits.
		lineRune := rune(firstNumber*10 + lastNumber)
		lineNumber := int(lineRune)

		// Add the line sum to the total sum.
		totalSum += lineNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Total Sum of Line Sums:", totalSum)
}

// Answer = 53921
