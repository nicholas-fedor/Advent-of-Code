package main

import (
	"bufio"
	"os"
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
	Value     string // Value holds the string value of the cell
	Type      string // Type specifies if the cell contains a digit or symbol
	NearDigit bool   // NearGear indicates if the cell is near a gear (*)
}

func main() {
	// Input filepath
	filepath := "sample.txt"

	// Opens the file
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Instantiates the scanner object with the file
	scanner := bufio.NewScanner(file)

	// Instantiates the schematic object to contain the file contents
	var schematic Schematic

	// Enumerates through the file and maps it to the schematic object
	for scanner.Scan() {
		line := scanner.Text()
		var row Row
		for _, b := range line {
			cell := Cell{
				Value: string(b),
				Type:  getType(b),
			}
			row.Cells = append(row.Cells, cell)
		}
		schematic.Rows = append(schematic.Rows, row)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Enumerates through the schematic and evaluates for gears that are near
	// digits.
	for y, row := range schematic.Rows {
		for x, cell := range row.Cells {
			if cell.Type == "gear" && isNearDigit(y, x, schematic) {
				// Evaluates for part numbers.
				// TODO: Write a function to generate a matrix of part number
				// pairs to be multiplied.
				// Final desired outcome is a sum of those multiplied part numbers.
			}
		}
	}

}

// getType evaluates if a rune is a digit or a gear
func getType(b rune) string {
	switch {
	case unicode.IsDigit(b):
		return "digit"
	case b == '*':
		return "gear"
	}
	return ""
}

// isNearGear evaluates if a cell is adjacent to a digit
func isNearDigit(y, x int, schematic Schematic) bool {
	maxY := len(schematic.Rows)
	maxX := len(schematic.Rows[0].Cells)

	// Enumerates cells starting from the top-right
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			evalY, evalX := y+i, x+j
			if evalY >= 0 && evalY < maxY && evalX >= 0 && evalX < maxX && schematic.Rows[evalY].Cells[evalX].Type == "digit" {
				return true
			}
		}
	}

	return false
}
