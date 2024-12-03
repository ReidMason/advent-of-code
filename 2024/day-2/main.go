package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := Day2Part2(string(input))
	fmt.Println(result)
}

func Day2Part1(input string) int {
	total := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if isLineValid(line) {
			total++
		}
	}

	return total
}

func Day2Part2(input string) int {
	total := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		chars := strings.Split(line, " ")
		for i := -1; i < len(chars); i++ {
			check := strings.Join(remove(chars, max(i, 0)), " ")
			if isLineValid(check) {
				total++
				break
			}
		}
	}

	return total
}

func remove(slice []string, index int) []string {
	newSlice := make([]string, len(slice)-1)
	copy(newSlice, slice[:index])
	copy(newSlice[index:], slice[index+1:])
	return newSlice
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func isLineValid(line string) bool {
	increasing := true

	characters := strings.Split(line, " ")
	numbers := make([]int, len(characters))
	for i, character := range characters {
		number, err := strconv.Atoi(character)
		if err != nil {
			return false
		}

		if i == 1 {
			increasing = numbers[0] < number
		}

		if i >= 1 {
			diff := math.Abs(float64(numbers[i-1] - number))
			if diff == 0 || diff > 3 {
				return false
			}
		}

		if i > 1 && increasing && numbers[i-1] > number {
			return false
		} else if i > 1 && !increasing && numbers[i-1] < number {
			return false
		}

		numbers[i] = number
	}

	return true
}
