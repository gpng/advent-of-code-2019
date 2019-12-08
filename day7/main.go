package day7

import (
	"log"

	"github.com/gpng/advent-of-code-2019/utils"
)

// Run day 7
func Run() {
	log.Println("Running day 7")
	defer utils.Timer("Day 7 total")()

	intcodes := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 43, 60, 81, 94, 175, 256, 337, 418, 99999, 3, 9, 101, 2, 9, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001, 9, 4, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001, 9, 2, 9, 1002, 9, 3, 9, 101, 4, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 4, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}

	res1 := linearMode(intcodes, []int{0, 1, 2, 3, 4})
	log.Printf("Part 1: %d", res1)

	res2 := feedbackLoopMode(intcodes, []int{5, 6, 7, 8, 9})
	log.Printf("Part 2: %d", res2)
}

func linearMode(intcodes []int, phases []int) int {
	defer utils.Timer("Part 1")()
	maxOutput := 0
	for _, v := range permutations(phases) {
		output := 0
		for _, i := range v {
			c := &com{0, 0, intcodes, false}
			c.setInput(i, false) // set phase
			for !c.end {
				c.setInput(output, false)
			}
			output = c.output
		}
		if output > maxOutput {
			maxOutput = output
		}
	}
	return maxOutput
}

func feedbackLoopMode(intcodes []int, phases []int) int {
	defer utils.Timer("Part 2")()
	maxOutput := 0
	for _, v := range permutations(phases) {
		// initialise all machines and set phase
		machines := map[int]*com{}
		for i, v := range v {
			machines[i] = &com{0, 0, intcodes, false}
			machines[i].setInput(v, false)
		}
		output := 0
		end := false
		for !end {
			for i := 0; i < len(machines); i++ {
				machines[i].setInput(output, true)
				output = machines[i].output
				if i == len(machines)-1 && machines[i].end {
					end = true
				}
			}
		}

		if output > maxOutput {
			maxOutput = output
		}
	}
	return maxOutput
}

func permutations(list []int) [][]int {
	if len(list) == 2 {
		return [][]int{[]int{list[0], list[1]}, []int{list[1], list[0]}}
	}
	res := [][]int{}
	for i := range list {
		base := list[i]
		rest := []int{}
		rest = append(rest, list[:i]...)
		rest = append(rest, list[i+1:]...)
		for _, v := range permutations(rest) {
			res = append(res, append([]int{base}, v...))
		}
	}
	return res
}

type com struct {
	output   int
	pos      int
	intcodes []int
	end      bool
}

func (c *com) setInput(input int, waitOutput bool) {
	intcodes, output, pos, end := intcodeCom(c.intcodes, c.output, c.pos, input, waitOutput)
	c.intcodes = intcodes
	c.output = output
	c.pos = pos
	c.end = end
}

func intcodeCom(intcodes []int, prevOutput int, startPos int, input int, waitOutput bool) ([]int, int, int, bool) {
	// copy slice to prevent modifying original instructions
	codes := make([]int, len(intcodes))
	copy(codes, intcodes)
	pos := startPos
	output := prevOutput
	end := false
	for codes[pos] != 99 {
		instr := splitInt(codes[pos])
		opCode := instr[0]
		switch opCode {
		case 1: // add
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			param3 := codes[pos+3]
			codes[param3] = param1 + param2
			if param3 != pos {
				pos += 4
			}
		case 2: // multiply
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			param3 := codes[pos+3]
			codes[param3] = param1 * param2
			if param3 != pos {
				pos += 4
			}
		case 3: // set input
			codes[codes[pos+1]] = input
			pos += 2
			if waitOutput {
				end = true
			} else {
				return codes, output, pos, false
			}
		case 4: // output
			param1 := paramValue(codes, instr, pos, 1)
			output = param1
			pos += 2
			if waitOutput && end {
				return codes, output, pos, false
			}
		case 5: // jump if true
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			if param1 > 0 {
				pos = param2
			} else {
				pos += 3
			}
		case 6: // jump if false
			param1 := paramValue(codes, instr, pos, 1)
			param2 := paramValue(codes, instr, pos, 2)
			if param1 == 0 {
				pos = param2
			} else {
				pos += 3
			}
		case 7: // less than
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
		case 8: // equals
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
	return codes, output, pos, true
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
