package main

import (
	"fmt"
	"time"
)
import (
	"github.com/mvmorin/aoc2023/day01"
	"github.com/mvmorin/aoc2023/day02"
	"github.com/mvmorin/aoc2023/day03"
	"github.com/mvmorin/aoc2023/day04"
	"github.com/mvmorin/aoc2023/day05"
	"github.com/mvmorin/aoc2023/day06"
	"github.com/mvmorin/aoc2023/day07"
	"github.com/mvmorin/aoc2023/day08"
	"github.com/mvmorin/aoc2023/day09"
)

func main() {
	fmt.Print("Advent of Code 2023\n\n")

	fns := []func()int {
		day01.Prob1,
		day01.Prob2,
		day02.Prob1,
		day02.Prob2,
		day03.Prob1,
		day03.Prob2,
		day04.Prob1,
		day04.Prob2,
		day05.Prob1,
		day05.Prob2,
		day06.Prob1,
		day06.Prob2,
		day07.Prob1,
		day07.Prob2,
		day08.Prob1,
		day08.Prob2,
		day09.Prob1,
		day09.Prob2,
	}

	start_time := time.Now()

	for i, f := range fns {
		fmt.Printf("Day %d, problem %d: ", i / 2 + 1, i % 2 + 1)
		tic := time.Now()
		f()
		fmt.Printf("Time: %s\n\n", time.Since(tic).String())
	}
	fmt.Printf("Total time: %s\n", time.Since(start_time).String())
}
