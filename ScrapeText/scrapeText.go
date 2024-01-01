/*
Copyright © 2023 Nicholas Fedor

This is a simple web scraper for Eric Wastl's Advent of Code project.
https://adventofcode.com

The advent calendar is always from December 1st to the 25th.
Oldest event: 2015
Most recent: 2023

The scraped data will be saved as text files.
Directory structure is as follows:
* Advent of Code
** Prompts
*** [Year]
**** [Day]
** Scraper

*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
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

// promptFileExists checks if the prompt file already exists for a given year and day.
func promptFileExists(year, day int) bool {
	filePath := fmt.Sprintf("../Data/Text/%d/%02d/prompt.txt", year, day)
	if _, err := os.Stat(filePath); err == nil {
		return true // File exists
	}
	return false
}

// getPrompt fetches the content for a given year and day from the Advent of Code website.
func getPrompt(year, day int) (string, string, string, error) {
	var header, prompt string
	var err error
	var url string

	// Adds values from secrets.env
	UserAgent, Cookie := getSecrets()

	// Creates Colly Collector
	c := colly.NewCollector(
		colly.AllowedDomains("adventofcode.com"),
	)

	// Find and store the content for the specified day.
	c.OnHTML(".day-desc", func(h *colly.HTMLElement) {

		// Outputs header first: ex. "--- Day 1: Sonar Sweep ---"
		header = h.ChildText("h2")

		// Outputs remaining text.
		prompt = strings.TrimPrefix(h.Text, header)
	})

	// Sets the cookie and user-agent for the request header.
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("user-agent", UserAgent)
		r.Headers.Set("cookie", Cookie)
	})

	// Error handling for request.
	c.OnError(func(r *colly.Response, e error) {
		err = e
	})

	// Example: https://adventofcode.com/2021/day/1
	url = fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)

	// Executes request
	c.Visit(url)

	return url, header, prompt, err
}

// Combines content from response for saving.
func createFileContent(url, header, prompt string) string {
	fileContent := fmt.Sprintf("URL: %s\n\n%s\n\n%s", url, header, prompt)
	return fileContent
}

func main() {
	// Loops from 2015 to current year (2023).
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
		yearDir := fmt.Sprintf("../Data/Text/%d", year)
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
			// If no prompt.txt file exists in the directory tree, then the
			// following code is executed.
			if !promptFileExists(year, day) {
				wg.Add(1)
				go func(y, d int) {
					defer wg.Done()
					url, header, prompt, err := getPrompt(y, d)
					if err != nil {
						log.Printf("Error getting prompt for %d day %d: %s\n", y, d, err)
						return
					}

					dayDir := fmt.Sprintf("../Data/Text/%d/%02d", y, d)
					err = os.MkdirAll(dayDir, os.ModePerm)
					if err != nil {
						log.Fatalf("Error creating directory for %d day %d: %s\n", y, d, err)
						return
					}

					filePath := fmt.Sprintf("%s/prompt.txt", dayDir)
					file, err := os.Create(filePath)
					if err != nil {
						log.Fatalf("Error creating file for %d day %d: %s\n", y, d, err)
						return
					}
					defer file.Close()

					fileContent := createFileContent(url, header, prompt)
					_, err = file.WriteString(fileContent)
					if err != nil {
						log.Fatalf("Error writing content for %d day %d: %s\n", y, d, err)
					}
				}(year, day)
			}
		}
	}

	wg.Wait()
}
