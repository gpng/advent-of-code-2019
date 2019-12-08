package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/gpng/advent-of-code-2019/utils"
)

// Run day 1
func Run() {
	log.Println("Running day 1")
	defer utils.Timer("Day 1 total")()

	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	masses := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("Failed to convert line to integer\nline: %s\nerror: %v", line, err)
		} else {
			masses = append(masses, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fuelRequirements := sumOfFuel(masses)
	log.Printf("Part 1: Sum of fuel requirements: %d", fuelRequirements)

	fuelRequirementsAdded := sumOfFuelAdded(masses)
	log.Printf("Part 2: Sum of fuel requirements with added reuqirements: %d", fuelRequirementsAdded)
}

func massFuel(mass int) int {
	res := (mass / 3) - 2
	if res > 0 {
		return res
	}
	return 0
}

func sumOfFuel(masses []int) int {
	defer utils.Timer("Part 1")()
	sum := 0

	for _, v := range masses {
		sum += massFuel(v)
	}

	return sum
}

func sumOfFuelAdded(masses []int) int {
	defer utils.Timer("Part 2")()
	sum := 0

	for _, v := range masses {
		lastFuel := massFuel(v)
		moduleSum := lastFuel
		for lastFuel > 0 {
			lastFuel = massFuel(lastFuel)
			moduleSum += lastFuel
		}
		sum += moduleSum
	}

	return sum
}
