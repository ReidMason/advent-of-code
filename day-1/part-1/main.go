package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput, err := getInput()
	if err != nil {
		log.Fatal("Failed to get input data")
	}

	input := parseInput(rawInput)

	findHighestNumberOfCalories(input)
}

func parseInput(input string) []string {
	result := strings.Split(input, "\n")
	for i, v := range result {
		result[i] = strings.TrimSpace(v)
	}

	return result
}

func findHighestNumberOfCalories(input []string) {
	var highestCarryAmount int
	total := 0
	for i, v := range input {
		if v != "" {
			weight, _ := strconv.Atoi(v)
			total += weight
		}

		if v == "" || i == len(input)-1 {
			if total > highestCarryAmount {
				highestCarryAmount = total
			}
			total = 0
		}
	}

	log.Printf("The Elf carrying the most calories has %d", highestCarryAmount)
}

func getInput() (string, error) {
	data, err := os.ReadFile("input.txt")
	return string(data), err
}
