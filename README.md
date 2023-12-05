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

this is a sample template for day 365. Copy and paste this into a new folder e.g. `day365` and create a `main.go` file.
text inputs are stored in the inputs folder. The input file name is the day number with the .txt extension.
In the root dir\'s `main.go` add the new package and the entry for the function mapper.

```Go
package day365

import (
 "fmt"
)

func Run(input []string, part int) {
 switch part {
 case 1:
  part1(input)
 case 2:
  part2(input)
 default:
  fmt.Println("No function defined for this part")
 }
}

func part1(lines []string) {}
func part2(lines []string) {}

```
