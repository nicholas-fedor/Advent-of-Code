package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Cell struct {
	Value      string
	X          int
	Y          int
	Digit      bool
	Symbol     bool
	NearSymbol bool
}

func main() {
	// Define the file path
	filePath := "input.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a slice to store the data
	var data [][]Cell

	// Store data from file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read each line as a byte array
		line := scanner.Text()

		// Create a new row to store the cells
		var row []Cell

		// Loop through line
		for _, b := range line {
			// Evaluate each cell.
			cell := Cell{
				Value:  string(b),
				Digit:  isDigit(b),
				Symbol: isSymbol(b),
			}
			// Add the cells to the row.
			row = append(row, cell)
		}
		// Append the row to the data array.
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Analyze data and update if a digit neighbors a symbol.
	for y, line := range data {
		for x, cell := range line {
			// Only use origin cells that are digits.
			if cell.Digit {
				// Evaluate cell if neighboring cells contain a symbol
				data[y][x].NearSymbol = isNearSymbol(y, x, data)
			}
		}
	}

	// Iterating through data (again) for digits and to build out numbers.
	var numbersCell [][]Cell
	var numberCell []Cell
	for _, line := range data {
		for _, cell := range line {
			// Initial condition for iterating over cells, thus ensuring not
			// outputting empty slices.
			switch {
			case cell.Digit:
				numberCell = append(numberCell, cell)
			case !cell.Digit:
				if len(numberCell) > 0 && isValid(numberCell) {
					numbersCell = append(numbersCell, numberCell)
				}
				numberCell = []Cell{}
			}
		}
	}

	var partNumbersStr [][]string
	var partNumber []string
	for _, cells := range numbersCell {
		for _, number := range cells {
			partNumber = append(partNumber, number.Value)
		}
		partNumbersStr = append(partNumbersStr, partNumber)
		partNumber = []string{}
	}

	// [][]string to []int conversion
	var partNumbersInt []int
	for _, slice := range partNumbersStr {
		str := strings.Join(slice, "")
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		partNumbersInt = append(partNumbersInt, num)
	}

	// Sum
	var partNumSum int
	for _, number := range partNumbersInt {
		partNumSum += number
	}

	// Output
	fmt.Println("Part Numbers:", partNumbersInt)
	fmt.Println("Sum of Part Numbers:", partNumSum)
}


func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isSymbol(r rune) bool {
	return r != '.' && r != ' ' && !isDigit(r)
}

// isNearSymbol evaluates if a cell neighbors a cell with a symbol.
func isNearSymbol(y, x int, data [][]Cell) bool {
	maxY := len(data)
	maxX := len(data[0])

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue // Skip the current cell
			}
			// Starts at top-left diagonal and moves left-to-right across while
			// iterating downwards.
			newY, newX := y+i, x+j
			// Boundary checks and then checks if cell value is a symbol.
			if newY >= 0 && newY < maxY && newX >= 0 && newX < maxX && data[newY][newX].Symbol {
				return true
			}
		}
	}
	return false
}

// isValid evaluates if the numberCell contains any cells.NearSymbol
func isValid(number []Cell) bool {
	for _, cell := range number {
		if cell.NearSymbol {
			return true
		}
	}
	return false
}