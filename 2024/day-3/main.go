package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := Day3Part1(string(input))
	fmt.Println(result)
}

func Day3Part1(input string) int {
	total := 0
	re := regexp.MustCompile(`(?m)mul\(\d{1,3},\d{1,3}\)`)
	for _, match := range re.FindAllString(input, -1) {
		re2 := regexp.MustCompile(`\d{1,3}`)
		numbers := re2.FindAllString(match, -1)
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
