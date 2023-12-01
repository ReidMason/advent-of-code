package main

import (
	"log"
	"os"
	"strings"
)

type Group []string

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatal("Failed to get input")
	}

	total := 0
	groups := parseInputs(input)
	for _, group := range groups {
		commonItem := findCommonItemInGroup(group)
		itemPiority := getItemPriority(commonItem)
		total += itemPiority
	}

	log.Print(total)
}

func parseInputs(input string) []Group {
	var groups []Group
	var currentGroup Group
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		currentGroup = append(currentGroup, line)

		if len(currentGroup) == 3 {
			groups = append(groups, currentGroup)
			currentGroup = nil
		}
	}

	return groups
}

func findCommonItemInGroup(group Group) rune {
	for _, r1 := range group[0] {
		for _, r2 := range group[1] {
			for _, r3 := range group[2] {
				if r1 == r2 && r2 == r3 {
					return r3
				}
			}
		}
	}

	return 0
}

func getItemPriority(item rune) int {
	asciiOffset := 96
	b := int(item) - asciiOffset

	capitalAsciiOffset := 58
	if b < 0 {
		b += capitalAsciiOffset
	}

	return b
}

func getInput() (string, error) {
	data, err := os.ReadFile("input.txt")
	return string(data), err
}
