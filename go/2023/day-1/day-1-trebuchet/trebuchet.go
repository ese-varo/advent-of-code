package main

import (
	"fmt"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	calVal := "treb7uchet"
	firstNumberFound := false
	fmt.Println(calVal)
	var currentVal = []int{0, 0}
	for _, r := range calVal {
		if unicode.IsDigit(r) {
			if !firstNumberFound {
				firstNumberFound = true
				currentVal[0] = int(r)
			}
			currentVal[1] = int(r)
		}
	}

	fmt.Printf("%c", currentVal)
}

// readFile, err := os.Open("input.txt")
// check(err)

// fileScanner := bufio.NewScanner(readFile)

// fileScanner.Split(bufio.ScanLines)

// var count int = 0
// for fileScanner.Scan() {
// 	count++
// 	fmt.Println("Line #", count, ":", fileScanner.Text())
// }
// fileScanner.Scan()
// var calVal string = fileScanner.Text()
