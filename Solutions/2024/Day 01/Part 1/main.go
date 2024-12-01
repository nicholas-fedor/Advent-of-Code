package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	sort.Ints(col1)
	sort.Ints(col2)

	return col1, col2, nil
}

func getDistances(col1, col2 []int) ([]int, error) {
	distances := make([]int, 0)

	for i, x := range col1 {
		distance := col2[i] - x
		distances = append(distances, distance)
	}

	return distances, nil
}

func getTotalDistance(distances []int) int {
	var totalDistance int
	for _, x := range distances {
		totalDistance += int(math.Abs(float64(x)))
	}
	return totalDistance
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

	distances, err := getDistances(col1, col2)
	if err != nil {
		log.Fatalln(err)
	}

	totalDistance := getTotalDistance(distances)
	fmt.Println(totalDistance)
}
