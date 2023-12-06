package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input
var inputFile string

func main() {
	lines := strings.Split(inputFile, "\n")
	sum := 0

	timeStart := time.Now()

	for _, line := range lines {
		gameNumber := 0
		gameValues := strings.FieldsFunc(line, Split)
		for i, value := range gameValues {
			if strings.Contains(value, "Game") {
				gameNumber = keepOnlyIntsFrom(value)
				gameValues = remove(gameValues, i)
				break
			}
		}

		if isPossible(gameValues) {
			sum = sum + gameNumber
		} else {
			fmt.Println("d")
		}
	}
	fmt.Println(sum)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func isPossible(gameValues []string) bool {
	isPossible := true
	for _, g := range gameValues {
		onlyInts := keepOnlyIntsFrom(g)
		if onlyInts >= 15 {
			return false
		}
		if strings.Contains(g, "blue") && onlyInts > 14 {
			isPossible = false
		} else if strings.Contains(g, "green") && onlyInts > 13 {
			isPossible = false
		} else if strings.Contains(g, "red") && onlyInts > 12 {
			isPossible = false
		}
	}
	return isPossible
}

func stringToInt(input []string) int {
	output, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Println(err)
	}
	return output
}

func Split(r rune) bool {
	return r == ':' || r == ',' || r == ';'
}

func keepOnlyIntsFrom(line string) int {
	expression := regexp.MustCompile("[0-9]+")
	return stringToInt(expression.FindAllString(line, -1))
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Time: 2.35ms...
