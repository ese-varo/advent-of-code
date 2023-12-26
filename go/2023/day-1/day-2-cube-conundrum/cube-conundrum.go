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
	colorsRegex := regexp.MustCompile(`blue|red|green`)
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

func main() {
	var validGameIds []int
	sumOfValidGameIds := 0
	readFile, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if isValidGame(text) {
			currentGameId := getGameId(text)
			validGameIds = append(validGameIds, currentGameId)
		}
	}

	for i := 0; i < len(validGameIds); i++ {
		sumOfValidGameIds += validGameIds[i]
	}
	p("Sum of valid games ids:", sumOfValidGameIds)
}
