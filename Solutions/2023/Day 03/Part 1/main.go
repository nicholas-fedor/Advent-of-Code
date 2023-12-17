package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Schematic represents the rows of cells within a schematic
type Schematic struct {
	Rows []Row
}

// Row represents a line of cells within the schematic
type Row struct {
	Cells []Cell
}

// Cell represents each rune in a line within a schematic
type Cell struct {
	Value      string // Value holds the string value of the cell
	Type       string // Type specifies if the cell contains a digit or symbol
	NearSymbol bool   // NearSymbol indicates if the cell is near a symbol
}

func main() {
	// Input filepath
	filePath := "input.txt"

	// Opens the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Instantiates the schematic
	var schematic Schematic

	// Instantiates the scanner object with the file
	scanner := bufio.NewScanner(file)

	// Enumerates through the file and maps it accordingly
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

	// Enumerates through the schematic and evaluates for digits that are near symbols
	for y, row := range schematic.Rows {
		for x, cell := range row.Cells {
			if cell.Type == "digit" {
				schematic.Rows[y].Cells[x].NearSymbol = isNearSymbol(y, x, schematic)
			}
		}
	}

	// Enumerate through the schematic and evaluate for cells that are part numbers
	var numbers [][]Cell
	var number []Cell
	for _, row := range schematic.Rows {
		for _, cell := range row.Cells {
			switch cell.Type {
			case "digit":
				number = append(number, cell)
			default:
				if isPartNumber(number) {
					numbers = append(numbers, number)
				}
				number = []Cell{}
			}
		}
		// EOL handling
		if isPartNumber(number) {
			numbers = append(numbers, number)
		}
		number = []Cell{}
	}

	// Enumerate through the cells in the numbers array and extract the values
	var partNumbersStr [][]string
	var partNumberStr []string
	for _, number := range numbers {
		for _, cell := range number {
			partNumberStr = append(partNumberStr, cell.Value)
		}
		partNumbersStr = append(partNumbersStr, partNumberStr)
		partNumberStr = []string{}
	}

	// Convert the array of part numbers from strings to integers
	partNumbersInt, err := convertToNumbers(partNumbersStr)
	if err != nil {
		panic(err)
	}

	// Calculate the sum of the part numbers
	numSum := calculateSum(partNumbersInt)

	// Output the sum of the part numbers
	fmt.Printf("Sum of Part Numbers: %d\n", numSum)
}

// getType evaluates if a rune is a number or a symbol
func getType(r rune) string {
	switch {
	case unicode.IsDigit(r): // rune is a digit 0-9
		return "digit"
	case !unicode.IsDigit(r) && r != '.': // rune is neither a digit nor a period
		return "symbol"
	}
	return ""
}

// isNearSymbol evaluates if a cell is adjacent to another cell that contains a symbol
func isNearSymbol(y, x int, schematic Schematic) bool {
	maxY := len(schematic.Rows)
	maxX := len(schematic.Rows[0].Cells)

	// Enumerate cells starting at the top-right
	// Move horizontally from left to right
	// And move vertically from top to bottom
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Skips the center cell
			if i == 0 && j == 0 {
				continue
			}
			seekY, seekX := y+i, x+j
			// Conditional to ensure handling within boundaries
			if seekY >= 0 && seekY < maxY && seekX >= 0 && seekX < maxX && schematic.Rows[seekY].Cells[seekX].Type == "symbol" {
				return true
			}
		}
	}
	return false
}

// isPartNumber evaluates if a slice of cells contains any cells that are near a symbol
func isPartNumber(number []Cell) bool {
	if len(number) > 0 {
		for _, cell := range number {
			if cell.NearSymbol {
				return true
			}
		}
	}
	return false
}

// convertToNumbers converts an matrix of string slices containing numbers into integers
func convertToNumbers(numbersStringMatrix [][]string) ([]int, error) {
	var numbersIntSlice []int
	for _, digit := range numbersStringMatrix {
		numberString := strings.Join(digit, "")
		numberInt, err := strconv.Atoi(numberString)
		if err != nil {
			return nil, err
		}
		numbersIntSlice = append(numbersIntSlice, numberInt)
	}
	return numbersIntSlice, nil
}

// calculateSum calculates the sum of a slice of integers
func calculateSum(partNumbersInt []int) int {
	var sum int
	for _, n := range partNumbersInt {
		sum += n
	}
	return sum
}
