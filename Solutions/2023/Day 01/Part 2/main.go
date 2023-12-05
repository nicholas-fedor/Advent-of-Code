// Code from https://github.com/TheN00bBuilder/AOC2023/blob/master/d1/p2/aocday1p2.go
package main

// don't judge a man based off his imports
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	handle, error := os.Open("test")
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(handle)
	lineNumberString := ""
	documentSumTotal := 0
	spelledNumbersArray := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		line := scanner.Text()

		// Debugging: Prints the original line.
		fmt.Println("Original Line:", line)

		lineNumberString = ""
		// for each element in the number string array
		for counted, element := range spelledNumbersArray {
			// get index
			substringIndex := strings.Index(line, element)

			fmt.Println("Substring Index:", substringIndex)

			// ok ok ok ok ok hear me out
			for substringIndex != -1 {
				// here's where it gets funky:
				// insert whichever one is found INSIDE the number we're working on currently
				// this prevents it from being re-detected forever in a loop
				// because we don't care about what it looks like, only what's in the string
				line = line[:substringIndex+1] + strconv.Itoa(counted) + line[substringIndex+1:]

				// Debugging: Prints the updated line.
				fmt.Println("Updated Line:", line)

				// then go onto the next match
				substringIndex = strings.Index(line, element)

				// Debugging: Prints the next matched substringIndex
				fmt.Println("Next Match Substring Index:", substringIndex)
			}
		}
		// then we use the same bit as before
		for _, char := range line {
			if unicode.IsNumber(char) {
				lineNumberString = lineNumberString + string(char)
			}
		}
		lineNumberString = string(lineNumberString[0]) + string(lineNumberString[len(lineNumberString)-1])
		lineNumberInt, _ := strconv.Atoi(lineNumberString)

		// Debugging: Print the lineNumberInt
		fmt.Printf("Line Number: %d\n\n", lineNumberInt)

		documentSumTotal = documentSumTotal + lineNumberInt
	}
	fmt.Print(documentSumTotal)
}
