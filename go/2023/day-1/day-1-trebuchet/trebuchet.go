package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func firstAndLastDigit(text string) int {
	firstNumberFound := false
	fmt.Println(text)
	var currentVal = []rune{'0', '0'}
	var stringValue string
	for _, r := range text {
		if unicode.IsDigit(r) {
			if !firstNumberFound {
				firstNumberFound = true
				currentVal[0] = r
			}
			currentVal[1] = r
		}
	}
	stringValue = string(currentVal[0])
	stringValue += string(currentVal[1])
	numberValue, err := strconv.Atoi(stringValue)
	check(err)
	return numberValue
}

func hasDigit(text string) bool {
	for _, r := range text {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func main() {
	readFile, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var calibrationValues []int
	sumOfCalibrationValues := 0
	for fileScanner.Scan() {
		currentValue := firstAndLastDigit(fileScanner.Text())
		calibrationValues = append(calibrationValues, currentValue)
	}
	for i := 0; i < len(calibrationValues); i++ {
		sumOfCalibrationValues += calibrationValues[i]
	}
	fmt.Println("calibration values:", calibrationValues)
	fmt.Println("sum of calibration values:", sumOfCalibrationValues)
}
