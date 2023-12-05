# Advent of Code: Day 2 - Part 1

## Prompt

You have a small bag and some red, green, and blue cubes.

Each round of the game involves being shown a sample of the cubes from the bag
before they are returned to it.

You are shown multiple samples of the bag's contents throughout the game.

## Puzzle Input

* Recorded data from playing multiple games.
* Semicolons separate sample groups
* Sample groups may be ordered differently in the raw data

### Format
Game [ID]: n blue, n red, n green; n red, n blue, n green; n red, n blue, n green

#### Example Data
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

##### Interpretation
In game 1, three sets of cubes are revealed from the bag (and then put back
again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red
cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

## Challenge Problem

The Elf would first like to know which games would have been possible if the bag
contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had
been loaded with that configuration. However, game 3 would have been impossible
because at one point the Elf showed you 20 red cubes at once; similarly, game 4
would also have been impossible because the Elf showed you 15 blue cubes at
once. If you add up the IDs of the games that would have been possible, you get
8.

Determine which games would have been possible if the bag had been loaded with
only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs
of those games?

## Analysis
Data is for 100 games total.

Task is to filter out the games that have samples indicating more than the
indicated n cubes for each color: 12 x Red, 13 x Green, 14 x Blue.

Raw data needs to be transformed prior to being filtered.
Each game row has x number of samples.
Not every game has the same number of samples.
The colors per sample are not ordered.

Last step is to take the ID# of the games that could have been possible and then
add them together.

## Solution
3035

The code to solve this iterates through the raw data and uses Regular
Expressions to transform the information into usable data that can be filtered
by the prescribed requirements.
