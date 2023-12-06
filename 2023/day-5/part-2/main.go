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

func process(input string) int {
	pos := 0
	lines := strings.Split(input, "\n")
	seedLine := lines[pos]

	pos += 3

	mappings := make([][]Mapping, 0)
	log.Println("Building mapping")
	for pos < len(lines) {
		mapping, newPos := buildMapping(lines, pos)
		pos = newPos
		mappings = append(mappings, mapping)
	}
	log.Println("Done building mapping")

	log.Println("Processing seeds")
	return getSeeds(seedLine, mappings)
}

func getSeeds(line string, mappings [][]Mapping) int {
	rawStrs := strings.Split(line, " ")[1:]

	c := make(chan int)
	var wg sync.WaitGroup

	log.Println("Starting seeds process")
	total := len(rawStrs)
	for i := 0; i < total; i += 2 {
		seedNum, _ := strconv.Atoi(rawStrs[i])
		seedRange, _ := strconv.Atoi(rawStrs[i+1])

		wg.Add(1)
		go func(seedStart, seedRange int) {
			defer wg.Done()
			seeds := make([]int, 0)
			for j := 0; j < seedRange; j++ {
				seed := seedStart + j
				seed = applyMapping(seed, mappings)
				seeds = append(seeds, seed)
			}
			c <- minSlice(seeds)
		}(seedNum, seedRange)
	}

	go func() {
		wg.Wait()
		close(c)
		log.Println("Seeds process done, collating...")
	}()

	minVal := -1
	for s := range c {
		if minVal == -1 || s < minVal {
			minVal = s
		}
	}

	log.Println("process done, collating...")
	return minVal
}

func applyMapping(seed int, mappings [][]Mapping) int {
	for _, mapping := range mappings {
		for _, m := range mapping {
			if seed >= m.start && seed <= m.end {
				seed += m.inc
				break
			}
		}
	}

	return seed
}

func get(mapping map[int]int, source int) int {
	val, ok := mapping[source]
	if !ok {
		return source
	}

	return val
}

func minSlice(vals []int) int {
	lowest := vals[0]
	for _, val := range vals {
		if val < lowest {
			lowest = val
		}
	}

	return lowest
}

type Mapping struct {
	start, end, inc int
}

func buildMapping(lines []string, pos int) ([]Mapping, int) {
	currLine := lines[pos]
	mapping := make([]Mapping, 0)

	for currLine != "" {
		nums := strings.Split(currLine, " ")
		steps, _ := strconv.Atoi(nums[2])
		sourceStart, _ := strconv.Atoi(nums[1])
		destStart, _ := strconv.Atoi(nums[0])
		sourceEnd := sourceStart + steps - 1

		mapping = append(mapping, Mapping{sourceStart, sourceEnd, destStart - sourceStart})

		pos++
		currLine = lines[pos]
	}

	pos += 2

	return mapping, pos
}
