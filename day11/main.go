package day11

import (
	"fmt"
	"github.com/gpng/advent-of-code-2019/utils"
	"log"
)

// Run day 11
func Run() {
	log.Println("Running day 11")
	defer utils.Timer("Day 11 total")()

	intcodes := utils.ScanFileLinesToInt("day11/input.txt", ",")

	painted := paint(intcodes, 0, 1)
	log.Printf("Part 1: Tiles painted: %d", painted)
	paint(intcodes, 1, 2)
}

func paint(intcodes []int, start int, part int) int {
	defer utils.Timer(fmt.Sprintf("Part %d", part))()
	coords := []int{0, 0}
	grid := map[int]map[int]int{0: map[int]int{0: start}}
	coloredGrid := map[int]map[int]bool{0: map[int]bool{}}

	c := &com{0, 0, 0, intcodes, false}
	dir := "U"
	count := 0
	for !c.end {
		x := coords[0]
		y := coords[1]
		input := 0
		if grid[x] != nil {
			input = grid[x][y]
		}
		c.setInput(input, true)
		color := c.output
		c.setInput(input, true)
		rotate := c.output
		if c.end {
			break
		}
		// paint
		if grid[x] == nil {
			grid[x] = map[int]int{}
			coloredGrid[x] = map[int]bool{}
		}
		grid[x][y] = color
		if !coloredGrid[x][y] {
			count++
			coloredGrid[x][y] = true
		}
		switch dir {
		case "U":
			switch rotate {
			case 0:
				coords = []int{x - 1, y}
				dir = "L"
			case 1:
				coords = []int{x + 1, y}
				dir = "R"
			}
		case "D":
			switch rotate {
			case 0:
				coords = []int{x + 1, y}
				dir = "R"
			case 1:
				coords = []int{x - 1, y}
				dir = "L"
			}
		case "L":
			switch rotate {
			case 0:
				coords = []int{x, y - 1}
				dir = "D"
			case 1:
				coords = []int{x, y + 1}
				dir = "U"
			}
		case "R":
			switch rotate {
			case 0:
				coords = []int{x, y + 1}
				dir = "U"
			case 1:
				coords = []int{x, y - 1}
				dir = "D"
			}
		}
	}
	if part == 1 {
		return count
	}
	rangeX := []int{0, 0}
	rangeY := []int{0, 0}
	for x, v := range grid {
		for y := range v {
			if y < rangeY[0] {
				rangeY[0] = y
			}
			if y > rangeY[1] {
				rangeY[1] = y
			}
		}
		if x < rangeX[0] {
			rangeX[0] = x
		}
		if x > rangeX[1] {
			rangeX[1] = x
		}
	}
	for y := rangeY[1]; y >= rangeY[0]; y-- {
		for x := rangeX[0]; x <= rangeX[1]; x++ {
			color := 0
			if grid[x] != nil {
				color = grid[x][y]
			}
			switch color {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("█")
			}
		}
		fmt.Print("\n")
	}
	return count
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
