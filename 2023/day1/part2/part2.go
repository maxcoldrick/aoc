package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// the one1one format is to account for oneight or twone .. which I found on the internet
var wordToDigitMapping = map[string]string{"one": "one1one", "two": "two2two", "three": "three3three", "four": "four4four",
	"five": "five5five", "six": "six6six", "seven": "seven7seven", "eight": "eight8eight", "nine": "nine9nine",
}

//go:embed input
var inputFile string

func main() {
	lines := strings.Split(inputFile, "\n")

	calibrationValueSum := 0

	for _, line := range lines {
		for word, digit := range wordToDigitMapping {
			//if string contains the word, replace
			if strings.Contains(line, word) {
				line = strings.Replace(line, word, digit, -1)
			}
		}

		//  9986fmfqhdmq8 becomes (0=9986, 1=8)
		onlyDigits := keepOnlyDigitsFrom(line)

		// join the slice together so (0=9986, 1=8) becomes "99868"
		digitsAsString := strings.Join(onlyDigits, "")

		if len(digitsAsString) == 1 {
			// if it's an individual digit then duplicate it, convert to an int and add it to the final sum
			calibrationValueSum = calibrationValueSum + duplicateSingleInt(onlyDigits)
		} else {
			// if it's multiple digits then take the first and last, convert to an int and add it to the final sum
			calibrationValueSum = calibrationValueSum + stringToInt(firstAndLastFrom(digitsAsString))
		}
	}
	fmt.Println(calibrationValueSum)
}

func firstAndLastFrom(input string) string {
	firstAndLastAsStringSlice := regexp.MustCompile("^.|.$").FindAllString(input, -1)
	return strings.Join(firstAndLastAsStringSlice, "")
}

func keepOnlyDigitsFrom(line string) []string {
	expression := regexp.MustCompile("[0-9]+")
	return expression.FindAllString(line, -1)
}

func duplicateSingleInt(input []string) int {
	var sb strings.Builder
	sb.WriteString(input[0])
	sb.WriteString(input[0])
	return stringToInt(sb.String())
}

func stringToInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}
