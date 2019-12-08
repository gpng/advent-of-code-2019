package day3

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gpng/advent-of-code-2019/utils"
)

// Run day 3
func Run() {
	log.Println("Running day 3")
	defer utils.Timer("Day 3 total")()

	file, err := os.Open("day3/input.txt")
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
	intersections := mapPath(path2, gridMap, false)

	minDistance := math.MaxFloat64
	leastSteps := math.MaxInt32

	for _, v := range intersections {
		distance := math.Abs(float64(v[0])) + math.Abs(float64(v[1]))
		if distance < minDistance {
			minDistance = distance
		}
		if v[2] < leastSteps {
			leastSteps = v[2]
		}
	}
	return minDistance, leastSteps
}

func mapPath(path []string, gridMap map[int]map[int][]int, first bool) [][]int {
	x := 0
	y := 0
	totalSteps := 0
	intersections := [][]int{}
	for _, instr := range path {
		checkAndInitMapX(gridMap, x)

		dir, steps := splitInstr(instr)
		if dir == "R" {
			for i := 0; i < steps; i++ {
				x++
				checkAndInitMapX(gridMap, x)
				incrSteps(gridMap, x, y, &totalSteps, first, &intersections)
			}
		}
		if dir == "L" {
			for i := 0; i < steps; i++ {
				x--
				checkAndInitMapX(gridMap, x)
				incrSteps(gridMap, x, y, &totalSteps, first, &intersections)
			}
		}
		if dir == "U" {
			for i := 0; i < steps; i++ {
				y++
				incrSteps(gridMap, x, y, &totalSteps, first, &intersections)
			}
		}
		if dir == "D" {
			for i := 0; i < steps; i++ {
				y--
				incrSteps(gridMap, x, y, &totalSteps, first, &intersections)
			}
		}
	}
	return intersections
}

// check if top level map has been initialised, otherwise initialise to prevent nil reference err
func checkAndInitMapX(gridMap map[int]map[int][]int, x int) {
	if gridMap[x] == nil {
		gridMap[x] = map[int][]int{}
	}
}

func incrSteps(gridMap map[int]map[int][]int, x int, y int, steps *int, first bool, intersections *[][]int) {
	*steps++
	if len(gridMap[x][y]) == 0 {
		gridMap[x][y] = []int{0, *steps}
	}
	if first {
		gridMap[x][y] = []int{1, *steps}
	} else if gridMap[x][y][0] > 0 {
		gridMap[x][y][0] = 2
		gridMap[x][y][1] += *steps
		*intersections = append(*intersections, []int{x, y, gridMap[x][y][1]})
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
