package main

import (
	"fmt"
	"time"
)
import (
	"github.com/mvmorin/aoc2023/day01"
	"github.com/mvmorin/aoc2023/day02"
)

func main() {
	fmt.Print("Advent of Code 2023\n\n")

	fns := []func()int {
		day01.Prob1,
		day01.Prob2,
		day02.Prob1,
		day02.Prob2,
	}

	start_time := time.Now()

	for i, f := range fns {
		fmt.Printf("Day %d, problem %d: ", i / 2, i % 2 + 1)
		tic := time.Now()
		f()
		fmt.Printf("Time: %s\n\n", time.Since(tic).String())
	}
	fmt.Printf("Total time: %s\n", time.Since(start_time).String())
}
