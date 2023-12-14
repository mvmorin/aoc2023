package day14

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Map struct {
	rows int
	cols int
	data []byte
}

func (m *Map) equal(a *Map) bool {
	for i := range m.data {
		if m.data[i] != a.data[i] {
			return false
		}
	}
	return true
}

func (m *Map) copy_map() Map {
	data := make([]byte, len(m.data))
	copy(data, m.data)
	return Map{m.rows, m.cols, data}
}

func (m *Map) get(i, j int) byte {
	if i < 0 || j < 0 || i >= m.rows || j >= m.cols {
		return ' '
	}
	return m.data[i+j*m.rows]
}

func (m *Map) set(i, j int, b byte) {
	if i < 0 || j < 0 || i >= m.rows || j >= m.cols {
		return
	}
	m.data[i+j*m.rows] = b
}

func (m *Map) print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%c", m.get(i, j))
		}
		fmt.Println()
	}
}

type Direction byte

const (
	NORTH Direction = iota
	WEST  Direction = iota
	SOUTH Direction = iota
	EAST  Direction = iota
)

func (m *Map) tilt(dir Direction) {
	var length int
	switch dir {
	case NORTH, SOUTH:
		length = m.cols
	case WEST, EAST:
		length = m.rows
	}

	for i := 0; i < length; i++ {
		m.tilt_single(i, dir)
	}
}

func (m *Map) tilt_single(col_or_row int, dir Direction) {
	get_row := func(i int) byte {
		return m.get(i, col_or_row)
	}
	set_row := func(i int, b byte) {
		m.set(i, col_or_row, b)
	}
	get_col := func(i int) byte {
		return m.get(col_or_row, i)
	}
	set_col := func(i int, b byte) {
		m.set(col_or_row, i, b)
	}

	var inc, i, length int
	var get func(int) byte
	var set func(int, byte)
	switch dir {
	case NORTH:
		inc, length, i = 1, m.rows, 0
		get, set = get_row, set_row
	case SOUTH:
		inc, length, i = -1, m.rows, m.rows-1
		get, set = get_row, set_row
	case WEST:
		inc, length, i = 1, m.cols, 0
		get, set = get_col, set_col
	case EAST:
		inc, length, i = -1, m.cols, m.cols-1
		get, set = get_col, set_col
	}

	for ; 0 <= i && i < length && get(i) != '.'; i += inc {
	}
	last_free := i
	for ; 0 <= i && i < length; i += inc {
		switch get(i) {
		case '#':
			i += inc
			for ; 0 <= i && i < length && get(i) != '.'; i += inc {
			}
			last_free = i
		case 'O':
			set(last_free, 'O')
			set(i, '.')
			last_free += inc
		}
	}
}

func (m *Map) cycle() {
	m.tilt(NORTH)
	m.tilt(WEST)
	m.tilt(SOUTH)
	m.tilt(EAST)
}

func (m *Map) total_north_load() int {
	load := 0
	for j := 0; j < m.cols; j++ {
		for i := 0; i < m.rows; i++ {
			if m.get(i, j) == 'O' {
				load += m.rows - i
			}
		}
	}
	return load
}

func parse_map(input string) Map {
	cols := strings.IndexByte(input, '\n')
	rows := len(input) / (cols + 1)
	m := Map{rows, cols, make([]byte, rows*cols)}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			k := i*(cols+1) + j
			m.set(i, j, input[k])
		}
	}
	return m
}

func Prob1() int {
	m := parse_map(input)
	m.tilt(NORTH)
	load := m.total_north_load()

	fmt.Println(load)
	return load
}

func Prob2() int {
	m := parse_map(input)
	path := make([]Map, 0, 256)
	path = append(path, m.copy_map())

	loop_index := -1
	for {
		m.cycle()
		for i := len(path) - 1; i >= 0; i-- {
			if m.equal(&path[i]) {
				loop_index = i
				break
			}
		}

		if loop_index < 0 {
			path = append(path, m.copy_map())
		} else {
			break
		}
	}

	n := 1000000000
	i := loop_index + (n-loop_index)%(len(path)-loop_index)
	load := path[i].total_north_load()

	fmt.Println(load)
	return load
}
