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
	"github.com/mvmorin/aoc2023/day10"
	"github.com/mvmorin/aoc2023/day11"
	"github.com/mvmorin/aoc2023/day12"
	"github.com/mvmorin/aoc2023/day13"
	"github.com/mvmorin/aoc2023/day14"
	"github.com/mvmorin/aoc2023/day15"
	"github.com/mvmorin/aoc2023/day16"
	"github.com/mvmorin/aoc2023/day17"
	"github.com/mvmorin/aoc2023/day18"
	"github.com/mvmorin/aoc2023/day19"
	"github.com/mvmorin/aoc2023/day20"
	"github.com/mvmorin/aoc2023/day21"
	"github.com/mvmorin/aoc2023/day22"
	"github.com/mvmorin/aoc2023/day23"
	"github.com/mvmorin/aoc2023/day24"
	"github.com/mvmorin/aoc2023/day25"
)

func main() {
	fmt.Print("Advent of Code 2023\n\n")

	fns := []struct{f func()int; run bool} {
		{day01.Prob1, true},
		{day01.Prob2, true},
		{day02.Prob1, true},
		{day02.Prob2, true},
		{day03.Prob1, true},
		{day03.Prob2, true},
		{day04.Prob1, true},
		{day04.Prob2, true},
		{day05.Prob1, true},
		{day05.Prob2, true},
		{day06.Prob1, true},
		{day06.Prob2, true},
		{day07.Prob1, true},
		{day07.Prob2, true},
		{day08.Prob1, true},
		{day08.Prob2, true},
		{day09.Prob1, true},
		{day09.Prob2, true},
		{day10.Prob1, true},
		{day10.Prob2, true},
		{day11.Prob1, true},
		{day11.Prob2, true},
		{day12.Prob1, false},
		{day12.Prob2, false},
		{day13.Prob1, true},
		{day13.Prob2, true},
		{day14.Prob1, true},
		{day14.Prob2, true},
		{day15.Prob1, true},
		{day15.Prob2, true},
		{day16.Prob1, true},
		{day16.Prob2, true},
		{day17.Prob1, true},
		{day17.Prob2, false},
		{day18.Prob1, false},
		{day18.Prob2, false},
		{day19.Prob1, false},
		{day19.Prob2, false},
		{day20.Prob1, false},
		{day20.Prob2, false},
		{day21.Prob1, false},
		{day21.Prob2, false},
		{day22.Prob1, true},
		{day22.Prob2, true},
		{day23.Prob1, true},
		{day23.Prob2, true},
		{day24.Prob1, true},
		{day24.Prob2, true},
		{day25.Prob1, true},
		{day25.Prob2, true},
	}

	start_time := time.Now()

	for i, t := range fns {
		if !t.run {
			continue
		}
		f := t.f
		fmt.Printf("Day %d, problem %d: ", i / 2 + 1, i % 2 + 1)
		tic := time.Now()
		f()
		fmt.Printf("Time: %s\n\n", time.Since(tic).String())
	}
	fmt.Printf("Total time: %s\n", time.Since(start_time).String())
}
