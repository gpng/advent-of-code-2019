package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	log.Println("Running day 3")
	defer timer("day 3 total")()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, strings.Split(line, ","))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minDistance, leastSteps := minDistanceAndSteps(instructions[0], instructions[1])
	log.Printf("Part 1: Minimum distance: %f", minDistance)
	log.Printf("Part 2: Least steps: %d", leastSteps)
}

func minDistanceAndSteps(path1 []string, path2 []string) (float64, int) {
	gridMap := map[int]map[int][]int{} // [x][y][no of visits, steps]
	mapPath(path1, gridMap, true)
	mapPath(path2, gridMap, false)

	minDistance := math.MaxFloat64
	leastSteps := math.MaxInt32

	for x, xv := range gridMap {
		for y, yv := range xv {
			if yv[0] > 1 {
				distance := math.Abs(float64(x)) + math.Abs(float64(y))
				if distance < minDistance {
					minDistance = distance
				}
				if yv[1] < leastSteps {
					leastSteps = yv[1]
				}
			}
		}
	}
	return minDistance, leastSteps
}

func mapPath(path []string, gridMap map[int]map[int][]int, first bool) {
	x := 0
	y := 0
	totalSteps := 0
	for _, instr := range path {
		checkAndInitMapX(gridMap, x)

		dir, steps := splitInstr(instr)
		if dir == "R" {
			for i := 0; i < steps; i++ {
				totalSteps++
				x++
				checkAndInitMapX(gridMap, x)
				incrSteps(gridMap, x, y, totalSteps, first)
			}
		}
		if dir == "L" {
			for i := 0; i < steps; i++ {
				totalSteps++
				x--
				checkAndInitMapX(gridMap, x)
				incrSteps(gridMap, x, y, totalSteps, first)
			}
		}
		if dir == "U" {
			for i := 0; i < steps; i++ {
				totalSteps++
				y++
				checkAndInitMapX(gridMap, x)
				incrSteps(gridMap, x, y, totalSteps, first)
			}
		}
		if dir == "D" {
			for i := 0; i < steps; i++ {
				totalSteps++
				y--
				checkAndInitMapX(gridMap, x)
				incrSteps(gridMap, x, y, totalSteps, first)
			}
		}
	}
}

// check if top level map has been initialised, otherwise initialise to prevent nil reference err
func checkAndInitMapX(gridMap map[int]map[int][]int, x int) {
	if gridMap[x] == nil {
		gridMap[x] = map[int][]int{}
	}
}

func incrSteps(gridMap map[int]map[int][]int, x int, y int, steps int, first bool) {
	if len(gridMap[x][y]) == 0 {
		gridMap[x][y] = []int{0, steps}
	}
	if first {
		gridMap[x][y] = []int{1, steps}
	} else if gridMap[x][y][0] > 0 {
		gridMap[x][y][0] = 2
		gridMap[x][y][1] += steps
	}
}

func splitInstr(instr string) (string, int) {
	dir := string(instr[0])
	steps, err := strconv.Atoi(instr[1:])
	if err != nil {
		log.Printf("Failed to convert text to integer\ninstr: %s\nerror: %v", instr, err)
	}
	return dir, steps
}
