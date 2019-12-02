package main

import "testing"

func TestProgram(t *testing.T) {
	testCases := []struct {
		opCodes []int
		result  int
	}{
		{opCodes: []int{1, 0, 0, 0, 99}, result: 2},
		{opCodes: []int{2, 3, 0, 3, 99}, result: 2},
		{opCodes: []int{2, 4, 4, 5, 99, 0}, result: 2},
		{opCodes: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, result: 30},
	}

	for _, tc := range testCases {
		result := program(tc.opCodes, tc.opCodes[1], tc.opCodes[2])
		if tc.result != result {
			t.Errorf("Expected %d but got %d", tc.result, result)
		}
	}
}
