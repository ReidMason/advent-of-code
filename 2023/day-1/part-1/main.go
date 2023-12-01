package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)
	res := getInputValueSumV2(input)
	log.Println(res)
}

func getInputValueSumV2(input string) int {
	total := 0
	lines := strings.Split(strings.TrimRight(input, "\n\r"), "\n")

	for _, line := range lines {
		firstDigit := getFirstDigitV2(line) - '0'
		lastDigit := getLastDigitV2(line) - '0'
		digit := int(firstDigit*10 + lastDigit)

		total += digit
	}

	return total
}

func getFirstDigitV2(line string) byte {
	for _, c := range []byte(line) {
		if c <= 57 {
			return c
		}
	}

	return 0
}

func getLastDigitV2(line string) byte {
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if c <= 57 {
			return c
		}
	}

	return 0
}

func getInputValueSum(input string) int {
	total := 0
	lines := strings.Split(strings.TrimRight(input, "\n\r"), "\n")

	for _, line := range lines {
		line = strings.TrimRight(line, "\n\r")
		firstDigit := string(getFirstDigit(line))
		lastDigit := string(getLastDigit(line))
		digit, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic(err)
		}

		total += digit
	}

	return total
}

func getFirstDigit(line string) byte {
	for i, c := range line {
		d := int(c - '0')
		if d <= 9 {
			return line[i]
		}
	}

	return 0
}

func getLastDigit(line string) byte {
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		d := int(c - '0')
		if d <= 9 {
			return c
		}
	}

	return 0
}
