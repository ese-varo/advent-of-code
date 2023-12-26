package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	s "strings"
)

var p = fmt.Println
var findDigit = regexp.MustCompile(`\d+`)
var colorsRegex = regexp.MustCompile(`blue|red|green`)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getGameId(text string) int {
	gameId, _ := strconv.Atoi(findDigit.FindString(text))
	return gameId
}

func isValidSet(set string) bool {
	cubesNumber, _ := strconv.Atoi(findDigit.FindString(set))
	switch colorsRegex.FindString(set) {
	case "red":
		if cubesNumber > 12 {
			return false
		}
	case "green":
		if cubesNumber > 13 {
			return false
		}
	case "blue":
		if cubesNumber > 14 {
			return false
		}
	}
	return true
}

func isValidGame(text string) bool {
	setsString := s.Split(text, ":")[1]
	sets := s.Split(setsString, ";")
	for i := 0; i < len(sets); i++ {
		set := s.Split(sets[i], ",")
		for j := 0; j < len(set); j++ {
			if !isValidSet(set[j]) {
				return false
			}
		}
	}
	return true
}

func getMinimumCubes(sets []string) map[string]int {
	var minimumCubes = map[string]int{"red": 0, "green": 0, "blue": 0}
	for i := 0; i < len(sets); i++ {
		set := s.Split(sets[i], ",")
		for j := 0; j < len(set); j++ {
			cubesNumber, _ := strconv.Atoi(findDigit.FindString(set[j]))
			color := colorsRegex.FindString(set[j])
			if cubesNumber > minimumCubes[color] {
				minimumCubes[color] = cubesNumber
			}
		}
	}
	return minimumCubes
}

func calculatePower(game string) int {
	setsString := s.Split(game, ":")[1]
	sets := s.Split(setsString, ";")
	minimumCubes := getMinimumCubes(sets)
	power := minimumCubes["red"] * minimumCubes["green"] * minimumCubes["blue"]
	return power
}

func main() {
	var gamePowers []int
	sumOfPowers := 0
	readFile, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		gamePower := calculatePower(text)
		gamePowers = append(gamePowers, gamePower)
	}

	for i := 0; i < len(gamePowers); i++ {
		sumOfPowers += gamePowers[i]
	}
	p("Sum of powers:", sumOfPowers)
}
