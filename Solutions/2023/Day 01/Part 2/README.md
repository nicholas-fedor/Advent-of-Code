# Advent of Code: Day 1 - Part 2

## Solution
54676

## Process
Ultimately used code from
<https://github.com/TheN00bBuilder/AOC2023/blob/master/d1/p2/aocday1p2.go> to
solve.

main.go.bak has my attempted code, which got a solution that was close, but not
fully accurate due to edge cases, such as when there were substrings of "twone".

My attempted solution was to use Regex to identify the matching substring;
however, was unable to get it to function for the right-most desired line number
digit value in those edge-cases.

The other solutions have proposed doing cleanup of the string to account for
these issues.

I'll revisit later, as spent about 1-2 days working on this.

I did use ChatGPT to help with my attempt; however, even it couldn't figure out
how to account for the edge cases (no surprise).

Other attempts at solutions can be found here: <https://www.reddit.com/r/adventofcode/comments/1883ibu/2023_day_1_solutions/>
