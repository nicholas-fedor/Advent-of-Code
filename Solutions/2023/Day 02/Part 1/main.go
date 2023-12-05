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
	// Input file
	Filename = "input"

	// Filter criteria
	RedMax   = 12
	GreenMax = 13
	BlueMax  = 14
)

// Game struct maps each game to an ID and a map of Colors to an int value.
type Game struct {
	ID     int
	Colors map[string]int
}

func main() {
	// File handling
	file, err := os.Open(Filename)
	check(err)
	defer file.Close()

	// Parsing of file data into usable array of data.
	games := parseGameData(file)

	// Total of summarized Game ID values.
	var gameIDSum int

	// Iterates through game data and filters based on above criteria.
	for _, game := range games {
		// If game meets criteria, then the game's ID is added to the sum.
		if isGameFiltered(game) {
			gameIDSum += game.ID
		}
	}
	// Final output of summarized Game ID totals after filtering.
	fmt.Println("Filtered GameID Total:", gameIDSum)
}

// isGameFiltered returns boolean based on selected criteria.
func isGameFiltered(game Game) bool {
	return game.Colors["red"] <= RedMax &&
		game.Colors["green"] <= GreenMax &&
		game.Colors["blue"] <= BlueMax
}

// parseGameData takes raw file data and returns data for each game, including
// the mapped maximums observed for each color across the samples.
func parseGameData(file *os.File) []Game {
	var gamesDataArray []Game
	gameRegex := regexp.MustCompile(`Game (\d+):(.+)`)
	scanner := bufio.NewScanner(file)
	var currentGame Game

	for scanner.Scan() {
		line := scanner.Text()

		// Regex filter condition to execute subsequent business logic.
		if gameRegex.MatchString(line) {

			// Slice of strings from leftmost match of Regex.
			matches := gameRegex.FindStringSubmatch(line)

			// Type conversion of GameID obtained from matches[1].
			gameID, _ := strconv.Atoi(matches[1])

			// Sets value in currentGame.ID to gameID.
			currentGame.ID = gameID

			// Returns maxCounts for respective colors obtained from matches[2].
			maxCounts := extractMaxCounts(matches[2])

			// Sets data for currentGame.Colors to maxCounts.
			currentGame.Colors = maxCounts

			// Appends currentGame data (ID and Colors) to games.
			gamesDataArray = append(gamesDataArray, currentGame)

			// Resets currentGame data prior to further loop repeat.
			currentGame = Game{}
		}
	}

	// Error handling of scanner.
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return gamesDataArray
}

// extractMaxCounts returns the maxCounts for each respective color across the
// samples for a game.
func extractMaxCounts(gameSamplesInfo string) map[string]int {

	// Returns an empty map array for [colors]maxValues.
	maxCounts := make(map[string]int)

	// Returns an array of the data for each sample within a game.
	samples := strings.Split(gameSamplesInfo, ";")

	// Iterates through each sample.
	for _, col := range samples {

		// Returns string array of the color and count info from each respective
		// sample.
		sampleDataArray := strings.Split(strings.TrimSpace(col), ",")

		// Loops through the sample to return respective data.
		for _, sampleDataStrings := range sampleDataArray {
			sampleData := strings.Split(strings.TrimSpace(sampleDataStrings), " ")
			sampleColorCount, _ := strconv.Atoi(sampleData[0])
			sampleColorName := sampleData[1]

			// Sets the count for respective color in the sample to the array
			// value if the sampleCount value is greater than the value already
			// present in the array.
			// i.e. returns maximum value for respective color across the
			// samples for the respective game.
			if sampleColorCount > maxCounts[sampleColorName] {
				maxCounts[sampleColorName] = sampleColorCount
			}
		}
	}

	return maxCounts
}

// check handles errors via log.Fatalln(err)
func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
