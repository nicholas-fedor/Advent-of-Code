package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func getColumns(file *os.File) ([]int, []int, error) {
	col1 := make([]int, 0)
	col2 := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		firstElement := strings.Split(line, "   ")[0]
		firstNum, err := strconv.Atoi(firstElement)
		if err != nil {
			return nil, nil, err
		}

		secondElement := strings.Split(line, "   ")[1]
		secondNum, err := strconv.Atoi(secondElement)
		if err != nil {
			return nil, nil, err
		}

		col1 = append(col1, firstNum)
		col2 = append(col2, secondNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return col1, col2, nil
}

func getSimilarity(col1, col2 []int) []int {
	var similarityArr []int

	var count int
	for _, x := range col1 {
		for _, y := range col2 {
			if x == y {
				count++

			}
			if x != y {
				continue
			}
		}

		similarityArr = append(similarityArr, count * x)
		count = 0
	}

	return similarityArr
}

func getTotalSimilarity(similarity []int) int {
	var totalSimilarity int

	for _, x := range similarity {
		totalSimilarity += x
	}

	return totalSimilarity
}

func main() {
	// file, err := openFile(SampleFile)
	file, err := openFile(InputFile)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	col1, col2, err := getColumns(file)
	if err != nil {
		log.Fatalln(err)
	}

	similarity := getSimilarity(col1, col2)

	totalSimilarity := getTotalSimilarity(similarity)
	fmt.Println(totalSimilarity)
}
