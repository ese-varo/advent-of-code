package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var mappedNumbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9"}
var hasNumber, _ = regexp.Compile(`one|two|three|four|five|six|seven|eight|nine|\d`)
var p = fmt.Println

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findFirstAndLastDigit(text string) int {
	firstDigit := findNumber(text, false)
	lastDigit := findNumber(text, true)
	twoDigitsNumber, _ := strconv.Atoi(string(firstDigit + lastDigit))
	p(twoDigitsNumber)
	return twoDigitsNumber
}

func findNumber(text string, startFromEnd bool) string {
	number := ""
	i := 1
	for len(number) == 0 {
		if startFromEnd {
			number = hasNumber.FindString(text[len(text)-i:])
		} else {
			number = hasNumber.FindString(text[0:i])
		}
		i++
	}
	if len(number) > 1 {
		number = mappedNumbers[number]
	}
	return number
}

func main() {
	var calibrationValues []int
	sumOfCalibrationValues := 0

	readFile, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		currentValue := findFirstAndLastDigit(text)
		calibrationValues = append(calibrationValues, currentValue)
	}
	for i := 0; i < len(calibrationValues); i++ {
		sumOfCalibrationValues += calibrationValues[i]
	}
	p("calibration values:", calibrationValues)
	p("sum of calibration values:", sumOfCalibrationValues)
}
