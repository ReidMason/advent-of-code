package main

import (
	"log"
	"os"
	"testing"
)

func TestGetCalibrationsValueSum(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

	expected := 281
	result := getInputValueSum(input)

	if result != expected {
		t.Errorf("Expected: %d found: %d", expected, result)
	}
}

func loadInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func BenchmarkGetCalibrationsValueSumV2(b *testing.B) {
	input := loadInput()

	b.ResetTimer()
	getInputValueSumV2(input)
}

func BenchmarkGetCalibrationsValueSum(b *testing.B) {
	input := loadInput()

	b.ResetTimer()
	getInputValueSum(input)
}
