package day2

import (
	"fmt"
	"log"
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

func part1(lines []string) {
	allowedCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	idSum := 0
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

func part2(lines []string) {
	sumPowers := int64(0)

	for _, line := range lines {
		game := strings.Split(line, ":")
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
			roundPower := minCubes["red"] * minCubes["blue"] * minCubes["green"]
			sumPowers = sumPowers + int64(roundPower)
		}
	}

	fmt.Println(sumPowers)
}
