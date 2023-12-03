package day03

import (
	_ "embed"
	"fmt"
	"slices"
	"unicode"
)

//go:embed input.txt
var input []byte

type Schematic struct {
	rows int
	cols int
	schm []byte
}

type Number struct {
	val   int
	row   int
	start int
	end   int
}

func parse(input []byte) Schematic {
	cols := slices.Index(input, '\n') + 1
	rows := len(input) / cols
	schm := Schematic{rows: rows, cols: cols, schm: input}

	return schm
}

func has_symbol_neighbour(nbr Number, schm Schematic) bool {
	for r := max(0, nbr.row-1); r <= min(schm.rows-1, nbr.row+1); r++ {
		for c := max(0, nbr.start-1); c <= min(schm.cols-1, nbr.end+1); c++ {
			b := schm.schm[c+r*schm.cols]

			if b != '\n' && !unicode.IsDigit(rune(b)) && b != '.' {
				return true
			}
		}
	}

	return false
}

func Prob1() int {
	schm := parse(input)

	var nbr *Number = nil

	sum := 0

	for r := 0; r < schm.rows; r++ {
		for c := 0; c < schm.cols; c++ {
			b := schm.schm[c+r*schm.cols]
			is_digit := unicode.IsDigit(rune(b))

			if is_digit && nbr == nil {
				nbr = &Number{
					val:   int(b - '0'),
					row:   r,
					start: c,
				}
			} else if is_digit && nbr != nil {
				nbr.val = 10*nbr.val + int(b-'0')
			} else if !is_digit && nbr != nil {
				nbr.end = c - 1
				if has_symbol_neighbour(*nbr, schm) {
					sum = sum + nbr.val
				}
				nbr = nil
			}
		}
	}

	fmt.Println(sum)
	return sum
}

func Prob2() int {
	schm := parse(input)

	var nbr *Number = nil
	gears := make([]struct {
		r int
		c int
	}, 0)
	nbrs := make([][]*Number, schm.cols)
	for r := range nbrs {
		nbrs[r] = make([]*Number, 0)
	}

	for r := 0; r < schm.rows; r++ {
		for c := 0; c < schm.cols; c++ {
			b := schm.schm[c+r*schm.cols]
			is_digit := unicode.IsDigit(rune(b))

			if b == '*' {
				gears = append(gears, struct {
					r int
					c int
				}{r, c})
			}

			if is_digit && nbr == nil {
				nbr = &Number{
					val:   int(b - '0'),
					row:   r,
					start: c,
				}
			} else if is_digit && nbr != nil {
				nbr.val = 10*nbr.val + int(b-'0')
			} else if !is_digit && nbr != nil {
				nbr.end = c - 1
				nbrs[r] = append(nbrs[r], nbr)
				nbr = nil
			}
		}
	}

	sum := 0
	for _, g := range gears {
		vals := make([]int, 0)
		for r := max(g.r-1, 0); r <= min(g.r+1, len(nbrs)); r++ {
			for _, nbr := range nbrs[r] {
				up := nbr.row - 1
				low := nbr.row + 1
				left := nbr.start - 1
				right := nbr.end + 1

				if up <= g.r && g.r <= low && left <= g.c && g.c <= right {
					vals = append(vals, nbr.val)
				}
			}
		}

		if len(vals) == 2 {
			sum += vals[0] * vals[1]
		}
	}

	fmt.Println(sum)
	return sum
}
