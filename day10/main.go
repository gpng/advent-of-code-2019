package day10

import (
	"log"
	"math"
	"sort"

	"github.com/gpng/advent-of-code-2019/utils"
)

// Run day 10
func Run() {
	log.Println("Running day 10")
	defer utils.Timer("Day 10 total")()

	lines := utils.ScanFileLinesToStrings("day10/input.txt", "")

	empty := [][]int{}
	asteroids := [][]int{}
	for i, line := range lines {
		for j, v := range line {
			coords := []int{j, i}
			switch v {
			case "#":
				asteroids = append(asteroids, coords)
			case ".":
				empty = append(empty, coords)
			}
		}
	}

	max, maxCoord, mapMaxCoords := newStation(asteroids)
	log.Printf("Part 1: %d", max)

	coord := kill(mapMaxCoords, maxCoord)
	log.Printf("Part 2: %d", coord[0]*100+coord[1])
}

func newStation(asteroids [][]int) (int, []int, map[float64][][]int) {
	defer utils.Timer("Part 1")()
	max := 0
	maxCoord := []int{0, 0}
	mapMaxCoords := map[float64][][]int{}
	for _, coords := range asteroids {
		count := 0
		gradientMap := map[float64][][]int{} // gradient: steps
		for _, asteroid := range asteroids {
			if asteroid[0] == coords[0] && asteroid[1] == coords[1] {
				continue
			}
			x := float64(asteroid[0] - coords[0])
			y := float64(asteroid[1] - coords[1])
			var gradient float64
			if x != 0 && y != 0 {
				gradient = math.Atan(math.Abs(y)/math.Abs(x)) * (180 / math.Pi)
				if y < 0 && x < 0 {
					gradient += 180
				}
				if y < 0 && x > 0 {
					gradient *= -1
				}
				if y > 0 && x < 0 {
					gradient += 90
				}
			} else if x < 0 {
				gradient = 180
			} else if x > 0 {
				gradient = 0
			} else if y < 0 {
				gradient = -90
			} else {
				gradient = 90
			}
			gradient += 90
			if gradient < 0 {
				gradient += 360
			}
			if gradientMap[gradient] == nil {
				count++
				gradientMap[gradient] = [][]int{asteroid}
			}
			gradientMap[gradient] = append(gradientMap[gradient], asteroid)
		}
		if count > max {
			max = count
			maxCoord = coords
			mapMaxCoords = gradientMap
		}
	}
	return max, maxCoord, mapMaxCoords
}

func kill(asteroidMap map[float64][][]int, coords []int) []int {
	defer utils.Timer("Part 2")()
	keys := []float64{}
	for k := range asteroidMap {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	step := 0
	for true {
		for _, k := range keys {
			asteroids := asteroidMap[k]
			if len(asteroids) > 0 {
				step++
				closestDistance := math.MaxFloat64
				var closestIndex int
				for i, v := range asteroids {
					x := math.Abs(float64(v[0] - coords[0]))
					y := math.Abs(float64(v[1] - coords[1]))
					distance := x + y
					if distance < closestDistance {
						closestDistance = distance
						closestIndex = i
					}
				}
				if step == 200 {
					return asteroids[closestIndex]
				}
				asteroidMap[k] = append(asteroidMap[k][0:closestIndex], asteroidMap[k][closestIndex+1:]...)
			}
		}
	}
	return []int{0, 0}
}
