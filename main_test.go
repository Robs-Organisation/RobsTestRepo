package main

import (
	"testing"
)

type CalcTest struct {
	inputOne int
	inputTwo int
	expected int
}

var calcTests = []CalcTest{
	{1, 2, 3},
	{-2, -3, -5},
}

func TestCalc(t *testing.T) {
	for _, test := range calcTests {
		result := add(test.inputOne, test.inputTwo)
		if result != test.expected {
			t.Fatal("Error: Input:", test.inputOne, test.inputTwo, "Output:", result, "Expected:", test.expected)
		}
	}
}
