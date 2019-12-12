package day9

import (
	"log"

	"github.com/gpng/advent-of-code-2019/utils"
)

// Run day 9
func Run() {
	log.Println("Running day 9")
	defer utils.Timer("Day 9 total")()

	intcodes := utils.ScanFileLinesToInt("day9/input.txt", ",")

	keycode := boostKeycode(intcodes)
	log.Printf("Part 1: BOOST keycode: %d", keycode)

	signal := distressSignal(intcodes)
	log.Printf("Part 2: Distress signal: %d", signal)
}

func boostKeycode(intcodes []int) int {
	defer utils.Timer("Part 1")()
	c := &com{0, 0, 0, intcodes, false}
	for !c.end {
		c.setInput(1, false)
	}
	return c.output
}

func distressSignal(intcodes []int) int {
	defer utils.Timer("Part 2")()
	c := &com{0, 0, 0, intcodes, false}
	for !c.end {
		c.setInput(2, false)
	}
	return c.output
}

type com struct {
	output       int
	pos          int
	relativeBase int
	intcodes     []int
	end          bool
}

func (c *com) setInput(input int, waitOutput bool) {
	intcodes, output, pos, relativeBase, end := intcodeCom(c.intcodes, c.output, c.pos, c.relativeBase, input, waitOutput)
	c.intcodes = intcodes
	c.output = output
	c.pos = pos
	c.relativeBase = relativeBase
	c.end = end
}

func intcodeCom(intcodes []int, prevOutput int, startPos int, relativeBase int, input int, waitOutput bool) ([]int, int, int, int, bool) {
	// copy slice to prevent modifying original instructions
	codes := make([]int, len(intcodes))
	copy(codes, intcodes)
	pos := startPos
	output := prevOutput
	base := relativeBase
	for codes[pos] != 99 {
		instr := splitInt(codes[pos])
		opCode := instr[0]
		switch opCode {
		case 1: // add
			param1 := paramValue(codes, instr, pos, base, 1, false)
			param2 := paramValue(codes, instr, pos, base, 2, false)
			param3 := paramValue(codes, instr, pos, base, 3, true)
			updateSliceLength(&codes, param3)
			codes[param3] = param1 + param2
			if param3 != pos {
				pos += 4
			}
		case 2: // multiply
			param1 := paramValue(codes, instr, pos, base, 1, false)
			param2 := paramValue(codes, instr, pos, base, 2, false)
			param3 := paramValue(codes, instr, pos, base, 3, true)
			updateSliceLength(&codes, param3)
			codes[param3] = param1 * param2
			if param3 != pos {
				pos += 4
			}
		case 3: // set input
			param1 := paramValue(codes, instr, pos, base, 1, true)
			updateSliceLength(&codes, param1)
			codes[param1] = input
			pos += 2
			if !waitOutput {
				return codes, output, pos, base, false
			}
		case 4: // output
			param1 := paramValue(codes, instr, pos, base, 1, false)
			output = param1
			pos += 2
			if waitOutput {
				return codes, output, pos, base, false
			}
		case 5: // jump if true
			param1 := paramValue(codes, instr, pos, base, 1, false)
			param2 := paramValue(codes, instr, pos, base, 2, false)
			if param1 > 0 {
				pos = param2
			} else {
				pos += 3
			}
		case 6: // jump if false
			param1 := paramValue(codes, instr, pos, base, 1, false)
			param2 := paramValue(codes, instr, pos, base, 2, false)
			if param1 == 0 {
				pos = param2
			} else {
				pos += 3
			}
		case 7: // less than
			param1 := paramValue(codes, instr, pos, base, 1, false)
			param2 := paramValue(codes, instr, pos, base, 2, false)
			param3 := paramValue(codes, instr, pos, base, 3, true)
			updateSliceLength(&codes, param3)
			if param1 < param2 {
				codes[param3] = 1
			} else {
				codes[param3] = 0
			}
			if param3 != pos {
				pos += 4
			}
		case 8: // equals
			param1 := paramValue(codes, instr, pos, base, 1, false)
			param2 := paramValue(codes, instr, pos, base, 2, false)
			param3 := paramValue(codes, instr, pos, base, 3, true)
			updateSliceLength(&codes, param3)
			if param1 == param2 {
				codes[param3] = 1
			} else {
				codes[param3] = 0
			}
			if param3 != pos {
				pos += 4
			}
		case 9: // adjust relative base
			param1 := paramValue(codes, instr, pos, base, 1, false)
			base = base + param1
			pos += 2
		}
		updateSliceLength(&codes, pos)
	}
	return codes, output, pos, base, true
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

func paramValue(codes []int, instr []int, pos int, base int, index int, returnIndex bool) int {
	mode := 0
	if len(instr) > index {
		mode = instr[index]
	}
	newIndex := codes[pos+index]
	switch mode {
	case 1:
		newIndex = pos + index
	case 2:
		newIndex = base + codes[pos+index]
	}
	if returnIndex {
		return newIndex
	}
	if newIndex >= len(codes) {
		return 0
	}
	return codes[newIndex]
}

func updateSliceLength(codes *[]int, pos int) {
	if pos >= len(*codes) {
		add := make([]int, pos+1-len(*codes))
		*codes = append(*codes, add...)
	}
}
