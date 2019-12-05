package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func timer(message string) func() {
	start := time.Now()
	return func() { fmt.Println(message, ": ", time.Since(start)) }
}

func main() {
	log.Println("Running day 5")
	defer timer("day 5 total")()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		text += line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	intcodeStrings := strings.Split(text, ",")
	intcodes := []int{}

	for _, v := range intcodeStrings {
		intcode, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("Failed to convert text to integer\nline: %s\nerror: %v", v, err)
		}
		intcodes = append(intcodes, intcode)
	}

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
