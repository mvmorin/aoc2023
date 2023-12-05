package day01

import (
	"fmt"
	_ "embed"
)

//go:embed input.txt
var input []byte

var digit_words [9]string = [9]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func try_parse_digit(buf []byte, i int, words bool) (int, int, bool) {
	val := 0
	found := false
	if '0' <= buf[i] && buf[i] <= '9' {
		found = true
		val = int(buf[i] - '0')
	}

	for wi := 0; words && !found && wi < len(digit_words); wi++ {
		word := digit_words[wi]
		sub := 0
		for i+sub < len(buf) && sub < len(word) && buf[i+sub] == word[sub] {
			sub++
		}
		found = sub == len(word)
		if found {
			val = wi + 1
			i += sub - 2
		}
	}

	i++
	return val, i, found
}

func Prob1() int {
	buf := input

	sum := 0
	first := -1
	last := -1
	i := 0
	for i < len(buf) {
		if buf[i] == '\n' || i+1 == len(buf) {
			sum += 10*first + last
			i++
			first = -1
			last = -1
			continue
		}

		val, i_new, found := try_parse_digit(buf, i, false)
		i = i_new
		if found {
			last = val
			if first < 0 {
				first = val
			}
		}
	}

	fmt.Println(sum)
	return sum
}

func Prob2() int {
	buf := input

	sum := 0
	first := -1
	last := -1
	i := 0
	for i < len(buf) {
		if buf[i] == '\n' || i+1 == len(buf) {
			cal := 10*first + last
			sum += cal
			i++
			first = -1
			last = -1
			continue
		}

		val, i_new, found := try_parse_digit(buf, i, true)
		i = i_new
		if found {
			last = val
			if first < 0 {
				first = val
			}
		}
	}

	fmt.Println(sum)
	return sum
}
