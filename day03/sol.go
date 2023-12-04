package day03

import (
	_ "embed"
	"fmt"
	"slices"
)

//go:embed input.txt
var input []byte

func is_digit(b byte) bool {
	return '0' <= b && b <= '9'
}

func parse_val(buf []byte, i int) (int, int, int) {
	for ; is_digit(buf[i]) && i > 0; i-- {}

	if !is_digit(buf[i]) {
		i++
	}
	start := i
	val := 0
	for ; is_digit(buf[i]); i++ {
		val = 10*val + int(buf[i] - '0')
	}
	end := i-1

	return val, start, end
}

func Prob1() int {
	buf := make([]byte, len(input))
	copy(buf, input)

	cols := slices.Index(buf, '\n') + 1
	rows := len(buf) / cols

	sum := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			i := c + r*cols
			b := buf[i]
			if is_digit(b) || b == '.' || b == '\n' {
				continue
			}

			for r_s := r-1; r_s <= r+1; r_s++ {
				for c_s := c-1; c_s <= c+1; c_s++ {
					i_s := c_s + r_s*cols
					if i_s < 0 || i_s >= len(buf) || (r_s == r && c_s == c) {
						continue
					}

					b_s := buf[i_s]
					if is_digit(b_s) {
						val, start, end := parse_val(buf, i_s)
						sum += val

						for i_d := start; i_d <= end; i_d++ {
							buf[i_d] = '.'
						}

					}

				}
			}

		}
	}

	fmt.Println(sum)
	return sum
}

func Prob2() int {
	buf := make([]byte, len(input))
	copy(buf, input)

	cols := slices.Index(buf, '\n') + 1
	rows := len(buf) / cols

	sum := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			i := c + r*cols
			b := buf[i]
			if b != '*' {
				continue
			}

			neighbours := 0
			ratio := 1

			for r_s := r-1; r_s <= r+1; r_s++ {
				skip_past := -1
				for c_s := c-1; c_s <= c+1; c_s++ {
					i_s := c_s + r_s*cols
					if i_s < 0 || i_s >= len(buf) || (r_s == r && c_s == c) || i_s <= skip_past {
						continue
					}

					b_s := buf[i_s]
					if is_digit(b_s) {
						val, _, end := parse_val(buf, i_s)
						skip_past = end
						ratio *= val
						neighbours++
					}
				}
			}
			if neighbours == 2 {
				sum += ratio
			}
		}
	}

	fmt.Println(sum)
	return sum
}
