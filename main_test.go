package main

import (
	"reflect"
	"testing"
)


func TestFirstUniqueChars(t *testing.T) {
	first := firstUniqChar("stsijn")
	if first != 1 {
		t.Error("For", "stsijn", "expected", 1, "got", first)
	}
}

func TestTwoSum(t *testing.T) {
	type TestPairIntegerSlices struct {
		expected []int
		input []int
		target int
	}

	testPairs := []TestPairIntegerSlices{
		TestPairIntegerSlices{expected: []int{0, 1}, input: []int{2,7,11,15}, target: 9},
		TestPairIntegerSlices{expected: []int{1,2}, input: []int{3,2,4}, target: 6},
		TestPairIntegerSlices{expected: []int{0,3}, input: []int{0,4,3,0}, target: 0},
	}

	for _, test := range testPairs {
		output := twoSum(test.input, test.target)
		if !reflect.DeepEqual(output, test.expected) {
			t.Error(
				"For", test.input,
				"with target", test.target,
				"got", output,
				"expected", test.expected,
				)
		}
	}
}
