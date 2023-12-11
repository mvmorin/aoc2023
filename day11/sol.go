package day11

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Idx struct {
	i int
	j int
}

func parse_galaxies(input string, expansion int) []Idx {
	cols := 0
	for ; cols < len(input) && input[cols] != '\n'; cols++ {
	}
	rows := len(input) / (cols + 1)

	row_is_not_empty := make([]bool, rows)
	col_is_not_empty := make([]bool, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			is_not_empty := input[i*(cols+1)+j] == '#'
			row_is_not_empty[i] = is_not_empty || row_is_not_empty[i]
			col_is_not_empty[j] = is_not_empty || col_is_not_empty[j]
		}
	}

	galaxies := make([]Idx, 0)
	empty_rows_passed := 0
	for i := 0; i < rows; i++ {
		if !row_is_not_empty[i] {
			empty_rows_passed++
		}

		empty_cols_passed := 0
		for j := 0; j < cols; j++ {
			if !col_is_not_empty[j] {
				empty_cols_passed++
			}

			if input[i*(rows+1)+j] == '#' {
				d_i := empty_rows_passed * (expansion - 1)
				d_j := empty_cols_passed * (expansion - 1)
				galaxies = append(galaxies, Idx{i + d_i, j + d_j})
			}
		}
	}

	return galaxies
}

func Prob1() int {
	galaxies := parse_galaxies(input, 2)
	sum := 0
	for i, a := range galaxies[:len(galaxies)-1] {
		for _, b := range galaxies[i+1:] {
			sum += abs(a.i-b.i) + abs(a.j-b.j)
		}
	}
	fmt.Println(sum)
	return sum
}

func Prob2() int {
	galaxies := parse_galaxies(input, 1000000)
	sum := 0
	for i, a := range galaxies[:len(galaxies)-1] {
		for _, b := range galaxies[i+1:] {
			sum += abs(a.i-b.i) + abs(a.j-b.j)
		}
	}
	fmt.Println(sum)
	return sum
}
