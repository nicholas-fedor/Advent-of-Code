# Advent of Code: Day 3 - Part 1

## Prompt
--- Day 3: Gear Ratios ---

You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

```console
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
```

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?

## Analysis

- Enumerate through the file, line by line, and load it into memory.
- While enumerating, evaluate if a rune is a digit or a symbol.
- Then, enumerate through the digits and evaluate if it neighbors a symbol.
- Enumerate line by line for digits, appending them to a "number" slice
  variable.
  - Evaluate if the number slice contains any digits that neighbor symbols.
  - If yes, then append that number to an array of part numbers.
- Enumerate through the array of part numbers and add them together for a total sum.

## Solution

Total Sum of Part Numbers: 539,590

## Learning Review

As a relative newbie to programming, this was a good challenge.
I will fully admit that I used ChatGPT to help me with deriving a solution to
this; however, it still took me some time to get this nailed down.

The biggest hurdle was finally narrowing down an edge-case that was resulting in
my output sum being too large.
In this issue, a number at the end of a row was not being terminated correctly,
as the proceeding row started with a digit.
This was ultimately a simple fix, but was difficult to pinpoint until I was able
to pinpoint the single number (6-digits vs the 3-digit average).