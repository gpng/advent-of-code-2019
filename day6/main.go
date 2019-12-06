package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func timer(message string) func() {
	start := time.Now()
	return func() { fmt.Println(message, ": ", time.Since(start)) }
}

func main() {
	log.Println("Running day 6")
	defer timer("day 6 total")()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	orbits := map[string][]string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		orbit := strings.Split(line, ")")

		orbits[orbit[0]] = append(orbits[orbit[0]], orbit[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	counts := map[string]int{}
	count3, orders := totalOrbits(orbits, counts)
	log.Printf("Part 1: Total orbits: %d", count3)

	min := minOrbits(counts, orders, "YOU", "SAN")

	log.Printf("Part 2: Minimum orbital transfers: %d", min)
}

func totalOrbits(orbits map[string][]string, counts map[string]int) (int, map[string][]string) {
	defer timer("part 1")()
	count := 0
	orders := map[string][]string{}
	com := orbits["COM"]
	recurOrbits(orbits, counts, 1, com, []string{}, orders)

	for _, v := range counts {
		count += v
	}
	return count, orders
}

// calculate the min transfers needed by finding the intersection between the 2 paths
func minOrbits(counts map[string]int, orders map[string][]string, keyStart string, keyEnd string) int {
	defer timer("part 2")()
	pathStart := orders[keyStart]
	pathEnd := orders[keyEnd]

	// find last common orbit
	var last string
	for i := 0; i < len(pathStart); i++ {
		if pathStart[i] != pathEnd[i] {
			last = pathStart[i-1]
			break
		}
	}
	return counts["YOU"] + counts["SAN"] - (2 * counts[last]) - 2
}

// returns number of orbits from COM to every object
// as well as a map of all paths per end node
func recurOrbits(orbits map[string][]string, counts map[string]int, step int, list []string, order []string, orders map[string][]string) {
	for _, v := range list {
		counts[v] = step
		newOrder := append(order, v)
		newList := orbits[v]
		if len(newList) > 0 {
			recurOrbits(orbits, counts, step+1, newList, newOrder, orders)
		} else {
			orders[v] = newOrder
		}
	}
}
