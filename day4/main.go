package day4

import (
	"fmt"
	"log"
	"time"
)

func timer(message string) func() {
	start := time.Now()
	return func() { fmt.Println(message, ": ", time.Since(start)) }
}

// Run day 4
func Run() {
	log.Println("Running day 4")
	defer timer("day 4 total")()
	start := 387638
	end := 919123

	// end := 387648
	count, count2 := decreasingAdjacentNumbers(start, end)

	log.Printf("Part 1: %d", count)
	log.Printf("Part 2: %d", count2)
}

func decreasingAdjacentNumbers(start int, end int) (int, int) {
	count := 0
	count2 := 0
	for i := start; i <= end; i++ {
		arr := splitInt(i)
		hasAdjacent := false
		hasDecrease := false
		sets := map[int]int{}
		setIndex := 0
		for j, v := range arr {
			if j == len(arr)-1 {
				break
			}
			next := arr[j+1]
			if next == v {
				if !hasAdjacent {
					hasAdjacent = true
				}
				if j > 0 && arr[j-1] != v {
					setIndex++
				}
				sets[setIndex]++
			}
			if next > v {
				hasDecrease = true
				break
			}
		}
		if hasAdjacent && !hasDecrease {
			count++
			for _, v := range sets {
				if v == 1 {
					count2++
					break
				}
			}
		}
	}
	return count, count2
}

func splitInt(input int) []int {
	arr := []int{}
	for i := 1; i < input; i *= 10 {
		arr = append(arr, (input%(i*10))/i)
	}
	return arr
}
