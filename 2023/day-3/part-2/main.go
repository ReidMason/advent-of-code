package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id, red, gree, blue int
}

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
	lineLength := strings.Index(input, "\n")
	idxs := []int{-lineLength - 1, -lineLength, -lineLength + 1, -1, 1, lineLength - 1, lineLength, lineLength + 1}

	input = strings.ReplaceAll(input, "\n", "")

	total := 0
	for i, c := range input {
		if c != '*' {
			continue
		}

		endIdx := -1
		foundNums := make([]int, 0)
		for _, j := range idxs {
			if i+j <= endIdx {
				continue
			}

			if hasNum(input, i+j) {
				num, x := getNum(input, i+j)
				endIdx = x
				foundNums = append(foundNums, num)
			}
		}
		if len(foundNums) == 2 {
			total += foundNums[0] * foundNums[1]
		}
	}

	return total
}

func hasNum(text string, idx int) bool {
	if idx < 0 || idx >= len(text) {
		return false
	}

	c := text[idx]
	return c >= '0' && c <= '9'
}

func getNum(text string, idx int) (int, int) {
	numStart := 0
	for i := idx; i > 0; i-- {
		c := text[i]
		if c > '9' || c < '0' {
			numStart = i + 1
			break
		}
	}

	currNum := ""
	endIdx := 0
	for i := numStart; i < len(text); i++ {
		c := text[i]
		if c > '9' || c < '0' {
			endIdx = i - 1
			break
		}

		currNum += string(c)
	}

	numVal, _ := strconv.Atoi(currNum)
	return numVal, endIdx
}
