package main

import (
	"log"
	"os"
	"testing"
)

func loadInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func TestGetCalibrationsValueSum(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

	expected := 142
	result := getInputValueSumV2(input)

	if result != expected {
		t.Errorf("Expected: %d found: %d", expected, result)
	}
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
