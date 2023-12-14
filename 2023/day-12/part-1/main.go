package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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

type Puzzle struct {
	input   string
	numbers []int
}

func process(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	puzzles := make([]Puzzle, 0)
	for _, line := range lines {
		split := strings.Fields(line)
		numbers := make([]int, 0)
		for _, numStr := range strings.Split(split[1], ",") {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}

		puzzles = append(puzzles, Puzzle{
			input:   split[0],
			numbers: numbers,
		})
	}

	var wg sync.WaitGroup
	c := make(chan int, len(puzzles))
	count := 0
	allPuzzles := len(puzzles)
	for _, puzzle := range puzzles {
		wg.Add(1)
		func(puzzle Puzzle) {
			defer wg.Done()
			c <- calculate(puzzle)
			count++
			log.Printf("%d/%d", count, allPuzzles)
		}(puzzle)
	}

	go func() {
		wg.Wait()
		close(c)
		log.Println("All done!")
	}()

	total := 0
	for res := range c {
		total += res
	}

	return total
}

func calculate(puzzle Puzzle) int {
	placeablePositions := make([]int, 0)
	for i, val := range puzzle.input {
		if val == '?' {
			placeablePositions = append(placeablePositions, i)
		}
	}

	return recurse(puzzle.input, puzzle.numbers)
}

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

	if len(nums) == 1 && !strings.Contains(input, ".") && len(input) == nums[0] {
		return 1
	}

	total := 0

	if len(input) > 1 && (input[0] == '.' || input[0] == '?') {
		total += recurse(input[1:], nums)
	}

	if input[0] == '#' || input[0] == '?' {
		if nums[0]+1 <= len(input) && !strings.Contains(input[:nums[0]], ".") && (nums[0] == len(input) || input[nums[0]] != '#') {
			total += recurse(input[nums[0]+1:], nums[1:])
		}
	}

	return total
}
