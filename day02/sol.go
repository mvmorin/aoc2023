package day02

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Prob1() int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	id_sum := 0

	for _, line := range lines {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ':' || r == ',' || r == ';' || r == ' '
		})

		id, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		possible := true

		for i := 2; i+1 < len(parts); i = i + 2 {
			color := parts[i+1]
			val, err := strconv.Atoi(parts[i])
			if err != nil {
				panic(err)
			}

			limit := -1
			switch color {
			case "red":
				limit = 12
			case "green":
				limit = 13
			case "blue":
				limit = 14
			}

			if val > limit {
				possible = false
				break
			}
		}

		if possible {
			id_sum = id_sum + id
		}
	}
	fmt.Println(id_sum)
	return id_sum
}

func Prob2() int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	power_sum := 0

	for _, line := range lines {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ':' || r == ',' || r == ';' || r == ' '
		})

		max_red := 0
		max_green := 0
		max_blue := 0

		for i := 2; i+1 < len(parts); i = i + 2 {
			color := parts[i+1]
			val, err := strconv.Atoi(parts[i])
			if err != nil {
				panic(err)
			}

			switch color {
			case "red":
				max_red = max(max_red, val)
			case "green":
				max_green = max(max_green, val)
			case "blue":
				max_blue = max(max_blue, val)
			}
		}

		power := max_red*max_green*max_blue

		power_sum = power_sum + power
	}
	fmt.Println(power_sum)
	return power_sum
}
