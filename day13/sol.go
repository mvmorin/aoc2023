package day13

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// TODO: add a TransposedMap struct that contains a Map and make the reflection functions generic so both can be
// accepted

type Map struct {
	rows int
	cols int
	data []byte
}

func (m *Map) get(i, j int) byte {
	if i < 0 || j < 0 || i >= m.rows || j >= m.cols {
		return ' '
	}
	return m.data[i*m.cols+j]
}

func parse_maps(input string) []Map {
	maps := make([]Map, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	new_map := true
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			new_map = true
			continue

		} else if new_map {
			cols := len(line)
			m := Map{
				rows: 0,
				cols: cols,
				data: make([]byte, 0),
			}
			maps = append(maps, m)
			new_map = false
		}

		m := &maps[len(maps)-1]
		m.rows++
		for i := 0; i < len(line); i++ {
			m.data = append(m.data, line[i])
		}

	}
	return maps
}

func vertical_reflection(m Map) (int, bool) {
	ref_col := 0
	var is_reflection bool
	for ; ref_col < m.cols-1; ref_col++ {
		is_reflection = true
		for col_off := 0; is_reflection && 0 <= ref_col-col_off && ref_col+col_off+1 < m.cols; col_off++ {
			for i := 0; is_reflection && i < m.rows; i++ {
				is_reflection = is_reflection && m.get(i, ref_col-col_off) == m.get(i, ref_col+col_off+1)
			}
		}

		if is_reflection {
			break
		}
	}

	return ref_col + 1, is_reflection
}

func horizontal_reflection(m Map) (int, bool) {
	ref_row := 0
	var is_reflection bool
	for ; ref_row < m.rows-1; ref_row++ {
		is_reflection = true
		for row_off := 0; is_reflection && 0 <= ref_row-row_off && ref_row+row_off+1 < m.rows; row_off++ {
			for j := 0; is_reflection && j < m.cols; j++ {
				is_reflection = is_reflection && m.get(ref_row-row_off, j) == m.get(ref_row+row_off+1, j)
			}
		}

		if is_reflection {
			break
		}
	}

	return ref_row + 1, is_reflection
}

func Prob1() int {
	maps := parse_maps(input)

	sum := 0
	for _, m := range maps {
		cols_before, col_reflection := vertical_reflection(m)
		rows_before, row_reflection := horizontal_reflection(m)

		if col_reflection {
			sum += cols_before
		} else if row_reflection {
			sum += rows_before*100
		}
	}
	fmt.Println(sum)
	return sum
}

func vertical_reflection_smudge(m Map) (int, bool) {
	ref_col := 0
	var is_reflection bool
	for ; ref_col < m.cols-1; ref_col++ {
		non_matching := 0
		for col_off := 0; 0 <= ref_col-col_off && ref_col+col_off+1 < m.cols; col_off++ {
			for i := 0; i < m.rows; i++ {
				if m.get(i, ref_col-col_off) != m.get(i, ref_col+col_off+1) {
					non_matching++
				}
			}
		}

		is_reflection = non_matching == 1

		if is_reflection {
			break
		}
	}

	return ref_col + 1, is_reflection
}

func horizontal_reflection_smudge(m Map) (int, bool) {
	ref_row := 0
	var is_reflection bool
	for ; ref_row < m.rows-1; ref_row++ {
		non_matching := 0
		for row_off := 0; 0 <= ref_row-row_off && ref_row+row_off+1 < m.rows; row_off++ {
			for j := 0; j < m.cols; j++ {
				if m.get(ref_row-row_off, j) != m.get(ref_row+row_off+1, j) {
					non_matching++
				}
			}
		}

		is_reflection = non_matching == 1

		if is_reflection {
			break
		}
	}

	return ref_row + 1, is_reflection
}

func Prob2() int {
	maps := parse_maps(input)

	sum := 0
	for _, m := range maps {
		cols_before, col_reflection := vertical_reflection_smudge(m)
		rows_before, row_reflection := horizontal_reflection_smudge(m)

		if col_reflection {
			sum += cols_before
		} else if row_reflection {
			sum += rows_before*100
		}
	}
	fmt.Println(sum)
	return sum
}
