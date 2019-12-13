package day13

import (
	"fmt"
	"log"

	"github.com/gpng/advent-of-code-2019/utils"
)

// Run day 13
func Run() {
	log.Println("Running day 13")
	defer utils.Timer("Day 13 total")()

	intcodes := utils.ScanFileLinesToInt("day13/input.txt", ",")

	blocks := part1(intcodes)
	log.Printf("Part 1: Number of blocks: %d", blocks)

	score := part2(intcodes)
	log.Printf("Part 2: Score: %d", score)
}

func part1(codes []int) int {
	defer utils.Timer("Part 1")()

	intcodes := append([]int{}, codes...)

	c := &com{0, 0, 0, intcodes, false}

	_, _, _, _, countBlocks := updateGrid(c, 0)
	return countBlocks
}

func part2(codes []int) int {
	defer utils.Timer("Part 2")()
	intcodes := append([]int{}, codes...)

	// Memory address 0 represents the number of quarters that have been inserted; set it to 2 to play for free.
	intcodes[0] = 2
	c := &com{0, 0, 0, intcodes, false}

	var score, countBlocks, posBall, posPaddle int
	// print initial grid
	grid, posBall, posPaddle, score, countBlocks := updateGrid(c, 0)
	// find grid range
	rangeX := []int{0, 0}
	rangeY := []int{0, 0}
	for y, v := range grid {
		for x := range v {
			if x < rangeX[0] {
				rangeX[0] = x
			}
			if x > rangeX[1] {
				rangeX[1] = x
			}
		}
		if y < rangeY[0] {
			rangeY[0] = y
		}
		if y > rangeY[1] {
			rangeY[1] = y
		}
	}

	for countBlocks > 0 {
		input := 0
		if posBall > posPaddle { // move right
			input = 1
		} else if posBall < posPaddle { // move left
			input = -1
		}
		grid, posBall, posPaddle, score, countBlocks = updateGrid(c, input)
		// print each frame result only for debugging
		// printGrid(rangeX[0], rangeX[1], rangeY[0], rangeY[1], grid)
		// log.Println("score", score)
		// log.Println("blocks left", countBlocks)
		// log.Println("pos ball", posBall)
		// log.Println("pos paddle", posPaddle)
	}
	return score
}

func printGrid(minX, maxX, minY, maxY int, grid map[int]map[int]int) {
	fmt.Println("----------------------")
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[y] != nil {
				fmt.Print(grid[y][x])
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println("")
	}
}

func updateGrid(c *com, input int) (map[int]map[int]int, int, int, int, int) {
	grid := map[int]map[int]int{}
	step := -1
	x := 0
	y := 0
	score := 0
	c.end = false
	c.pos = 0
	c.relativeBase = 0
	posBall := 0
	pospaddle := 0
	countBlocks := 0

	for !c.end {
		step++
		c.setInput(input, true)
		switch step {
		case 0:
			x = c.output
		case 1:
			y = c.output
		case 2:
			if x == -1 && y == 0 {
				if c.output > 0 { // for some reason the last step outputs 0 score
					score = c.output
				}
			} else {
				if grid[y] == nil {
					grid[y] = map[int]int{}
				}
				grid[y][x] = c.output
				switch c.output {
				case 2:
					countBlocks++
				case 3:
					posBall = x
				case 4:
					pospaddle = x
				}
			}
			step = -1
		}
	}
	return grid, posBall, pospaddle, score, countBlocks
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
