package day01

import (
	"fmt"
	// "os"
	"strings"
	"unicode"
	"embed"
)

//go:embed input.txt input_test.txt
var efs embed.FS

func read_lines(f string) []string {
	bytes, err := efs.ReadFile(f)
    if err != nil {
        // fmt.Print(err)
		panic(err)
    }

	s := string(bytes)
	lines := strings.Split(s, "\n")
	return lines[:len(lines)-1]
}

func Prob1() int {
	lines := read_lines("input.txt")

	sum := 0
	for _, line := range lines {
		digit := -1
		for _, r := range line {
			if unicode.IsDigit(r) {
				cur_digit := int(r - '0')
				if digit == -1 {
					sum = sum + 10*cur_digit
				}
				digit = cur_digit
			}
		}
		sum = sum + digit
	}

	fmt.Println(sum)
	return sum

}

func Prob2() int {
	lines := read_lines("input.txt")

	digit_words := [...]string{
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

	sum := 0
	for _, line := range lines {
		first_digit := -1
		second_digit := -1

		for i, r := range line {
			found_digit := -1

			if unicode.IsDigit(r) {
				found_digit = int(r - '0')
			} else {
				for d, w := range digit_words {
					if i+len(w) <= len(line) && w == line[i:i+len(w)] {
						found_digit = d + 1
						break
					}
				}
			}

			if found_digit >= 0 {
				if first_digit < 0 {
					first_digit = found_digit
				}
				second_digit = found_digit
			}
		}
		sum = sum + 10*first_digit + second_digit
	}

	fmt.Println(sum)
	return sum
}
