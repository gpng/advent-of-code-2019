package main

import (
	"flag"
	"github.com/gpng/advent-of-code-2019/day7"
	"github.com/gpng/advent-of-code-2019/day8"
	"log"
	"os"

	"github.com/gpng/advent-of-code-2019/day1"
	"github.com/gpng/advent-of-code-2019/day2"
	"github.com/gpng/advent-of-code-2019/day3"
	"github.com/gpng/advent-of-code-2019/day4"
	"github.com/gpng/advent-of-code-2019/day5"
	"github.com/gpng/advent-of-code-2019/day6"
)

func main() {
	day := flag.Int("d", 0, "Day to run")

	flag.Parse()

	if *day == 0 {
		log.Fatalf("-d Day flag required")
	}

	switch *day {
	case 0:
		log.Fatalf("-d Day flag required")
	case 1:
		day1.Run()
	case 2:
		day2.Run()
	case 3:
		day3.Run()
	case 4:
		day4.Run()
	case 5:
		day5.Run()
	case 6:
		day6.Run()
	case 7:
		day7.Run()
	case 8:
		day8.Run()
	default:
		log.Fatalf("Invalid day %d", *day)
	}

	os.Exit(0)
}
