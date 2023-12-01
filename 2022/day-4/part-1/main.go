package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type AssignedRange struct {
	min int
	max int
}

func main() {
	rawInput, err := getInput()
	if err != nil {
		log.Fatal("Failed to get input data")
	}

	result := getPairsThatFullyContain(rawInput)
	log.Print(result)
}

func getPairsThatFullyContain(input string) int {
	pairsThatFullyContain := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		elf1 := parseSectionRange(pairs[0])
		elf2 := parseSectionRange(pairs[1])

		elf1ContainsElf2 := elf1.min <= elf2.min && elf1.max >= elf2.max
		elf2ContainsElf1 := elf2.min <= elf1.min && elf2.max >= elf1.max
		if elf1ContainsElf2 || elf2ContainsElf1 {
			pairsThatFullyContain++
		}
	}

	return pairsThatFullyContain
}

func parseSectionRange(sectionRangeString string) AssignedRange {
	sectionRange := strings.Split(sectionRangeString, "-")
	min, _ := strconv.Atoi(sectionRange[0])
	max, _ := strconv.Atoi(sectionRange[1])
	return AssignedRange{min: min, max: max}
}

func getInput() (string, error) {
	data, err := os.ReadFile("input.txt")
	return string(data), err
}
