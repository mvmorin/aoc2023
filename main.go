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
