/*
Copyright © 2024 Nicholas Fedor

This is a simple web scraper for Eric Wastl's Advent of Code project.
https://adventofcode.com

The advent calendar is always from December 1st to the 25th.
Oldest event: 2015
Most recent: 2024
*/

// scrapeHTML obtains all of the data from the Advent of Code and organizes it
// into a local website.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Gets the HTML for the specified URL and writes to respective file.
func extractAndSaveSection(url, filePath string) error {
	// Obtains response object from web request.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Creates GoQuery document object with content from response object.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	// Returns object matching content filtered by selector.
	articleContent := doc.Find("html body main article")

	// Creates file to save htmlContent.
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Returns the HTML string from articleContent.
	htmlContent, err := articleContent.Html()
	if err != nil {
		return err
	}

	// Writes HTML to the file.
	_, err = file.WriteString(htmlContent)
	if err != nil {
		return err
	}

	// Returns nil if function runs successfully.
	return nil
}

func main() {
	var wg sync.WaitGroup

	// Data directory.
	dir := "../Data/HTML"

	// Creates data directory if not present.
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating directory: %s\n", err)
		}
	}

	// Loops from 2015 to current year.
	currentYear := time.Now().Year()
	for year := 2015; year <= currentYear; year++ {
		// Create directory for respective year, if not already present.
		yearDir := fmt.Sprintf("%s/%d", dir, year)
		if _, err := os.Stat(yearDir); os.IsNotExist(err) {
			err := os.Mkdir(yearDir, os.ModePerm)
			if err != nil {
				log.Fatalf("Error creating directory: %s\n", err)
			}
		}

		// Buffer to hold the links to each respective day.
		var links strings.Builder

		// Sets the last day to 25 or the current day if the year is the current year.
		endDay := 25
		if year == currentYear {
			endDay = time.Now().Day()
		}

		// Loops from day 1 to the last day.
		for day := 1; day <= endDay; day++ {
			// Creates filePath string for respective year and day.
			// Format: "[data dir]/[year]/[year]-[day].html"
			// Example: "./Web/2023/2023-01.html"
			filePath := fmt.Sprintf("%s/%d/%d-%02d.html", dir, year, year, day)

			// Loop that generates HTML for respective year - day, if it does
			// not already exist.
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				wg.Add(1)
				go func(y, d int) {
					defer wg.Done()
					// Request URL.
					url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", y, d)
					// Executes request and file write.
					err := extractAndSaveSection(url, filePath)
					if err != nil {
						log.Printf("Error fetching %d Day %d: %s\n", y, d, err)
					}
				}(year, day)
			}

			// Generates link for respective day to write to the respective
			// year's index.html file.
			links.WriteString(fmt.Sprintf(`<li><a href="%d-%02d.html">Day %02d</a></li>`, year, day, day))
		}

		// Content for respective year's index.html file.
		indexContent := fmt.Sprintf(`<!DOCTYPE html>
		<html>
		<head>
			<title>Advent of Code Prompts</title>
		</head>
		<body>
			<h1>Advent of Code Prompts - %d</h1>
			<ul>
			 <li><a href="../index.html">Home</a></li>
				%s
			</ul>
		</body>
		</html>`, year, links.String())

		// Respective year's index.html filepath.
		indexFilePath := fmt.Sprintf("%s/%d/index.html", dir, year)

		// Creates the respective year's index.html file.
		indexFile, err := os.Create(indexFilePath)
		if err != nil {
			log.Fatalf("Error creating index file for year %d: %s\n", year, err)
		}
		defer indexFile.Close()

		// Writes content to respective year's index.html file.
		_, err = indexFile.WriteString(indexContent)
		if err != nil {
			log.Fatalf("Error writing content to index file for year %d: %s\n", year, err)
		}
	}

	// Buffer to hold URL links for each year's index.html file location.
	// Used in generating the primary index.html file.
	var yearLinks strings.Builder

	// Generates year links for primary index.html file.
	for year := 2015; year <= time.Now().Year(); year++ {
		yearLinks.WriteString(fmt.Sprintf(`<li><a href="%d/index.html">%d</a></li>`, year, year))
	}

	// Generates the content for the primary index.html file.
	mainContent := fmt.Sprintf(`<!DOCTYPE html>
	<html>
	<head>
		<title>Advent of Code Prompts</title>
	</head>
	<body>
		<h1>Advent of Code Prompts</h1>
		<ul>
			%s
		</ul>
	</body>
	</html>`, yearLinks.String())

	// File path of main index.html file that contains links to each year.
	mainFilePath := fmt.Sprintf("%s/index.html", dir)

	// Creates main index.html file.
	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		log.Fatalf("Error creating main index file: %s\n", err)
	}
	defer mainFile.Close()

	// Writes content to main index.html file.
	_, err = mainFile.WriteString(mainContent)
	if err != nil {
		log.Fatalf("Error writing content to main index file: %s\n", err)
	}

	wg.Wait()
}
