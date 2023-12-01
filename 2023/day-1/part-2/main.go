package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numberWords = [...]struct {
	digit string
	value byte
}{
	{
		digit: "one",
		value: 1,
	},

	{
		digit: "two",
		value: 2,
	},
	{
		digit: "three",
		value: 3,
	},
	{
		digit: "four",
		value: 4,
	},
	{
		digit: "five",
		value: 5,
	},
	{
		digit: "six",
		value: 6,
	},
	{
		digit: "seven",
		value: 7,
	},
	{
		digit: "eight",
		value: 8,
	},
	{
		digit: "nine",
		value: 9,
	},
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)
	res := getInputValueSum(input)
	log.Println(res)
}

func getInputValueSumV2(input string) int {
	total := 0
	lines := strings.Split(strings.TrimRight(input, "\n\r"), "\n")

	for _, line := range lines {
		firstDigit := getFirstDigitV2(line)
		lastDigit := getLastDigitV2(line)
		digit := int(firstDigit*10 + lastDigit)

		total += digit
	}

	return total
}

func getFirstDigitV2(line string) byte {
	firstWordIndex := len(line)
	var digitValue byte = 0
	for _, word := range numberWords {
		pos := strings.Index(line, word.digit)
		if pos == -1 {
			continue
		}

		if pos < firstWordIndex {
			firstWordIndex = pos
			digitValue = word.value
		}
	}

	for i := 0; i < firstWordIndex; i++ {
		c := line[i]
		if c <= 57 {
			return c - '0'
		}
	}

	return digitValue
}

func getLastDigitV2(line string) byte {
	lastWordIndex := 0
	var digitValue byte = 0
	for _, word := range numberWords {
		pos := strings.LastIndex(line, word.digit)
		if pos == -1 {
			continue
		}

		if pos > lastWordIndex {
			lastWordIndex = pos
			digitValue = word.value
		}
	}

	for i := len(line) - 1; i >= lastWordIndex; i-- {
		c := line[i]
		if c <= 57 {
			return c - '0'
		}
	}

	return digitValue
}

var numberWords1 = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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
	for i, word := range numberWords1 {
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
	for i, word := range numberWords1 {
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
