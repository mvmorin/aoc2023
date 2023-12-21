package day17

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Queue[T any] struct {
	start  int
	length int
	slice  []T
}

func new_queue[T any]() Queue[T] {
	return Queue[T]{0, 0, make([]T, 0, 8192)}
}

func (q *Queue[T]) push(e T) {
	if q.length == cap(q.slice) {
		new_slice := make([]T, 0, 2*cap(q.slice))

		for i := q.start; i < len(q.slice); i++ {
			new_slice = append(new_slice, q.slice[i])
		}
		for i := 0; i < q.start; i++ {
			new_slice = append(new_slice, q.slice[i])
		}

		q.start = 0
		q.slice = new_slice
	}

	ins_i := (q.start + q.length) % cap(q.slice)
	if ins_i < len(q.slice) {
		q.slice[ins_i] = e
	} else {
		q.slice = append(q.slice, e)
	}
	q.length++
}

func (q *Queue[T]) peek() T {
	return q.slice[q.start]
}

func (q *Queue[T]) pop() T {
	e := q.peek()
	q.start = (q.start + 1) % cap(q.slice)
	q.length--
	return e
}

type Map struct {
	rows int
	cols int
	data string
}

func (m *Map) valid_index(i, j int) bool {
	return !(i < 0 || j < 0 || i >= m.rows || j >= m.cols)
}

func (m *Map) get(i, j int) byte {
	return m.data[i*(m.cols+1)+j]
}

func (m *Map) print_map() {
	fmt.Println(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%c", m.get(i, j))
		}
		fmt.Println()
	}
}

func parse_map(input string) Map {
	cols := strings.Index(input, "\n")
	rows := len(input) / (cols + 1)
	return Map{rows, cols, input}
}

type Direction byte

const (
	UP    Direction = iota
	DOWN  Direction = iota
	LEFT  Direction = iota
	RIGHT Direction = iota
)

func next_coord(i, j int, dir Direction, l int) (int, int) {
	switch dir {
	case UP:
		return i - l, j
	case DOWN:
		return i + l, j
	case LEFT:
		return i, j - l
	case RIGHT:
		return i, j + l
	}
	panic("invalid")
}

type State struct {
	i            int
	j            int
	straight_len int
	cost         int
	dir          Direction
}

type StateHash struct {
	i            int
	j            int
	straight_len int
	dir          Direction
}

func Prob1() int {
	m := parse_map(input)

	q := new_queue[State]()
	q.push(State{0, 0, 0, 0, RIGHT})
	q.push(State{0, 0, 0, 0, DOWN})

	best := make(map[StateHash]int)
	for q.length > 0 {
		s := q.pop()

		state_hash := StateHash{s.i, s.j, s.straight_len, s.dir}
		state_best, ok := best[state_hash]
		if !ok {
			best[state_hash] = s.cost
		} else if state_best > s.cost {
			best[state_hash] = s.cost
		} else {
			continue
		}

		if s.i == m.rows-1 && s.j == m.cols-1 {
			continue
		}

		var dirs [3]Direction
		dirs[0] = s.dir
		switch s.dir {
		case UP, DOWN:
			dirs[1] = LEFT
			dirs[2] = RIGHT
		case LEFT, RIGHT:
			dirs[1] = UP
			dirs[2] = DOWN
		}

		for _, d := range dirs {
			i, j := next_coord(s.i, s.j, d, 1)
			if !m.valid_index(i, j) || (d == s.dir && s.straight_len >= 2) {
				continue
			}

			step_cost := int(m.get(i, j) - '0')
			new_path := State{i: i, j: j, cost: s.cost + step_cost, dir: d}
			if d == s.dir {
				new_path.straight_len = s.straight_len + 1
			} else {
				new_path.straight_len = 0
			}

			q.push(new_path)
		}
	}

	goal_hashes := [...]StateHash{
		{m.rows-1, m.cols-1, 0, DOWN},
		{m.rows-1, m.cols-1, 0, RIGHT},
		{m.rows-1, m.cols-1, 1, DOWN},
		{m.rows-1, m.cols-1, 1, RIGHT},
		{m.rows-1, m.cols-1, 2, DOWN},
		{m.rows-1, m.cols-1, 2, RIGHT},
	}

	best_cost := -1
	for _, state_hash := range goal_hashes {
		b, ok := best[state_hash]

		if ok && (best_cost < 0 || best_cost > b) {
			best_cost = b
		}
	}
	fmt.Println(best_cost)
	return best_cost
}

func Prob2() int {
	q := new_queue[int]()
	fmt.Println(q.start, q.length, q.slice)

	i := 0
	for ; i < 3; i++ {
		q.push(i)
		fmt.Println(q.start, q.length, q.slice, q.peek())
	}
	for ; i < 20; i++ {
		q.push(i)
		q.pop()
		fmt.Println(q.start, q.length, q.slice, q.peek())
	}
	for q.length > 0 {
		res := q.pop()
		fmt.Println(q.start, q.length, q.slice, res)
	}
	return 0
}
