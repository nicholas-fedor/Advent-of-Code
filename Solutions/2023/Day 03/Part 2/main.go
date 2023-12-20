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
	Value string // Value holds the string value of the cell
	Type  string // Type specifies if the cell contains a digit or symbol
}

func main() {
	// Input filepath
	filepath := "input.txt"

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

	var cps [][][]int // all coordinate pairs
	var cp [][]int    // coordinate pairs for a gear
	var nc []int      // coordinate of left-most number
	for y, row := range schematic.Rows {
		for x, cell := range row.Cells {
			if cell.Type == "gear" {
				for j := -1; j <= 1; j++ {
					for i := -1; i <= 1; i++ {

						// Sets starting eval position to -1,-1
						evalY, evalX := j+y, i+x

						// Skips center cell
						if j == y && i == x {
							continue
						}

						// Looks for in-bounds cells with digits
						for isWithinBounds(evalX, evalY, schematic) && isDigit(evalX, evalY, schematic) {
							nc = []int{evalY, evalX}
							evalX--
						}
					}
					if len(nc) > 0 {
						cp = append(cp, nc)
					}
					nc = []int{}
				}
				if len(cp) > 1 {
					cps = append(cps, cp)
				}
				cp = [][]int{}
			}
		}
	}

	var partNumbersStr [][]string
	for _, pair := range cps {
		for _, coordinate := range pair {
			y, x := coordinate[0], coordinate[1]
			partNumbersStr = append(partNumbersStr, getPartNumber(y, x, schematic))
		}
	}

	partNumbersInt, err := convertToNumbers(partNumbersStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part Numbers:", partNumbersInt)

	var total int
	var product int
	var num1, num2 int

	// partNumLength := len(partNumbersInt)
	// fmt.Println("# of Part Numbers", partNumLength)
	fmt.Println()

	for x, number := range partNumbersInt {
		switch x % 2 {
		case 0:
			num1 = number
			// fmt.Printf("Part Number %d: %d\n", x+1, num1)
		case 1:
			num2 = number
			// fmt.Printf("Part Number %d: %d\n", x+1, num2)

			product = num1 * num2
			// fmt.Println("Gear Ratio:", product)
			// fmt.Println()
			total += product
		}
		
	}

	fmt.Println("Total Sum of Gear Ratios:", total)
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

func getPartNumber(y, x int, schematic Schematic) []string {
	var partNumber []string
	for isWithinBounds(x, y, schematic) && isDigit(x, y, schematic) {
		partNumber = append(partNumber, schematic.Rows[y].Cells[x].Value)
		x++
	}
	return partNumber
}

func isDigit(x, y int, schematic Schematic) bool {
	return schematic.Rows[y].Cells[x].Type == "digit"
}

func isWithinBounds(x, y int, schematic Schematic) bool {
	maxY := len(schematic.Rows)
	maxX := len(schematic.Rows[0].Cells)

	return x >= 0 && x < maxX && y >= 0 && y < maxY
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
