package actions

import (
	"testing"
)

type CalcTest struct {
	input    string
	expected string
}

var calcTests = []CalcTest{
	{"1+1", "2.00"},
	{"5/4*3", "3.75"},
	{"5*4/3", "6.67"},
	{"3-4", "-1.00"},
	{"2/0", "Division by Zero isn't allowed!"},
}

func TestCalc(t *testing.T) {
	for _, test := range calcTests {
		result := calc(test.input)
		if result != test.expected {
			t.Fatal("Error: Input:", test.input, "Output:", result, "Expected:", test.expected)
		}
	}
}
