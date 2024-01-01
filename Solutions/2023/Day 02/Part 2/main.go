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
	Filename = "input.txt"
)

type Game struct {
	Colors map[string]int
}

func main() {
	// File handling.
	file, err := os.Open(Filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Returns data from file.
	gameData := parseGameData(file)

	// Solving for Part 2 solution.
	var sumProduct int
	for i, games := range gameData {
		fmt.Printf("Game: %d\n", i+1)
		var gameProduct int = 1
		for color, colorCount := range games.Colors {
			fmt.Printf("%s: %d\n", color, colorCount)
			gameProduct *= colorCount
		}
		fmt.Printf("Game Product: %d\n", gameProduct)
		fmt.Println(strings.Repeat("-", 15))
		sumProduct += gameProduct
	}
	fmt.Println("Total for all games:", sumProduct)
}

func parseGameData(file *os.File) []Game {
	var gamesDataArray []Game
	gameRegex := regexp.MustCompile(`Game (\d+):(.+)`)
	scanner := bufio.NewScanner(file)

	var currentGame Game
	for scanner.Scan() {
		line := scanner.Text()
		if gameRegex.MatchString(line) {
			matches := gameRegex.FindStringSubmatch(line)
			maxCounts := extractMaxCounts(matches[2])
			currentGame.Colors = maxCounts
			gamesDataArray = append(gamesDataArray, currentGame)
			currentGame = Game{}
		}
	}
	return gamesDataArray
}

func extractMaxCounts(gameSampleData string) map[string]int {
	maxCounts := make(map[string]int)
	samples := strings.Split(gameSampleData, ";")
	for _, sampleColor := range samples {
		sampleDataArray := strings.Split(strings.TrimSpace(sampleColor), ",")
		for _, sampleDataString := range sampleDataArray {
			sampleData := strings.Split(strings.TrimSpace(sampleDataString), " ")
			sampleColorCount, _ := strconv.Atoi(sampleData[0])
			sampleColorName := sampleData[1]

			if sampleColorCount > maxCounts[sampleColorName] {
				maxCounts[sampleColorName] = sampleColorCount
			}
		}
	}
	return maxCounts
}
