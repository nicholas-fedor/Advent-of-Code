# Advent of Code

<https://adventofcode.com/>

I am using this as a decent excuse to learn Golang.

## Solutions

Each day has two separate parts that are completed sequentially.
The scrapers only get the first parts and are missing the necessary input files.

I will attempt to figure out a decent format for including the prompts for both
parts in the solutions directory.

## Scraping Prompts and Inputs

The scrapers pull the initial prompts for each day into a Data directory;
They do not obtain the secondary prompts that are unlocked after completing part one of each day.

### ScrapeText -> main.go

Colly web-scraping that saves output as .txt files.
Outputs to "Data/Text" directory separated by year and day.

### ScrapeHTML -> main.go

GoQuery web-scraping that saves output at .html files.
Outputs to "Data/HTML" directory separated by year and day.
Generates index.html files for each year and a primary index.html file linking
to each year.

## Future Work

I may eventually get around to restructuring this repository to properly manage things.
