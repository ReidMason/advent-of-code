package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.TrimSpace(string(data))
	res := process(input)
	log.Println(res)
}

func process(input string) int {
	lines := strings.Split(strings.TrimRight(input, "\n\r"), "\n")

	tallies := make([]int, len(lines))
	for i := range lines {
		tallies[i] = 1
	}

	total := 0
	for i, line := range lines {
		score := parseLine(line)
		multiplier := tallies[i]
		for x := 1; x <= score; x++ {
			nextIdx := i + x
			tallies[nextIdx] += multiplier
		}

		total += 1 * multiplier
	}

	return total
}

func parseLine(line string) int {
	line += " "
	pos := strings.Index(line, ":") + 2

	score := 0
	winningNums := make([]string, 0)
	gettingWinners := true
	currNum := ""
	for pos < len(line) {
		c := line[pos]
		if c >= '0' && c <= '9' {
			currNum += string(c)
		} else if c == ' ' && len(currNum) > 0 {
			if gettingWinners {
				winningNums = append(winningNums, currNum)
			} else if contains(winningNums, currNum) {
				score += 1
			}
			currNum = ""
		} else if c == '|' {
			gettingWinners = false
		}

		pos++
	}

	return score
}

func contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}

	return false
}
