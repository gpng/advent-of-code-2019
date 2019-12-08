package day3

import (
	"strings"
	"testing"
)

func TestProgram(t *testing.T) {
	testCases := []struct {
		path1       []string
		path2       []string
		minDistance float64
		leastSteps  int
	}{
		{
			path1:       strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ","),
			path2:       strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ","),
			minDistance: 159,
			leastSteps:  610,
		},
		{
			path1:       strings.Split("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", ","),
			path2:       strings.Split("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", ","),
			minDistance: 135,
			leastSteps:  410,
		},
	}

	for _, tc := range testCases {
		minDistance, leastSteps := minDistanceAndSteps(tc.path1, tc.path2)
		if tc.minDistance != minDistance || tc.leastSteps != leastSteps {
			t.Errorf("Expected %f, %d but got %f, %d", tc.minDistance, tc.leastSteps, minDistance, leastSteps)
		}
	}
}
