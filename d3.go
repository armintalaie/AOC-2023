package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func day3(part int) {
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

	file, err := os.Open("inputs/input3.txt")

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

	var reg = regexp.MustCompile("(\\d)")
	for i, line := range lines {
		startNumIndex := -1
		endNumIndex := -1
		for j, char := range line {
			if reg.MatchString(string(char)) {
				if startNumIndex == -1 {
					startNumIndex = j
				} else {
					endNumIndex = j
				}
			} else {

				applyCase(startNumIndex, endNumIndex, &sum, lines, i)
				startNumIndex = -1
				endNumIndex = -1

			}

		}
		applyCase(startNumIndex, endNumIndex, &sum, lines, i)

	}
	fmt.Println(sum)
}

func applyCase(startNumIndex int, endNumIndex int, sum *int, lines []string, lineIndex int) {
	if startNumIndex != -1 {
		if endNumIndex == -1 {
			endNumIndex = startNumIndex
		}
		start, err := strconv.Atoi(lines[lineIndex][startNumIndex : endNumIndex+1])
		// fmt.Println(startNumIndex, endNumIndex, start)
		if err != nil {
			log.Fatal(err)
		}
		if isSuitable(startNumIndex, endNumIndex, lineIndex, lines) {
			// fmt.Println(startNumIndex, endNumIndex, start)
			*sum += start
		}
	}

}

func isSuitable(start int, end int, row int, lines []string) bool {
	var reg = regexp.MustCompile("[^\\d.]")
	// testString := "*"

	// fmt.Println(reg.MatchString(testString))

	for i := max(0, start-1); i <= min(end+1, len(lines[0])-1); i++ {
		for j := max(0, row-1); j <= min(row+1, len(lines)-1); j++ {
			if reg.MatchString(string(lines[j][i])) {
				return true
			}
		}
	}
	return false
}

func day3part2() {

	file, err := os.Open("inputs/input3.txt")

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

	var reg = regexp.MustCompile("[\\*]")
	for i, line := range lines {
		for j, char := range line {
			if reg.MatchString(string(char)) {
				gearRatio := findGearRatio(lines, j, i)
				fmt.Println(gearRatio)
				sum += gearRatio
			}
		}
	}
	fmt.Println(sum)

}

func findGearRatio(lines []string, index int, row int) int {
	var reg = regexp.MustCompile("[\\d]")
	nums := make([]int, 0)

	for i := max(0, row-1); i <= min(row+1, len(lines)-1); i++ {
		for j := max(0, index-1); j <= min(index+1, len(lines[0])-1); j++ {
			if reg.MatchString(string(lines[i][j])) {
				fmt.Println(i, j, lines[i][j], lines[i], index, row)

				findNumber(lines, j, i, &nums)
			}
		}
	}

	prod := 1
	fmt.Print(nums)

	if len(nums) == 0 {
		return 0
	} else {
		for _, num := range nums {
			prod *= num
		}
	}
	return prod
}

func findNumber(lines []string, index int, row int, nums *[]int) {
	var reg = regexp.MustCompile("[\\d]")
	for i := max(0, row-1); i <= min(row+1, len(lines)-1); i++ {
		for j := max(0, index-1); j <= min(index+1, len(lines[0])-1); j++ {
			if reg.MatchString(string(lines[j][i])) {
				*nums = append(*nums, findNumberHelper(lines, j, i))
			}
		}
	}
}

func findNumberHelper(lines []string, index int, row int) int {
	var reg = regexp.MustCompile("(\\d)")
	startNumIndex := -1
	endNumIndex := -1
	for j, char := range lines[row] {
		if reg.MatchString(string(char)) {
			if startNumIndex == -1 {
				startNumIndex = j
			} else {
				endNumIndex = j
			}
		} else {

			if startNumIndex != -1 {
				if endNumIndex == -1 {
					endNumIndex = startNumIndex
				}
				start, err := strconv.Atoi(lines[row][startNumIndex : endNumIndex+1])
				// fmt.Println(startNumIndex, endNumIndex, start)
				if err != nil {
					log.Fatal(err)
				}
				return start
			}
			startNumIndex = -1
			endNumIndex = -1

		}

	}
	if startNumIndex != -1 {
		if endNumIndex == -1 {
			endNumIndex = startNumIndex
		}
		start, err := strconv.Atoi(lines[row][startNumIndex : endNumIndex+1])
		// fmt.Println(startNumIndex, endNumIndex, start)
		if err != nil {
			log.Fatal(err)
		}
		return start
	}
	return -1
}
