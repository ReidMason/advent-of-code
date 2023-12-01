package main

import "testing"

func TestGetCalibrationsValueSum(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	expected := 142
	result := getInputValueSum(input)

	if result != expected {
		t.Errorf("Expected: %d found: %d", expected, result)
	}
}
