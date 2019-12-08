package day5

import "testing"

func TestProgram(t *testing.T) {
	testCases := []struct {
		opCodes []int
		input   int
		result  int
	}{
		{opCodes: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: 8, result: 1},
		{opCodes: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, input: 7, result: 1},
		{opCodes: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: 8, result: 1},
		{opCodes: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, input: 7, result: 1},
		{opCodes: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, input: 7, result: 999},
		{opCodes: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, input: 8, result: 1000},
		{opCodes: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, input: 9, result: 1001},
	}

	for _, tc := range testCases {
		result := alg(tc.opCodes, tc.input)
		if tc.result != result {
			t.Errorf("Expected %d but got %d", tc.result, result)
		}
	}
}
