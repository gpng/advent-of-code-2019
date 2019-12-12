package day12

import (
	"log"
	"math"

	"github.com/gpng/advent-of-code-2019/utils"
)

// Run day 11
func Run() {
	log.Println("Running day 12")
	defer utils.Timer("Day 12 total")()

	steps := 1000

	positions := [][]int{
		[]int{17, 5, 1},
		[]int{-2, -8, 8},
		[]int{7, -6, 14},
		[]int{1, -10, 4},
	}

	velocities := [][]int{}
	for range positions {
		velocities = append(velocities, []int{0, 0, 0})
	}

	total := part1(positions, steps)
	log.Printf("Part 1: %f", total)

	count := part2(positions)
	log.Printf("Part 2: %d", count)
}

func part2(input [][]int) int {
	defer utils.Timer("Part 2")()
	positions := [][]int{}

	inputStateX := state(input, 0)
	inputStateY := state(input, 1)
	inputStateZ := state(input, 2)

	velocities := [][]int{}
	for _, v := range input {
		positions = append(positions, []int{v[0], v[1], v[2]})
		velocities = append(velocities, []int{0, 0, 0})
	}

	steps := 1
	var stepsX, stepsY, stepsZ int
	for true {
		steps++
		step(&positions, &velocities)
		if stepsX == 0 {
			if isRepeat(state(positions, 0), inputStateX) {
				stepsX = steps
			}
		}
		if stepsY == 0 {
			if isRepeat(state(positions, 1), inputStateY) {
				stepsY = steps
			}
		}
		if stepsZ == 0 {
			if isRepeat(state(positions, 2), inputStateZ) {
				stepsZ = steps
			}
		}
		if stepsX > 0 && stepsY > 0 && stepsZ > 0 {
			break
		}
	}
	return lcm(stepsX, stepsY, stepsZ)
}

func state(positions [][]int, pos int) []int {
	res := []int{}
	for _, v := range positions {
		res = append(res, v[pos])
	}
	return res
}

func isRepeat(state []int, inputState []int) bool {
	for i, vv := range state {
		if inputState[i] != vv {
			return false
		}
	}
	return true
}

func part1(input [][]int, steps int) float64 {
	defer utils.Timer("Part 1")()
	positions := [][]int{}

	velocities := [][]int{}
	for _, v := range input {
		positions = append(positions, []int{v[0], v[1], v[2]})
		velocities = append(velocities, []int{0, 0, 0})
	}

	for i := 0; i < steps; i++ {
		step(&positions, &velocities)
	}

	return totalEnergy(positions, velocities)
}

func totalEnergy(positions [][]int, velocities [][]int) float64 {
	var total float64
	for i := 0; i < len(positions); i++ {
		var ke, pe float64
		for j := 0; j < 3; j++ {
			ke += math.Abs(float64(positions[i][j]))
			pe += math.Abs(float64(velocities[i][j]))
		}
		total += (ke * pe)
	}
	return total
}

func step(positions *[][]int, velocities *[][]int) {
	for i := 0; i < len(*positions)-1; i++ {
		for j := i + 1; j < len(*positions); j++ {
			for k := 0; k < len((*positions)[i]); k++ {
				if (*positions)[i][k] < (*positions)[j][k] {
					(*velocities)[i][k]++
					(*velocities)[j][k]--
				} else if (*positions)[i][k] > (*positions)[j][k] {
					(*velocities)[i][k]--
					(*velocities)[j][k]++
				}
			}
		}
	}
	for i, v := range *velocities {
		for j, vv := range v {
			(*positions)[i][j] += vv
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
