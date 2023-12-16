package day16

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Direction byte

const (
	UP    Direction = 0
	DOWN  Direction = 1
	RIGHT Direction = 2
	LEFT  Direction = 3
)

type Map struct {
	rows int
	cols int
	data string
}

func (m *Map) get(i, j int) byte {
	if i < 0 || j < 0 || i >= m.rows || j >= m.cols {
		return ' '
	}
	return m.data[i*(m.cols+1)+j]
}

func to_map(input string) Map {
	cols := strings.Index(input, "\n")
	rows := len(input) / (cols + 1)
	return Map{rows, cols, input}
}

type Visited struct {
	rows int
	cols int
	data [][4]bool
}

func (v *Visited) is_visited(i, j int, dir Direction) bool {
	return v.data[i*v.cols+j][dir]
}

func (v *Visited) any_visited(i, j int) bool {
	for d := UP; d <= LEFT; d++ {
		if v.is_visited(i, j, d) {
			return true
		}
	}
	return false
}

func (v *Visited) visit(i, j int, dir Direction) {
	v.data[i*v.cols+j][dir] = true
}

func new_visited(rows, cols int) Visited {
	data := make([][4]bool, rows*cols)
	return Visited{rows, cols, data}
}

type Beam struct {
	i   int
	j   int
	dir Direction
}

func next(i, j int, dir Direction) (int, int) {
	switch dir {
	case UP:
		return i - 1, j
	case DOWN:
		return i + 1, j
	case LEFT:
		return i, j - 1
	case RIGHT:
		return i, j + 1
	}
	panic("invalid")
}

func count_visited(m Map, row, col int, dir Direction) int {
	v := new_visited(m.rows, m.cols)

	front := make([]Beam, 0, 1024)
	front = append(front, Beam{row, col, dir})
	for len(front) > 0 {
		b := front[len(front)-1]
		front = front[:len(front)-1]

		tile := m.get(b.i, b.j)
		if tile == ' ' || v.is_visited(b.i, b.j, b.dir) {
			// off map or already visited
			continue
		}

		v.visit(b.i, b.j, b.dir)

		switch tile {
		case '.':
			i, j := next(b.i, b.j, b.dir)
			front = append(front, Beam{i, j, b.dir})
		case '|':
			if b.dir == UP || b.dir == DOWN {
				i, j := next(b.i, b.j, b.dir)
				front = append(front, Beam{i, j, b.dir})
			} else {
				i, j := next(b.i, b.j, UP)
				front = append(front, Beam{i, j, UP})
				i, j = next(b.i, b.j, DOWN)
				front = append(front, Beam{i, j, DOWN})
			}
		case '-':
			if b.dir == LEFT || b.dir == RIGHT {
				i, j := next(b.i, b.j, b.dir)
				front = append(front, Beam{i, j, b.dir})
			} else {
				i, j := next(b.i, b.j, LEFT)
				front = append(front, Beam{i, j, LEFT})
				i, j = next(b.i, b.j, RIGHT)
				front = append(front, Beam{i, j, RIGHT})
			}
		case '\\':
			var next_dir Direction
			switch b.dir {
			case UP:
				next_dir = LEFT
			case DOWN:
				next_dir = RIGHT
			case LEFT:
				next_dir = UP
			case RIGHT:
				next_dir = DOWN
			}
			i, j := next(b.i, b.j, next_dir)
			front = append(front, Beam{i, j, next_dir})
		case '/':
			var next_dir Direction
			switch b.dir {
			case UP:
				next_dir = RIGHT
			case DOWN:
				next_dir = LEFT
			case LEFT:
				next_dir = DOWN
			case RIGHT:
				next_dir = UP
			}
			i, j := next(b.i, b.j, next_dir)
			front = append(front, Beam{i, j, next_dir})
		}

	}

	sum := 0
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if v.any_visited(i,j) {
				sum++
			}
		}
	}

	return sum
}

func Prob1() int {
	m := to_map(input)
	sum := count_visited(m, 0, 0, RIGHT)
	fmt.Println(sum)
	return sum
}

func Prob2() int {
	m := to_map(input)
	best := 0
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if i == 0 {
				best = max(best, count_visited(m, i, j, DOWN))
			}
			if i == m.rows-1 {
				best = max(best, count_visited(m, i, j, UP))
			}
			if j == 0 {
				best = max(best, count_visited(m, i, j, RIGHT))
			}
			if j == m.cols-1 {
				best = max(best, count_visited(m, i, j, RIGHT))
			}
		}
	}
	fmt.Println(best)
	return best

}
