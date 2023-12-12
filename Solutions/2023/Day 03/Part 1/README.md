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

Looking for sum of all "part numbers" that are adjacent to a symbol.
Any number adjacent to a symbol, even diagonally, is a "part number" and should be included in the sum.
Periods do not count as symbols.

Mapping of the input data to account for (x,y) position will allow for easier
handling of business logic.

1) Map digits and symbols (except periods)
2) Iterate left-to-right to identify number series, i.e. 2+ digits.
3) Validate if "part number" by evaluating if index positions surrounding the
   number series have symbols.

i.e. Index from left-to-right and top-to-bottom.

### Example Table Conversion

```console
9x9 Table Position [x,y] Mapping
[0,0][1,1][2,1][3,1][4,1][5,1][6,1][7,1][8,1][9,1]
[0,1][1,1][2,1][3,1][4,1][5,1][6,1][7,1][8,1][9,1]
[0,2][1,2][2,2][3,2][4,2][5,2][6,2][7,2][8,2][9,2]
[0,3][1,3][2,3][3,3][4,3][5,3][6,3][7,3][8,3][9,3]
[0,4][1,4][2,4][3,4][4,4][5,4][6,4][7,4][8,4][9,4]
[0,5][1,5][2,5][3,5][4,5][5,5][6,5][7,5][8,5][9,5]
[0,6][1,6][2,6][3,6][4,6][5,6][6,6][7,6][8,6][9,6]
[0,7][1,7][2,7][3,7][4,7][5,7][6,7][7,7][8,7][9,7]
[0,8][1,8][2,8][3,8][4,8][5,8][6,8][7,8][8,8][9,8]
[0,9][1,9][2,9][3,9][4,9][5,9][6,9][7,9][8,9][9,9]

Example Data
[4][6][7][.][.][1][1][4][.][.]
[.][.][.][*][.][.][.][.][.][.]
[.][.][3][5][.][.][6][3][3][.]
[.][.][.][.][.][.][#][.][.][.]
[6][1][7][*][.][.][.][.][.][.]
[.][.][.][.][.][+][.][5][8][.]
[.][.][5][9][2][.][.][.][.][.]
[.][.][.][.][.][.][7][5][5][.]
[.][.][.][$][.][*][.][.][.][.]
[.][6][6][4][.][5][9][8][.][.]

Transformed to indicate d for digit and s for symbol.
[d][d][d][.][.][d][d][d][.][.]
[.][.][.][s][.][.][.][.][.][.]
[.][.][d][d][.][.][d][d][d][.]
[.][.][.][.][.][.][s][.][.][.]
[d][d][d][s][.][.][.][.][.][.]
[.][.][.][.][.][s][.][d][d][.]
[.][.][d][d][d][.][.][.][.][.]
[.][.][.][.][.][.][d][d][d][.]
[.][.][.][s][.][s][.][.][.][.]
[.][d][d][d][.][d][d][d][.][.]

Further transformation to indicate y for yes, if d = adjacent to s.
[y][y][y][.][.][n][n][n][.][.]
[.][.][.][s][.][.][.][.][.][.]
[.][.][y][y][.][.][y][y][y][.]
[.][.][.][.][.][.][s][.][.][.]
[y][y][y][s][.][.][.][.][.][.]
[.][.][.][.][.][s][.][n][n][.]
[.][.][y][y][y][.][.][.][.][.]
[.][.][.][.][.][.][y][y][y][.]
[.][.][.][s][.][s][.][.][.][.]
[.][y][y][y][.][y][y][y][.][.]
```

Iterate through for digits that are yes and split by symbol or space (period) to
get "Part Numbers":
467

35 633

617

592
755

644 598

Answer = 467 + 35 + 633 + 617 + 592 + 755 + 644 + 598

## Solution

Total Sum of Part Numbers: 539,590
