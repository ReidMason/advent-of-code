package main

import (
	"log"
	"os"
	"strconv"
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
	lines := strings.Split(strings.TrimSpace(input), "\n")

	total := 0
	for _, line := range lines {
		nums := toInt(strings.Fields(line))
		total += getDiffs(nums) + nums[len(nums)-1]
	}

	return total
}

func toInt(input []string) []int {
	output := make([]int, 0, len(input))
	for _, val := range input {
		num, _ := strconv.Atoi(val)
		output = append(output, num)
	}

	return output
}

func getDiffs(nums []int) int {
	diffs := make([]int, 0)

	increment := nums[len(nums)-1] - nums[len(nums)-2]
	allSame := true
	for i := 1; i < len(nums); i++ {
		newVal := nums[i] - nums[i-1]
		if allSame && newVal != increment {
			allSame = false
		}
		diffs = append(diffs, newVal)
	}

	if allSame {
		return increment
	}

	childIncrement := getDiffs(diffs)
	return diffs[len(diffs)-1] + childIncrement
}
