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
	input := string(data)
	res := process(input)
	log.Println(res)
}

func process(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		total += parseLine(line)
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
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
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
