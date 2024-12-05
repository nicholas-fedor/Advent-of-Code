package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
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

func parseFile(file *os.File) ([][]int, error) {
	var numbers2DArray [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		numbers, err := getNumbers(line)
		if err != nil {
			return nil, err
		}
		numbers2DArray = append(numbers2DArray, numbers)
	}
	return numbers2DArray, nil
}

func getNumbers(line []byte) ([]int, error) {
	var numbers []int

	// Split the line into separate byte arrays delimited by a blank space
	numbersByteArray := bytes.Split(line, []byte(" "))
	// Iterate through the newly split array
	for _, numberBytes := range numbersByteArray {
		// Convert the bytes into strings and then integers
		// This first joins them together before the type conversion
		number, err := strconv.Atoi(string(numberBytes))
		if err != nil {
			return nil, err
		}
		// Add the result to the output
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func getDifferences(data [][]int) [][]int {
	var differences [][]int

	// Iterate through the data array of reports
	for _, report := range data {
		// Calculate the differences and append the result to the differences array
		differences = append(differences, getReportDifferences(report))
	}
	return differences
}

func getReportDifferences(report []int) []int {
	var reportDifferences []int

	// Evaluate a single line of the data array
	for i := 0; i < len(report)-1; i++ {
		reportDifferences = append(reportDifferences, report[i+1]-report[i])
	}

	return reportDifferences
}

func getSafeReportsCount(differences [][]int) int {
	var safeReportsCount int

	for _, report := range differences {
		if isSafe(report) {
			safeReportsCount++
		}
	}

	return safeReportsCount
}

func isSafe(report []int) bool {
	if (isAllPositive(report) || isAllNegative(report)) && isWithinRange(report) {
		return true
	}
	return false
}

func isAllPositive(differences []int) bool {
	var count int
	// Evaluate if each element in the array
	for _, difference := range differences {
		// If the element meets the critera, then increment the counter
		if difference > 0 {
			count++
		}
	}
	// If the count is equal to the number of elements in the array, then all of the elements meet the criteria
	if count == len(differences) {
		return true
	}
	return false
}

func isAllNegative(differences []int) bool {
	var count int
	// Evaluate if each element in the array
	for _, difference := range differences {
		// If the element meets the critera, then increment the counter
		if difference < 0 {
			count++
		}
	}
	// If the count is equal to the number of elements in the array, then all of the elements meet the criteria
	if count == len(differences) {
		return true
	}
	return false
}

func isWithinRange(differences []int) bool {
	var count int
	// Evaluate if each element in the array
	for _, difference := range differences {
		// If the element meets the critera, then increment the counter
		if (difference >= 1 && difference <= 3) || (difference <= -1 && difference >= -3) {
			count++
		}
	}
	// If the count is equal to the number of elements in the array, then all of the elements meet the criteria
	if count == len(differences) {
		return true
	}

	return false
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
	data, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	// Evaluate the data to calculate the differences between the n and n+1 elements
	differences := getDifferences(data)
	fmt.Println(differences)

	// Evaluate the differences to determine if all the elements in the row meet the criteria of being "safe"
	safeReportsCount := getSafeReportsCount(differences)

	// Output the final tally of "safe reports"
	fmt.Printf("There are %d safe reports\n", safeReportsCount)
}
