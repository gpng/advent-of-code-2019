package main

import "testing"

func TestSumOfFuel(t *testing.T) {
	testCases := []struct {
		masses []int
		result int
	}{
		{masses: []int{12}, result: 2},
		{masses: []int{14}, result: 2},
		{masses: []int{1969}, result: 654},
		{masses: []int{100756}, result: 33583},
		{masses: []int{12, 14, 1969, 100756}, result: 34241},
	}

	for _, tc := range testCases {
		result := sumOfFuel(tc.masses)
		if tc.result != result {
			t.Errorf("Expected %d but got %d", tc.result, result)
		}
	}
}

func TestSumOfFuelAdded(t *testing.T) {
	testCases := []struct {
		masses []int
		result int
	}{
		{masses: []int{14}, result: 2},
		{masses: []int{1969}, result: 966},
		{masses: []int{100756}, result: 50346},
		{masses: []int{14, 1969, 100756}, result: 51314},
	}

	for _, tc := range testCases {
		result := sumOfFuelAdded(tc.masses)
		if tc.result != result {
			t.Errorf("Expected %d but got %d", tc.result, result)
		}
	}
}
