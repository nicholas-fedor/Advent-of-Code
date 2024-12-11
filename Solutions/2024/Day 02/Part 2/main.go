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
	// Debug output
	// fmt.Println(data)

	// TODO: Implement logic to omit x levels and then determine if the report is safe.

	// Thought - expand data into variants based on if x levels are omitted.
	// i.e. [1 2 3 4 5] with a single omitted level expands into the following:
	// [2 3 4 5], [1 3 4 5], [1 2 4 5], [1 2 3 5], [1 2 3 4]
	// In the problem, the criteria is still as follows:
	// - The levels are either all increasing or decreasing
	// - Any two adjacent levels differ by at least one and as most three
	// As an example, report #3 = [9 7 6 2 1]
	// If omitting a single level, then that expands to the following:
	// a) [7 6 2 1]
	// b) [9 6 2 1]
	// c) [9 7 2 1]
	// d) [9 7 6 1]
	// e) [9 7 6 2]
	// With the following differences (left to right)
	// a) [-1 -4 -1]
	// b) [-3 -4 -1]
	// c) [-2 -5 -1]
	// d) [-2 -1 -5]
	// e) [-2 -1 -4]
	// When evaluated based on the criteria:
	// a) [-1 -4 -1] = -4 is greater than -3 = unsafe
	// b) [-3 -4 -1] = -4 is greater than -3 = unsafe
	// c) [-2 -5 -1] = -5 is greater than -3 = unsafe
	// d) [-2 -1 -5] = -5 is greater than -3 = unsafe
	// e) [-2 -1 -4] = -4 is greater than -3 = unsafe
	// If there were any of the variations that were safe, then we could +1 a safeReportSet counter.

	numFieldsOmitted := 1
	allPermutations := getAllPermutations(data, numFieldsOmitted)
	// Debug output
	for _, reportPermutations := range allPermutations {
		fmt.Println(reportPermutations)
	}
}
