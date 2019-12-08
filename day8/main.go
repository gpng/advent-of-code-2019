package day8

import (
	"fmt"
	"log"
	"math"

	"github.com/gpng/advent-of-code-2019/utils"
)

// Run day 8
func Run() {
	log.Println("Running day 8")
	defer utils.Timer("Day 8 total")()

	data := utils.ScanFileLinesToInt("day8/input.txt", "")

	width := 25
	height := 6

	layers, res := minZeroes(data, width, height)
	log.Printf("Part 1: %d", res)
	printImage(layers, width, height)
}

func minZeroes(data []int, width int, height int) ([][]int, int) {
	defer utils.Timer("Part 1")()
	count0 := 0
	count1 := 0
	count2 := 0

	layers := [][]int{}
	layerStats := [][]int{}
	layerIndex := 0
	layer := []int{}

	steps := 0
	currentMin := math.MaxInt32
	currentMinIndex := 0
	for _, v := range data {
		layer = append(layer, v)
		switch v {
		case 0:
			count0++
		case 1:
			count1++
		case 2:
			count2++
		}
		steps++
		if steps == width*height {
			c := []int{}
			copy(c, layer)
			layers = append(layers, layer)
			layerStats = append(layerStats, []int{count0, count1, count2})
			if count0 < currentMin {
				currentMinIndex = layerIndex
				currentMin = count0
			}
			layer = []int{}
			count0 = 0
			count1 = 0
			count2 = 0
			layerIndex++
			steps = 0
		}
	}

	return layers, layerStats[currentMinIndex][1] * layerStats[currentMinIndex][2]
}

func printImage(layers [][]int, width int, height int) {
	defer utils.Timer("Part 2")()
	index := 0
	for i := 0; i < width*height; i++ {
		for j := 0; j < len(layers); j++ {
			v := layers[j][i]
			if v == 0 {
				fmt.Print(" ")
				break
			}
			if v == 0 || v == 1 {
				fmt.Print("█")
				break
			}
			if j == len(layers)-1 { // last layer
				fmt.Print(2)
			}
		}
		index++
		if index == width {
			fmt.Print("\n")
			index = 0
		}
	}
}
