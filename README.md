# Advent of Code 2023

[https://adventofcode.com/](https://adventofcode.com/)

> Advent of Code is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like. People use them as interview prep, company training, university coursework, practice problems, a speed contest, or to challenge each other.

## File Structure

The main.go file contains the main function which calls the day functions. Each day function calls the part functions. The part functions contain the code for the solution.

to run the code for a specific day and part, run the following command:

```bash
go run main.go  <day> <part>
```

where day is the day number and part is either 1 or 2.

## File Template

this is a sample template for day 3. Copy and paste this into a new file and replace the day number with the correct day number.
text inputs are stored in the inputs folder. The input file name is the day number with the .txt extension.

```Go
package main

import (
 "fmt"
)

func day3(part int) {
 fmt.Println("Day 3")
 switch part {
 case 1:
  day3part1()
 case 2:
  day3part2()
 default:
  fmt.Println("No function defined for this part")
 }
}

func day3part1() {

}

func day3part2() {

}

```
