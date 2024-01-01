package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Input filepath
	filePath := "input.txt"

	// Opens the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Instantiates the scanner object with the file
	scanner := bufio.NewScanner(file)
	
	// totalPoints tracks the overall total points earned across all cards
	var totalPoints int

	// Enumerates through the file and handles each line as a card
	for scanner.Scan() {
		card := scanner.Text()
		// Outputs the line/card for debugging purposes
		fmt.Println(card)

		// cardPoints tracks the total points for a card
		var cardPoints int

		// Handles the card by parsing the sets of numbers and matches
		re := regexp.MustCompile(`\d+:\s*([\d\s]+?)\s*\|\s*([\d\s]+)`)
		match := re.FindStringSubmatch(card)
		if len(match) >= 3 {
			// Handling of winning numbers i.e. left of the pipe "|"
			winningNumbersStrSlice := strings.Fields(match[1]) // Split by spaces
			winningNumbersIntSlice := make([]int, len(winningNumbersStrSlice))
			for i, winningNumberStr := range winningNumbersStrSlice {
				winningNumbersIntSlice[i] = atoi(winningNumberStr)
			}
			
			// Handling of card numbers i.e. right of the pipe "|"
			cardNumbersStrSlice := strings.Fields(match[2]) // Split by spaces
			cardNumbersIntSlice := make([]int, len(cardNumbersStrSlice))
			for i, cardNumberStr := range cardNumbersStrSlice {
				cardNumbersIntSlice[i] = atoi(cardNumberStr)
			}


			// Iterates through each card's numbers to compare against the
			// card's winning numbers
			for _, cardNumber := range cardNumbersIntSlice {
				for _, winningNumber := range winningNumbersIntSlice {
					// If the card number is a winning number...
					if cardNumber == winningNumber {
						fmt.Println("Matching Number:", cardNumber)
						switch {
						// If it's the first match, cardPoints will equal 0
						case cardPoints == 0:
							// Increment cardPoints by 1
							cardPoints++
						// For subsequent matches...
						case cardPoints > 0:
							// Multiply cardPoints by 2
							cardPoints = cardPoints * 2
						}
						// Output for debugging
						fmt.Println("Card Points:", cardPoints)
					}
				}
			}
			// Sums the total points for each card
			totalPoints += cardPoints
			// Output for debugging
			fmt.Println("Total Points:", totalPoints)
			fmt.Println()
		}
		
	}
	
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	
	// Outputs the final answer
	fmt.Println()
	fmt.Println("Total Points:", totalPoints)
	fmt.Println()
}

// atoi converts a string digit to an integer
func atoi(s string) int {
	n := 0
	for _, r := range s {
		n = n*10 + int(r-'0')
	}
	return n
}
