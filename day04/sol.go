package day04

import (
	"fmt"
	_ "embed"
)

//go:embed input.txt
var input_bytes []byte

func bytes_to_int(buf []byte, i int) (int, int) {
	val := 0
	for ; '0' <= buf[i] && buf[i] <= '9'; i++ {
		val = 10*val + int(buf[i] - '0')
	}
	return val, i
}

func parse_game(buf []byte, i int) (int, int) {
	var val int
	winning := make([]int, 0, 10)
	matches := 0

	for ; buf[i] != ':'; i++ {}
	for i++; buf[i] == ' '; i++ {}

	for buf[i] != '|' {
		val, i = bytes_to_int(buf, i)
		winning = append(winning, val)

		for ; buf[i] == ' '; i++ {}
	}
	for i++; buf[i] == ' '; i++ {}

	for buf[i] != '\n' {
		val, i = bytes_to_int(buf, i)

		for _, w := range winning {
			if w == val {
				matches++
				break
			}
		}

		for ; buf[i] == ' '; i++ {}
	}

	i++
	return matches, i
}

func parse_matches_per_game(buf []byte) []int {
	matches := 0
	matches_per_game := make([]int, 0, 100)
	for i:= 0; i < len(buf)-1; {
		matches, i = parse_game(buf, i)
		matches_per_game = append(matches_per_game, matches)
	}
	return matches_per_game
}


func Prob1() int {
	matches_per_game := parse_matches_per_game(input_bytes)
	total_points := 0
	for _, matches := range matches_per_game {
		points := (1 << matches) >> 1
		total_points += points
	}
	fmt.Println(total_points)
	return total_points
}

func Prob2() int {
	matches_per_game := parse_matches_per_game(input_bytes)

	copies_per_game := make([]int, len(matches_per_game))
	for i := range copies_per_game {
		copies_per_game[i] = 1
	}

	for game, copies := range copies_per_game {
		matches := matches_per_game[game]

		for i := game+1; i <= game+matches; i++ {
			copies_per_game[i] += copies
		}
	}

	sum := 0
	for _, copies := range copies_per_game {
		sum += copies
	}

	fmt.Println(sum)
	return sum
}
