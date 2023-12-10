package day10

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

type Idx struct {
	i int
	j int
}

type Map struct {
	rows int
	cols int
	data []byte
}

func (m *Map) get(pos Idx) byte {
	if pos.i < 0 || pos.j < 0 || pos.i >= m.rows || pos.j >= m.cols {
		return ' '
	}
	return m.data[pos.i*m.cols+pos.j]
}

func (m *Map) set(pos Idx, b byte) bool {
	if pos.i < 0 || pos.j < 0 || pos.i >= m.rows || pos.j >= m.cols {
		return false
	}
	m.data[pos.i*m.cols+pos.j] = b
	return true
}

func (m *Map) connections(pos Idx) (Idx, Idx) {
	b := m.get(pos)
	switch b {
	case '|':
		return Idx{pos.i + 1, pos.j}, Idx{pos.i - 1, pos.j}
	case '-':
		return Idx{pos.i, pos.j + 1}, Idx{pos.i, pos.j - 1}
	case 'L':
		return Idx{pos.i - 1, pos.j}, Idx{pos.i, pos.j + 1}
	case 'J':
		return Idx{pos.i - 1, pos.j}, Idx{pos.i, pos.j - 1}
	case '7':
		return Idx{pos.i + 1, pos.j}, Idx{pos.i, pos.j - 1}
	case 'F':
		return Idx{pos.i + 1, pos.j}, Idx{pos.i, pos.j + 1}
	}
	return pos, pos
}

func parse_map(input string) (Map, Idx, Idx) {
	cols := 0
	for ; cols < len(input) && input[cols] != '\n'; cols++ {
	}
	rows := len(input) / (cols + 1)
	data := make([]byte, cols*rows)
	start := Idx{}

	dest := 0
	for i := 0; i < len(input); i++ {
		b := input[i]
		if b == '\n' {
			continue
		}
		if b == 'S' {
			start.i = dest / cols
			start.j = dest - start.i*cols
		}
		data[dest] = b
		dest++
	}
	m := Map{rows, cols, data}

	first_tunnel := Idx{}
	ds := [4]struct{ i, j int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, d := range ds {
		test := Idx{start.i + d.i, start.j + d.j}

		a, b := m.connections(test)
		if (a.i == start.i && a.j == start.j) || (b.i == start.i && b.j == start.j) {
			first_tunnel = test
		}
	}

	return m, start, first_tunnel
}

func Prob1() int {
	m, start, first_tunnel := parse_map(input)

	length := 1
	prev_idx := start
	cur_idx := first_tunnel
	for m.get(cur_idx) != 'S' {
		a_idx, b_idx := m.connections(cur_idx)
		length++

		if a_idx != prev_idx {
			prev_idx = cur_idx
			cur_idx = a_idx
		} else {
			prev_idx = cur_idx
			cur_idx = b_idx
		}
	}
	furthest := length / 2
	fmt.Println(furthest)
	return furthest
}

func segments_sides(m *Map, front *[]Idx, start, end Idx) {
	d_i := -(end.j - start.j)
	d_j := end.i - start.i

	idx := Idx{end.i + d_i, end.j + d_j}
	if m.get(idx) == '.' {
		m.set(idx, 'X')
		*front = append(*front, idx)
	}
	idx = Idx{end.i - d_i, end.j - d_j}
	if m.get(idx) == '.' {
		m.set(idx, 'Y')
		*front = append(*front, idx)
	}
	idx = Idx{start.i + d_i, start.j + d_j}
	if m.get(idx) == '.' {
		m.set(idx, 'X')
		*front = append(*front, idx)
	}
	idx = Idx{start.i - d_i, start.j - d_j}
	if m.get(idx) == '.' {
		m.set(idx, 'Y')
		*front = append(*front, idx)
	}
}

func Prob2() int {
	m, start, first_tunnel := parse_map(input)

	clean_data := make([]byte, len(m.data))
	for i := range clean_data {
		clean_data[i] = '.'
	}
	m_flood := Map{rows: m.rows, cols: m.cols, data: clean_data}
	m_flood.set(start, 'S')

	// find loop and copy it over to a clean map
	prev_idx := start
	cur_idx := first_tunnel
	for m.get(cur_idx) != 'S' {
		m_flood.set(cur_idx, m.get(cur_idx))
		a_idx, b_idx := m.connections(cur_idx)
		if a_idx != prev_idx {
			prev_idx = cur_idx
			cur_idx = a_idx
		} else {
			prev_idx = cur_idx
			cur_idx = b_idx
		}
	}

	// color the tiles of both sides of the loop segments with different colors
	expand_front := make([]Idx, 0)
	prev_idx = start
	cur_idx = first_tunnel
	for m.get(cur_idx) != 'S' {
		segments_sides(&m_flood, &expand_front, prev_idx, cur_idx)

		a_idx, b_idx := m.connections(cur_idx)
		if a_idx != prev_idx {
			prev_idx = cur_idx
			cur_idx = a_idx
		} else {
			prev_idx = cur_idx
			cur_idx = b_idx
		}
	}
	segments_sides(&m_flood, &expand_front, prev_idx, cur_idx)

	// expand each colored tiles until there is nothing left to color
	for len(expand_front) > 0 {
		last_i := len(expand_front) - 1
		cur := expand_front[last_i]
		expand_front = expand_front[:last_i]

		ds := [...]struct{ i, j int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		for _, d := range ds {
			test := Idx{cur.i + d.i, cur.j + d.j}

			if m_flood.get(test) == '.' {
				if m_flood.set(test, m_flood.get(cur)) {
					expand_front = append(expand_front, test)
				}
			}
		}
	}

	// count the color tiles and identify which color touches the sides
	x_is_outside := true
	x_cnt := 0
	y_cnt := 0
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			b := m_flood.get(Idx{i, j})
			if b == 'X' {
				x_cnt++
			} else if b == 'Y' {
				y_cnt++
			}

			if b == 'Y' && (i == 0 || j == 0 || i == m.rows-1 || j == m.cols-1) {
				x_is_outside = false
			}
		}
	}
	cnt := x_cnt
	if x_is_outside {
		cnt = y_cnt
	}

	fmt.Println(cnt)
	return cnt
}
