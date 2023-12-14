package main

import (
	"fmt"
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
		split := strings.Fields(line)
		numbers := make([]int, 0)
		for _, numStr := range strings.Split(split[1], ",") {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}

		inputStr := split[0]
		numbersOriginal := numbers
		for i := 0; i < 4; i++ {
			inputStr += "?" + split[0]
			numbers = append(numbers, numbersOriginal...)
		}

		total += calculate(inputStr, numbers)
	}

	return total
}

func calculate(input string, nums []int) int {
	placeablePositions := make([]int, 0)
	for i, val := range input {
		if val == '?' {
			placeablePositions = append(placeablePositions, i)
		}
	}

	return recurse(input, nums)
}

var cache = make(map[string]int)

func recurse(input string, nums []int) int {
	if input == "" {
		if len(nums) == 0 {
			return 1
		}

		// Input is empty but we need more nums INVALID
		return 0
	}

	// No more nums needed
	if len(nums) == 0 {
		// There are leftover hashes INVALID
		if strings.Contains(input, "#") {
			return 0
		}
		return 1
	}

	memo := input
	for _, num := range nums {
		memo += fmt.Sprintf("%d,", num)
	}
	res, found := cache[memo]
	if found {
		return res
	}

	if len(nums) == 1 && !strings.Contains(input, ".") && len(input) == nums[0] {
		return 1
	}

	total := 0

	if input[0] == '.' || input[0] == '?' {
		if len(input) > 1 {
			total += recurse(input[1:], nums)
		} else {
			total += recurse("", nums)
		}
	}

	if input[0] == '#' || input[0] == '?' {
		if nums[0] <= len(input) && !strings.Contains(input[:nums[0]], ".") && (nums[0] == len(input) || input[nums[0]] != '#') {

			if nums[0]+1 >= len(input) {
				total += recurse("", nums[1:])
			} else {
				total += recurse(input[nums[0]+1:], nums[1:])
			}
		}
	}

	cache[memo] = total
	return total
}
