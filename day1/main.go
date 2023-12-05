package day1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func convertToDigit(num string) int {
	curr, _ := strconv.Atoi(num)
	return curr
}

func replaceWordsWithNumbers(inputString string) string {
	wordToNumber := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	for num, replacement := range wordToNumber {
		inputString = strings.ReplaceAll(inputString, num, replacement)
	}

	return inputString
}

func part1(lines []string) {
	re, err := regexp.Compile("\\d{1}")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	for _, line := range lines {
		line = replaceWordsWithNumbers(line)
		values := re.FindAllString(line, -1)
		curr := convertToDigit(values[0] + values[len(values)-1])
		sum = sum + curr
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	re, err := regexp.Compile("\\d{1}")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for _, line := range lines {
		line = replaceWordsWithNumbers(line)
		values := re.FindAllString(line, -1)
		curr := convertToDigit(values[0] + values[len(values)-1])
		sum = sum + curr
	}
	fmt.Println(sum)
}
