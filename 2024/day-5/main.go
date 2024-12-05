package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := Day5Part1(string(input))
	fmt.Println(result)
}

func Day5Part1(input string) int {
	lines := strings.Split(input, "\n")

	// Get rules
	orderRules := make([][]string, 0)
	pos := 0
	for i, line := range lines {
		pos = i
		if line == "" {
			break
		}
		orderRules = append(orderRules, strings.Split(line, "|"))
	}

	// Get print orders
	score := 0
	for pos < len(lines) {
		line := lines[pos]
		if line == "" {
			pos++
			continue
		}

		numbers := strings.Split(line, ",")
		if isOrderValid(orderRules, numbers) {
			score += getMidpoint(numbers)
		}
		pos++
	}

	return score
}

func getMidpoint(numbers []string) int {
	midIndex := len(numbers) / 2
	number, err := strconv.Atoi(numbers[midIndex])
	if err != nil {
		panic(err)
	}
	return number
}

func isOrderValid(orderRules [][]string, numbers []string) bool {
	for _, orderRule := range orderRules {
		firstIndex, secondIndex := getIndex(orderRule[0], orderRule[1], numbers)
		if firstIndex != -1 && secondIndex != -1 && firstIndex > secondIndex {
			return false
		}
	}

	return true
}

func getIndex(first, second string, numbers []string) (int, int) {
	firstIndex := -1
	secondIndex := -1
	for i, number := range numbers {
		if number == first {
			firstIndex = i
		}
		if number == second {
			secondIndex = i
		}

		if firstIndex != -1 && secondIndex != -1 {
			return firstIndex, secondIndex
		}
	}

	return firstIndex, secondIndex
}
