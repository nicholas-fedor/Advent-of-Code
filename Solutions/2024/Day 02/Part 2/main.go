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

func getAllPermutations(rows [][]int, n int) [][][]int {
	result := make([][][]int, len(rows))

	for i, row := range rows {
		result[i] = make([][]int, 0)
		indices := make([]int, n)
		generatePermutationsRecursive(row, indices, 0, n, len(row), &result[i])
	}
	return result
}

func generatePermutationsRecursive(row []int, indices []int, index, n, length int, result *[][]int) {
	if index == n {
		// Generate permutation by excluding the indices
		perm := make([]int, 0, length-n)
		j := 0
		for i := 0; i < length; i++ {
			if j < n && i == indices[j] {
				j++
				continue
			}
			perm = append(perm, row[i])
		}
		*result = append(*result, perm)
		return
	}

	var start int
	switch {
	case index == 0:
		// First index, start from 0
		start = 0
	default:
		// Subsequent indices, start from last chosen index + 1
		start = indices[index-1] + 1
	}

	for i := start; i < length-(n-index-1); i++ {
		indices[index] = i
		generatePermutationsRecursive(row, indices, index+1, n, length, result)
	}
}

func getAllDifferences(allPermutations [][][]int) [][][]int {
	var allDifferences [][][]int

	for _, allReportPermutations := range allPermutations {
		var reportDifferences [][]int
		for _, permutation := range allReportPermutations {
			reportDifferences = append(reportDifferences, getReportDifferences(permutation))
		}
		allDifferences = append(allDifferences, reportDifferences)
	}

	return allDifferences
}

func getReportDifferences(permutation []int) []int {
	var reportDifferences []int

	for i := 0; i < len(permutation)-1; i++ {
		reportDifferences = append(reportDifferences, permutation[i+1]-permutation[i])
	}

	return reportDifferences
}

func getSafeReportsCount(allDifferences [][][]int) int {
	var safeReportsCount int

	for _, row := range allDifferences {
		var safePermutationCount int
		for _, permutation := range row {
			if isSafe(permutation) {
				safePermutationCount++
			}
		}
		if safePermutationCount > 0 {
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
		// If the element meets the criteria, then increment the counter
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
		// If the element meets the criteria, then increment the counter
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
		// If the element meets the criteria, then increment the counter
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

	numFieldsOmitted := 1
	allPermutations := getAllPermutations(data, numFieldsOmitted)

	allDifferences := getAllDifferences(allPermutations)

	safeReportsCount := getSafeReportsCount(allDifferences)

	fmt.Printf("There are %d safe reports\n", safeReportsCount)
}
