package main

import "testing"

func TestGetCalibrationsValueSum(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	input = "1fivethree4two5threetmm"

	expected := 281
	expected = 13
	result := getInputValueSum(input)

	if result != expected {
		t.Errorf("Expected: %d found: %d", expected, result)
	}
}
