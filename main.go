package main

import (
	"flag"
	"github.com/gpng/advent-of-code-2019/day11"
	"github.com/gpng/advent-of-code-2019/day12"
	"github.com/gpng/advent-of-code-2019/day13"
	"log"
	"os"

	"github.com/gpng/advent-of-code-2019/day1"
	"github.com/gpng/advent-of-code-2019/day10"
	"github.com/gpng/advent-of-code-2019/day2"
	"github.com/gpng/advent-of-code-2019/day3"
	"github.com/gpng/advent-of-code-2019/day4"
	"github.com/gpng/advent-of-code-2019/day5"
	"github.com/gpng/advent-of-code-2019/day6"
	"github.com/gpng/advent-of-code-2019/day7"
	"github.com/gpng/advent-of-code-2019/day8"
	"github.com/gpng/advent-of-code-2019/day9"
)

func main() {
	day := flag.Int("d", 0, "Day to run")

	flag.Parse()

	runMap := map[int]func(){
		1:  day1.Run,
		2:  day2.Run,
		3:  day3.Run,
		4:  day4.Run,
		5:  day5.Run,
		6:  day6.Run,
		7:  day7.Run,
		8:  day8.Run,
		9:  day9.Run,
		10: day10.Run,
		11: day11.Run,
		12: day12.Run,
		13: day13.Run,
	}

	// run all
	if *day == 0 {
		for i := 1; i <= len(runMap); i++ {
			runMap[i]()
		}
		os.Exit(0)
	}

	fn, ok := runMap[*day]
	if !ok {
		log.Fatalf("Invalid day %d", *day)
	}

	fn()

	os.Exit(0)
}
