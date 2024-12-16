package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func parseFile(file *os.File) [][][]string {
	var matches [][][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches = append(matches, matchString(line))
	}

	return matches
}

func matchString(line string) [][]string {
	re := regexp.MustCompile("(?i)mul\\(\\d+,\\d+\\)")
	return re.FindAllStringSubmatch(line, -1)
}

func cleanupData(data [][][]string) [][]string {
	var numArray [][]string

	re := regexp.MustCompile("\\d+,\\d+")
	for _, x := range data {
		for _, y := range x {
			for _, z := range y {

				s := re.FindAllStringSubmatch(z, -1)
				for _, nums := range s {
					for _, num := range nums {
						ns := strings.Split(num, ",")
						numArray = append(numArray, ns)
					}
				}
			}
		}
	}

	return numArray
}

func convertData(data [][]string) [][]int {
	var nums [][]int

	for _, row := range data {
		var rowInt []int
		for _, y := range row {
			numInt, err := strconv.Atoi(y)
			if err != nil {
				log.Fatal(err)
			}
			rowInt = append(rowInt, numInt)
		}
		nums = append(nums, rowInt)
	}

	return nums
}

func multiplyGroup(nums [][]int) []int {
	var products []int

	for _, row := range nums {
		products = append(products, row[0] * row[1])
	}

	return products
}

func sumGroup(nums []int) int {
	var sum int

	for _, num := range nums {
		// fmt.Println(num)
		sum += num
	}

	return sum
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
	data := parseFile(file)

	cleanedData := cleanupData(data)

	numArray := convertData(cleanedData)

	productNums := multiplyGroup(numArray)

	fmt.Println(sumGroup(productNums))
}
