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
	return func() { fmt.Println(message, time.Since(start)) }
}

func main() {
	log.Println("Running day 2")
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

	opCodesStrings := strings.Split(text, ",")
	opCodes := []int{}

	for _, v := range opCodesStrings {
		opCode, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("Failed to convert text to integer\nline: %s\nerror: %v", v, err)
		}
		opCodes = append(opCodes, opCode)
	}

	final := program(opCodes, 12, 2)
	log.Printf("Part 1: Final value: %d", final)

	final2 := pair(opCodes)
	log.Printf("Part 2: Final value: %d", final2)
}

func alg(opCodes []int, noun int, verb int) int {
	// copy slice to prevent modifying original instructions
	codes := make([]int, len(opCodes))
	copy(codes, opCodes)
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
func program(opCodes []int, noun int, verb int) int {
	defer timer("program")()
	return alg(opCodes, noun, verb)
}

func pair(opCodes []int) int {
	defer timer("pair")()
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			res := alg(opCodes, i, j)
			if res == 19690720 {
				return (100 * i) + j
			}
		}
	}
	return 0
}
