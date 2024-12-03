package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := Day3Part2(string(input))
	fmt.Println(result)
}

func Day3Part1(input string) int {
	total := 0
	re := regexp.MustCompile(`(?m)mul\(\d{1,3},\d{1,3}\)`)
	for _, match := range re.FindAllString(input, -1) {
		content := match[4 : len(match)-1]
		numbers := strings.Split(content, ",")
		for i := 0; i < len(numbers); i += 2 {
			num1, err := strconv.Atoi(numbers[i])
			if err != nil {
				continue
			}
			num2, err := strconv.Atoi(numbers[i+1])
			if err != nil {
				continue
			}
			total += num1 * num2
		}
	}

	return total
}

func Day3Part2(input string) int {
	total := 0
	re := regexp.MustCompile(`(?m)do\(\)|don\'t\(\)|mul\(\d{1,3},\d{1,3}\)`)
	mulEnabled := true
	for _, match := range re.FindAllString(input, -1) {
		if match == "do()" {
			mulEnabled = true
			continue
		} else if match == "don't()" {
			mulEnabled = false
			continue
		}

		if !mulEnabled {
			continue
		}

		content := match[4 : len(match)-1]
		numbers := strings.Split(content, ",")
		for i := 0; i < len(numbers); i += 2 {
			num1, err := strconv.Atoi(numbers[i])
			if err != nil {
				continue
			}
			num2, err := strconv.Atoi(numbers[i+1])
			if err != nil {
				continue
			}
			total += num1 * num2
		}
	}

	return total
}
