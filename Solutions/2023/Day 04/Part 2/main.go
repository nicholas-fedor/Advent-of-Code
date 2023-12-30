package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const FilePath = "input.txt"

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
	Value string   // Value holds the string value of the cell
	Type  CellType // Type specifies if the cell contains a digit or symbol
}

// CellType represents the type of value within a cell
type CellType string

const (
	Digit CellType = "digit"
	Gear  CellType = "gear"
)

func main() {
	// Opens the file
	file, err := os.Open(FilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Returns a schematic object from the file
	schematic, err := getSchematic(file)
	if err != nil {
		panic(err)
	}

	// Returns all coordinates for the first digits of gear part numbers
	var allGearCoordinates [][][]int
	for y, row := range schematic.Rows {
		for x, cell := range row.Cells {
			if cell.Type == Gear {

				var gearNumberCoordinates [][]int
				// Starts iterating from [-1][-1], i.e. top, left and moves
				// first from left to right and then top to bottom.
				for j := -1; j <= 1; j++ {
					for i := -1; i <= 1; i++ {
						evalY, evalX := y+j, x+i

						// Skips the center i.e. the (*) gear symbol
						if j == 0 && i == 0 {
							continue
						}

						// Returns the coordinate for the first digit of a
						// possible part number
						var firstDigitCoordinate []int
						for isWithinBounds(evalX, evalY, *schematic) && isDigit(evalX, evalY, *schematic) {
							firstDigitCoordinate = []int{evalY, evalX}
							evalX-- // Ensures the loop continues until hitting the first digit while moving right to left
						}
						// Moves the returned coordinate to a slice matrix for
						// the respective gear
						if len(firstDigitCoordinate) > 0 {
							gearNumberCoordinates = append(gearNumberCoordinates, firstDigitCoordinate)
							firstDigitCoordinate = []int{}
						}
					}
				} // End loop for gear

				// Returns coordinates for possible part numbers when there is
				// more than one neighboring digit.
				if len(gearNumberCoordinates) > 1 {

					// First coordinate pair
					coordinate1 := gearNumberCoordinates[0]

					// Last coordinate pair (there are often more than two
					// coordinate pairs per gear, but we are only interested in
					// the first and last pairs.)
					coordinate2 := gearNumberCoordinates[len(gearNumberCoordinates)-1]
					gearCoordinate := [][]int{coordinate1, coordinate2}

					// Avoids duplicated coordinates
					if gearCoordinate[0][0] == gearCoordinate[1][0] && gearCoordinate[0][1] == gearCoordinate[1][1] {
						gearCoordinate = [][]int{}
						continue
					}

					// Places all the returned coordinates into a matrix
					allGearCoordinates = append(allGearCoordinates, gearCoordinate)
				}
			}
		}
	}

	// Returns the part number pairs for all gears
	var values [][]int
	var valueA, valueB []string
	for _, gearCoordinate := range allGearCoordinates {
		gearAY := gearCoordinate[0][0]
		gearAX := gearCoordinate[0][1]
		gearBY := gearCoordinate[1][0]
		gearBX := gearCoordinate[1][1]

		// Returns the first part number
		for isWithinBounds(gearAX, gearAY, *schematic) && isDigit(gearAX, gearAY, *schematic) {
			digit := schematic.Rows[gearAY].Cells[gearAX].Value
			valueA = append(valueA, digit)
			gearAX++
		}

		// Joins the string values together to form the number, returned as an integer
		intA, err := strconv.Atoi(strings.Join(valueA, ""))
		if err != nil {
			panic(err)
		}
		valueA = []string{}

		// Returns the second part number
		for isWithinBounds(gearBX, gearBY, *schematic) && isDigit(gearBX, gearBY, *schematic) {
			digit := schematic.Rows[gearBY].Cells[gearBX].Value
			valueB = append(valueB, digit)
			gearBX++
		}

		// Joins the string values together to form the number, returned as an integer
		intB, err := strconv.Atoi(strings.Join(valueB, ""))
		if err != nil {
			panic(err)
		}
		valueB = []string{}

		// Merges the two part numbers together into a slice matrix
		joinedValues := []int{intA, intB}
		values = append(values, joinedValues)
	}

	// Returns and outputs the following:
	// Gear: [gear number] | [part numbers]
	// Product: [product of part numbers] | SumTotal: [cumulative total of part number products]
	var sumTotal int
	for x, line := range values {
		x++
		fmt.Println("Gear:", x, "|", line)
		gearProduct := line[0] * line[1]
		sumTotal += gearProduct
		fmt.Println("Product:", gearProduct, "|", "SumTotal:", sumTotal)
	}
}

// isWithinBounds evaluates if x and y are within the schematic
func isWithinBounds(x, y int, schematic Schematic) bool {
	maxY := len(schematic.Rows)
	maxX := len(schematic.Rows[0].Cells)
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}

// isDigit evaluates if the cell within the schematic is a digit
func isDigit(x, y int, schematic Schematic) bool {
	return schematic.Rows[y].Cells[x].Type == Digit
}

// getType evaluates if a rune is a digit or a gear
func getType(b rune) CellType {
	switch {
	case unicode.IsDigit(b):
		return Digit
	case b == '*':
		return Gear
	}
	return ""
}

// getSchematic evaluates the input file and returns a schematic object
func getSchematic(file *os.File) (*Schematic, error) {
	// Instantiates the scanner object with the file
	scanner := bufio.NewScanner(file)

	// Instantiates the schematic object to contain the file contents
	schematic := &Schematic{}

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
		return nil, err
	}
	return schematic, nil
}
