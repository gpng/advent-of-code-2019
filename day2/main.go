package day2

import (
	"log"

	"github.com/gpng/advent-of-code-2019/utils"
)

// Run day 2
func Run() {
	log.Println("Running day 2")
	defer utils.Timer("Day 2 total")()

	intcodes := utils.ScanFileLinesToInt("day2/input.txt", ",")

	final := program(intcodes, 12, 2)
	log.Printf("Part 1: Final value: %d", final)

	final2 := pair(intcodes)
	log.Printf("Part 2: Final value: %d", final2)
}

func alg(intcode []int, noun int, verb int) int {
	// copy slice to prevent modifying original instructions
	codes := make([]int, len(intcode))
	copy(codes, intcode)
	codes[1] = noun
	codes[2] = verb

	pos := 0

	for codes[pos] != 99 {
		instr := codes[pos]
		if instr == 1 {
			codes[codes[pos+3]] = codes[codes[pos+1]] + codes[codes[pos+2]]
		}
		if instr == 2 {
			codes[codes[pos+3]] = codes[codes[pos+1]] * codes[codes[pos+2]]
		}
		pos += 4
	}
	return codes[0]
}

// move function execution to separate function for timing
func program(intcodes []int, noun int, verb int) int {
	defer utils.Timer("Part 1")()
	return alg(intcodes, noun, verb)
}

func pair(intcodes []int) int {
	defer utils.Timer("Part 2")()
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			res := alg(intcodes, i, j)
			if res == 19690720 {
				return (100 * i) + j
			}
		}
	}
	return 0
}
