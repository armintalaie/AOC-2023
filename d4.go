package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func day4(part int) {
	switch part {
	case 1:
		day4part1()
	case 2:
		day4part2()
	default:
		fmt.Println("No function defined for this part")
	}
}

func day4part1() {

	file, err := os.Open("inputs/input4.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0

	var reg = regexp.MustCompile("(\\d+)")
	for _, line := range lines {
		numberPart := strings.Split(line, ":")[1]
		numberGroups := strings.Split(numberPart, "|")
		groupLeft := reg.FindAllString(strings.TrimSpace(numberGroups[0]), -1)
		elfNumbers := map[int]int{}
		groupRight := reg.FindAllString(strings.TrimSpace(numberGroups[1]), -1)

		reward := 0

		for _, numStr := range groupRight {
			num, _ := strconv.Atoi(string(numStr))
			elfNumbers[num]++
		}

		for _, numStr := range groupLeft {
			num, _ := strconv.Atoi(string(numStr))
			// fmt.Println(num, elfNumbers[num])
			if elfNumbers[num] > 0 {
				elfNumbers[num]-- // ?
				if reward >= 1 {
					reward *= 2
				} else {
					reward = 1
				}
			}
		}

		// fmt.Println(groupLeft, groupRight, reward)
		sum += reward

	}
	fmt.Println(sum)
}

func day4part2() {

	file, err := os.Open("inputs/input4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	copies := 0
	wonCopies := map[int]int{}

	var reg = regexp.MustCompile("(\\d+)")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		numberPart := strings.Split(line, ":")[1]
		numberGroups := strings.Split(numberPart, "|")
		groupLeft := reg.FindAllString(strings.TrimSpace(numberGroups[0]), -1)
		groupRight := reg.FindAllString(strings.TrimSpace(numberGroups[1]), -1)
		attempts := 1 + wonCopies[i]
		copies += attempts
		scratchcard(i, groupLeft, groupRight, &wonCopies, len(lines), attempts )
	}
	fmt.Println(copies)
}

func scratchcard(i int, groupLeft []string, groupRight []string, wonCopies *map[int]int, length int, attempts int) {
	elfNumbers := map[int]int{}
	count := 0

	for _, numStr := range groupRight {
		num, _ := strconv.Atoi(string(numStr))
		elfNumbers[num]++
	}

	for _, numStr := range groupLeft {
		num, _ := strconv.Atoi(string(numStr))
		if elfNumbers[num] > 0 {
			elfNumbers[num]--
			count++
		}
	}

	if count > 0 {
		for j := i + 1; j < min(i+count+1, length); j++ {
			(*wonCopies)[j]+= attempts
		}
	}

}
