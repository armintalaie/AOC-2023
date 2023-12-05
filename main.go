package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type dayFunction func(num int)

var dayFunctions = map[string]dayFunction{
	"1": day1,
	"2": day2,
	"3": day3,
	"4": day4,
}

func main() {
	start := time.Now()
	day := os.Args[1]
	part, _ := strconv.Atoi(os.Args[2])

	if dayFunc, ok := dayFunctions[day]; ok {
		fmt.Printf("Running day %s\t Part %d\n", day, part)
		fmt.Println("------------------------------")
		dayFunc(part)
		fmt.Println("------------------------------")
	} else {
		fmt.Println("No function defined for this day")
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
