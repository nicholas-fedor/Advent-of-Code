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
	filePath := "sample.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a slice to store the data
	var data [][]Cell

	// Instantiate y index for data
	var yIndex int

	// Store data from file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read each line as a byte array
		line := scanner.Text()

		// Create a new row to store the cells
		var row []Cell

		// Loop through line
		for xIndex, b := range line {
			// Evaluate each cell.
			cell := Cell{
				Value:  string(b),
				X:      xIndex,
				Y:      yIndex,
				Digit:  isDigit(b),
				Symbol: isSymbol(b),
			}
			// Add the cells to the row.
			row = append(row, cell)
		}
		// Append the row to the data array.
		data = append(data, row)

		// Increment y index
		yIndex++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Analyze data and update if a digit neighbors a symbol.
	for y, line := range data {
		for x, cell := range line {
			// Only use origin cells that are digits.
			if cell.Digit {
				// Evaluate cell if neighboring cells containing a symbol and
				// update the data array
				data[y][x].NearSymbol = isNearSymbol(y, x, data)
			}
		}
	}

	// Find valid number series.
	var partNumbersStr [][]string
	var number []string
	for _, line := range data {
		for _, cell := range line {
			// Initial condition for iterating over cells, thus ensuring not
			// outputting empty slices.
			switch {
			case cell.Digit:
				// TODO: Handle joining together digits for numbers that are
				// near symbols.
				number = append(number, cell.Value)
			case !cell.Digit:
				if len(number) > 0 {
					partNumbersStr = append(partNumbersStr, number)
					number = []string{}
				}
			}
		}
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

	// Output
	fmt.Println(partNumbersInt)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isSymbol(r rune) bool {
	return r != '.' && r != ' ' && !isDigit(r)
}

// Evaluate if a given cell neighbors a cell with a symbol.
func isNearSymbol(y, x int, data [][]Cell) bool {
	maxY := len(data)
	maxX := len(data[0])

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue // Skip the current cell
			}

			newY, newX := y+i, x+j
			if newY >= 0 && newY < maxY && newX >= 0 && newX < maxX && data[newY][newX].Symbol {
				return true
			}
		}
	}
	return false
}
