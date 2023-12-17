package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Schematic struct {
	Rows []Row
}

type Row struct {
	Cells []Cell
}

type Cell struct {
	Value      string
	Type       string
	NearSymbol bool
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var schematic Schematic

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row Row
		for _, b := range line {
			cell := Cell{
				Value:      string(b),
				Type:       getType(b),
				NearSymbol: false,
			}
			row.Cells = append(row.Cells, cell)
		}
		schematic.Rows = append(schematic.Rows, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for y, row := range schematic.Rows {
		for x, cell := range row.Cells {
			if cell.Type == "digit" {
				schematic.Rows[y].Cells[x].NearSymbol = isNearSymbol(y, x, schematic)
			}
		}
	}

	for _, row := range schematic.Rows {
		for _, cell := range row.Cells {
			if cell.Type == "symbol" {

			}
		}
	}

	var numbers [][]Cell
	var number []Cell
	for _, row := range schematic.Rows {
		for _, cell := range row.Cells {
			switch cell.Type {
			case "digit":
				number = append(number, cell)
			default:
				if len(number) > 0 && isPartNumber(number) {
					numbers = append(numbers, number)
				}
				number = []Cell{}
			}
		}
		if len(number) > 0 && isPartNumber(number) {
			numbers = append(numbers, number)
		}
		number = []Cell{}
	}

	var partNumbers [][]string
	var partNumber []string
	for _, number := range numbers {
		for _, cell := range number {
			partNumber = append(partNumber, cell.Value)
		}
		partNumbers = append(partNumbers, partNumber)
		partNumber = []string{}
	}

	partNumbersInt := numStrToInt(partNumbers)
	fmt.Println(partNumbersInt)

	var numSum int
	for _, n := range partNumbersInt {
		numSum += n
	}
	fmt.Println(numSum)
}

func numStrToInt(partNumbers [][]string) []int {
	var numbersInt []int
	for _, v := range partNumbers {
		str := strings.Join(v, "")
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		numbersInt = append(numbersInt, num)
	}
	return numbersInt
}

func getType(r rune) string {
	switch {
	case unicode.IsDigit(r):
		return "digit"
	case !unicode.IsDigit(r) && r != '.':
		return "symbol"
	}
	return ""
}

func isNearSymbol(y, x int, schematic Schematic) bool {
	maxY := len(schematic.Rows)
	maxX := len(schematic.Rows[0].Cells)

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			seekY, seekX := y+i, x+j
			if seekY >= 0 && seekY < maxY && seekX >= 0 && seekX < maxX && schematic.Rows[seekY].Cells[seekX].Type == "symbol" {
				return true
			}
		}
	}
	return false
}

func isPartNumber(number []Cell) bool {
	for _, cell := range number {
		if cell.NearSymbol {
			return true
		}
	}
	return false
}
