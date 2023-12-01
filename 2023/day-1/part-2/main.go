package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numberWords = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)
	res := getInputValueSum(input)
	log.Println(res)
}

func getInputValueSum(input string) int {
	total := 0
	lines := strings.Split(strings.TrimRight(input, "\n\r"), "\n")

	for _, line := range lines {
		line = strings.TrimRight(line, "\n\r")
		firstDigit := getFirstDigit(line)
		lastDigit := getLastDigit(line)
		digit, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic(err)
		}

		total += digit
	}

	return total
}

func getFirstDigit(line string) string {
	soonestIndex := len(line)
	for i, c := range line {
		d := int(c - '0')
		if d <= 9 {
			soonestIndex = i
			break
		}
	}

	firstWordIndex := len(line)
	firstWord := 0
	for i, word := range numberWords {
		pos := strings.Index(line, word)
		if pos == -1 {
			continue
		}

		if pos < firstWordIndex && pos < soonestIndex {
			firstWordIndex = pos
			firstWord = i
		}
	}

	if firstWordIndex < len(line) {
		return fmt.Sprint(firstWord + 1)
	}

	return string(line[soonestIndex])
}

func getLastDigit(line string) string {
	lastIndex := 0
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		d := int(c - '0')
		if d <= 9 {
			lastIndex = i
			break
		}
	}

	lastWordIndex := -1
	lastWord := 0
	for i, word := range numberWords {
		pos := strings.LastIndex(line, word)
		if pos == -1 {
			continue
		}

		if pos > lastWordIndex && pos > lastIndex {
			lastWordIndex = pos
			lastWord = i
		}
	}

	if lastWordIndex > -1 {
		return fmt.Sprint(lastWord + 1)
	}

	return string(line[lastIndex])
}
