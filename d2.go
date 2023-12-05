package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2(part int) {
	switch part {
	case 1:
		day2part1()
	case 2:
		day2part2()
	default:
		fmt.Println("No function defined for this part")
	}
}

func day2part1() {
	file, err := os.Open("inputs/input2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	allowedCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	idSum := 0

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		game := strings.Split(line, ":")
		gameNumber := game[0][5:]
		gameNumberInt, err := strconv.Atoi(gameNumber)
		if err != nil {
			log.Fatal(err)
		}

		rounds := strings.Split(game[1], ";")

		allowedGame := true
		for _, round := range rounds {
			colors := strings.Split(round, ",")
			cubes := map[string]int{
				"red":   0,
				"blue":  0,
				"green": 0,
			}

			for _, color := range colors {
				parsed := strings.Split(strings.TrimSpace(color), " ")
				colorInt, _ := strconv.Atoi(parsed[0])
				cubes[parsed[1]] = colorInt

			}

			// fmt.Println(cubes)

			if cubes["red"] > allowedCubes["red"] || cubes["blue"] > allowedCubes["blue"] || cubes["green"] > allowedCubes["green"] {
				allowedGame = false
				break
			}

			if !allowedGame {
				break
			}
		}

		if allowedGame {
			idSum = idSum + gameNumberInt
		}
	}

	fmt.Println(idSum)

}

func day2part2() {
	file, err := os.Open("inputs/input2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sumPowers := int64(0)

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		game := strings.Split(line, ":")
		// gameNumber := game[0][5:]
		// gameNumberInt, err := strconv.Atoi(gameNumber)
		if err != nil {
			log.Fatal(err)
		}

		rounds := strings.Split(game[1], ";")

		minCubes := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		allowedGame := true
		for _, round := range rounds {
			colors := strings.Split(round, ",")
			cubes := map[string]int{
				"red":   0,
				"blue":  0,
				"green": 0,
			}

			for _, color := range colors {
				parsed := strings.Split(strings.TrimSpace(color), " ")
				colorInt, _ := strconv.Atoi(parsed[0])
				cubes[parsed[1]] = colorInt
			}

			minCubes["red"] = max(minCubes["red"], cubes["red"])
			minCubes["blue"] = max(minCubes["blue"], cubes["blue"])
			minCubes["green"] = max(minCubes["green"], cubes["green"])
		}

		if allowedGame {
			// fmt.Println(minCubes)
			roundPower := minCubes["red"] * minCubes["blue"] * minCubes["green"]

			sumPowers = sumPowers + int64(roundPower)
		}
	}

	fmt.Println(sumPowers)

}
