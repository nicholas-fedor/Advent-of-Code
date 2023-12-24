// TODO: Determine best method for obtaining coordinate pairs.
// Currently not filtering second coordinate from known coordinate pairs prior
// to applying as coord2.

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
	Value string   // Value holds the string value of the cell
	Type  CellType // Type specifies if the cell contains a digit or symbol
}

type CellType string

const (
	Digit CellType = "digit"
	Gear  CellType = "gear"
)

func main() {
	// Input filepath
	filepath := "input.txt"

	// Opens the file
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Returns a schematic object from the file
	schematic, err := getSchematic(file)
	if err != nil {
		panic(err)
	}

	// Returns the part numbers from the schematic
	partNumbers, err := getPartNumbers(*schematic)
	if err != nil {
		panic(err)
	}

	// Returns the gear ratios (i.e. the product of part number pairs)
	gearRatios := getGearRatios(partNumbers)

	// Returns the total sum of the gear ratios
	totalSum := getTotalSum(gearRatios)

	// Outputs the total sum
	fmt.Println("Total Sum of Gear Ratios:", totalSum)
}

// getTotalSum returns the total sum of all gear ratios
func getTotalSum(gearRatios []int) int {
	var totalSum int
	for _, gearRatio := range gearRatios {
		// fmt.Println(gearRatio)
		totalSum += gearRatio
	}
	return totalSum
}

// getGearRatios returns the product of part number pairs
func getGearRatios(partNumbers []int) []int {
	var gearRatios []int
	for x := 0; x < len(partNumbers)-1; x += 2 {
		gearRatios = append(gearRatios, partNumbers[x]*partNumbers[x+1])
	}
	return gearRatios
}

// getPartNumbers evaluates the schematic for all part numbers
func getPartNumbers(schematic Schematic) ([]int, error) {
	// Returns all of the coordinate pairs for the first digits of part numbers
	coordinatePairs := getCoordinatePairs(schematic)

	// Returns all of the part numbers as strings
	var partNumbersStr [][]string
	for _, coordinatePair := range coordinatePairs { // All coordinate pairs
		for _, coordinate := range coordinatePair { // Coordinate pairs for a gear
			y, x := coordinate[0], coordinate[1]
			partNumbersStr = append(partNumbersStr, getPartNumber(x, y, schematic))
		}
	}

	// Converts the strings into integers
	partNumbersInt, err := convertToNumbers(partNumbersStr)
	if err != nil {
		return nil, err
	}
	return partNumbersInt, nil
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

// getPartNumber returns a slice of string values consisting of digits
func getPartNumber(x, y int, schematic Schematic) []string {
	var partNumber []string
	for isWithinBounds(x, y, schematic) && isDigit(x, y, schematic) {
		partNumber = append(partNumber, schematic.Rows[y].Cells[x].Value)
		x++
	}
	return partNumber
}

// isWithinBounds evaluates if x and y are within the schematic
func isWithinBounds(x, y int, schematic Schematic) bool {
	maxY := len(schematic.Rows)
	maxX := len(schematic.Rows[0].Cells)
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}

// getCoordinate returns the coordinates for a the first digit in a number
func getCoordinate(x, y int, schematic Schematic) []int {
	var coordinate []int
	for isWithinBounds(x, y, schematic) && isDigit(x, y, schematic) {
		coordinate = []int{y, x}
		x-- // Seek left
	}
	return coordinate
}

// getCoordinatePair returns a pair of coordinates for a specific gear
func getCoordinatePair(x, y int, schematic Schematic) [][]int {
	var coordinatePair [][]int
	var coordinate []int
	for j := -1; j <= 1; j++ { // Row iterator
		for i := -1; i <= 1; i++ { // Column iterator
			evalY, evalX := j+y, i+x
			if j == 0 && i == 0 {
				continue
			}
			coordinate = getCoordinate(evalX, evalY, schematic)
			if len(coordinate) == 2 {
				coordinatePair = append(coordinatePair, coordinate)
				fmt.Println(coordinatePair)
			}
		}
	}
	return coordinatePair
}

// getCoordinatePairs returns the coordinates for the first digits of part
// numbers that are paired to gears.
func getCoordinatePairs(schematic Schematic) [][][]int {
	var coordinatePairs [][][]int // all coordinate pairs
	var coordinatePair [][]int
	for y, row := range schematic.Rows {
		for x, cell := range row.Cells {
			if cell.Type == Gear {
				coordinatePair = getCoordinatePair(x, y, schematic)
				if len(coordinatePair) == 2 {
					coordinatePairs = append(coordinatePairs, coordinatePair)
				}
			}
		}
	}
	return coordinatePairs
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
