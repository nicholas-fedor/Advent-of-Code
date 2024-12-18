// Scrape Advent of Code for the Inputs
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// getSecrets loads the user agent and cookie from the secrets.env file.
func getSecrets() (string, string) {
	err := godotenv.Load("secrets.env")
	if err != nil {
		log.Fatal(err)
	}
	UserAgent := os.Getenv("USER_AGENT")
	Cookie := os.Getenv("COOKIE")

	return UserAgent, Cookie
}

// inputFileExists checks if the input file already exists for a given year and day.
func inputFileExists(year, day int) bool {
	filePath := fmt.Sprintf("../Data/Input/%d/%02d/input.txt", year, day)
	if _, err := os.Stat(filePath); err == nil {
		return true // File exists
	}
	return false
}

// getInput fetches the content for a given year and day from the Advent of Code website.
func getInput(year, day int) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	UserAgent, Cookie := getSecrets()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", Cookie))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status: %s", resp.Status)
	}

	return string(body), nil
}

func main() {
	// Loops from 2015 to current year (2024).
	// If current year, then only outputs to current day.
	// Otherwise, outputs days 1 to 25.
	startYear := 2015
	currentYear := time.Now().Year()

	startDay := 1
	endDay := 25
	currentDay := time.Now().Day()

	var wg sync.WaitGroup

	// Loops through years starting at startYear through currentYear.
	for year := startYear; year <= currentYear; year++ {
		yearDir := fmt.Sprintf("../Data/Inputs/%d", year)
		err := os.MkdirAll(yearDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating directory for year %d: %s\n", year, err)
		}

		// If year evaluated is the currentYear, then the endDay will be the
		// currentDay. This avoids attempted scraping of nonexistent days.
		if year == currentYear {
			endDay = currentDay
		}

		// Loops through days starting at 1 until either endDay or currentDay.
		for day := startDay; day <= endDay; day++ {
			// If no input.txt file exists in the directory tree, then the
			// following code is executed.
			if !inputFileExists(year, day) {
				wg.Add(1)
				go func(y, d int) {
					defer wg.Done()
					input, err := getInput(y, d)
					if err != nil {
						log.Printf("Error getting input for %d/%02d: %s\n", y, d, err)
						return
					}

					dayDir := fmt.Sprintf("../Data/Inputs/%d/%02d", y, d)
					err = os.MkdirAll(dayDir, os.ModePerm)
					if err != nil {
						log.Fatalf("Error creating directory for %d/%02d: %s\n", y, d, err)
						return
					}

					filePath := fmt.Sprintf("%s/input.txt", dayDir)
					file, err := os.Create(filePath)
					if err != nil {
						log.Fatalf("Error creating file for %d/%02d: %s\n", y, d, err)
						return
					}
					defer file.Close()

					_, err = file.WriteString(input)
					if err != nil {
						log.Fatalf("Error writing content for %d/%02d: %s\n", y, d, err)
					}
				}(year, day)
			}
		}
	}

	wg.Wait()
}
