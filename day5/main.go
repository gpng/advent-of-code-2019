package main

import (
	"fmt"
	"log"
	"time"
)

func timer(message string) func() {
	start := time.Now()
	return func() { fmt.Println(message, ": ", time.Since(start)) }
}

func main() {
	log.Println("Running day 5")
	defer timer("day 5 total")()

	intcodes := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1002, 36, 25, 224, 1001, 224, -2100, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1102, 31, 84, 225, 1102, 29, 77, 225, 1, 176, 188, 224, 101, -42, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 223, 224, 223, 2, 196, 183, 224, 1001, 224, -990, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 224, 223, 223, 102, 14, 40, 224, 101, -1078, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1001, 180, 64, 224, 101, -128, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 223, 224, 223, 1102, 24, 17, 224, 1001, 224, -408, 224, 4, 224, 1002, 223, 8, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1101, 9, 66, 224, 1001, 224, -75, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 223, 224, 223, 1102, 18, 33, 225, 1101, 57, 64, 225, 1102, 45, 11, 225, 1101, 45, 9, 225, 1101, 11, 34, 225, 1102, 59, 22, 225, 101, 89, 191, 224, 1001, 224, -100, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 329, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 1001, 223, 1, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 359, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 374, 101, 1, 223, 223, 1008, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 389, 101, 1, 223, 223, 8, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 404, 101, 1, 223, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 419, 1001, 223, 1, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 434, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 449, 1001, 223, 1, 223, 107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 464, 1001, 223, 1, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 479, 1001, 223, 1, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 494, 1001, 223, 1, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 524, 101, 1, 223, 223, 1007, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 539, 1001, 223, 1, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 569, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 584, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 599, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 629, 101, 1, 223, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 644, 1001, 223, 1, 223, 108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 659, 1001, 223, 1, 223, 7, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}

	diagnosticCode := program1(intcodes, 1)
	log.Printf("Part 1: Final value: %d", diagnosticCode)
	diagnosticCode2 := program2(intcodes, 5)
	log.Printf("Part 2: Final value: %d", diagnosticCode2)
}

func alg(intcodes []int, input int) int {
	// copy slice to prevent modifying original instructions
	codes := make([]int, len(intcodes))
	copy(codes, intcodes)
	pos := 0
	lastOutput := 0
	for codes[pos] != 99 {
		instr := splitInt(codes[pos])
		opCode := instr[0]
		if opCode == 1 { // add
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			param3 := codes[pos+3]
			codes[param3] = param1 + param2
			if param3 != pos {
				pos += 4
			}
		}
		if opCode == 2 { // multiply
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			param3 := codes[pos+3]
			codes[param3] = param1 * param2
			if param3 != pos {
				pos += 4
			}
		}
		if opCode == 3 { // set input
			codes[codes[pos+1]] = input
			pos += 2
		}
		if opCode == 4 { // output
			param1 := paramValue(codes, instr, pos, 1)
			lastOutput = param1
			pos += 2
		}
		if opCode == 5 { // jump if true
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			if param1 > 0 {
				pos = param2
			} else {
				pos += 3
			}
		}
		if opCode == 6 { // jump if false
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			if param1 == 0 {
				pos = param2
			} else {
				pos += 3
			}
		}
		if opCode == 7 { // less than
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			param3 := codes[pos+3]
			if param1 < param2 {
				codes[param3] = 1
			} else {
				codes[param3] = 0
			}
			if param3 != pos {
				pos += 4
			}
		}
		if opCode == 8 { // equals
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			param3 := codes[pos+3]
			if param1 == param2 {
				codes[param3] = 1
			} else {
				codes[param3] = 0
			}
			if param3 != pos {
				pos += 4
			}
		}
	}
	return lastOutput
}

func splitInt(input int) []int {
	if input < 10 {
		return []int{input}
	}
	arr := []int{}
	for i := 1; i < input; i *= 10 {
		if i == 10 {
			continue
		}
		arr = append(arr, (input%(i*10))/i)
	}
	return arr
}

func paramValue(codes []int, instr []int, pos int, index int) int {
	mode := 0
	if len(instr) > index {
		mode = instr[index]
	}
	if mode == 1 {
		return codes[pos+index]
	}
	return codes[codes[pos+index]]
}

func program1(intcodes []int, input int) int {
	defer timer("part 1")()
	return alg(intcodes, input)
}

func program2(intcodes []int, input int) int {
	defer timer("part 2")()
	return alg(intcodes, input)
}
