package main

import (
	"advent/main/day1"
	"advent/main/day2"
	"advent/main/day3"
	"advent/main/day4"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type dayFunction func(input []string, num int)
const INPUT_LOCATION = "inputs"

var dayFunctions = map[int]dayFunction{
	1: day1.Run,
	2: day2.Run,
	3: day3.Run,
	4: day4.Run,
	// "5": day5.Run,
}

func main() {
	start := time.Now()
	day, err := strconv.Atoi(os.Args[1])
	part, err2 := strconv.Atoi(os.Args[2])

	if err != nil || err2 != nil {
		panic("Invalid argument format should be <day number> <part number>")
	}

	if dayFunc, ok := dayFunctions[day]; ok {
		fmt.Printf("Running Day %d, Part %d\n", day, part)
		fmt.Println("------------------------------")
		dayFunc(getInputFile(day), part)
		fmt.Println("------------------------------")
	} else {
		fmt.Println("No function defined for this day")
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}


func getInputFile(day int) []string {

	path := fmt.Sprintf("%s/input%d.txt", INPUT_LOCATION, day)
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}