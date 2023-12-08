package day08

import (
	_ "embed"
	"fmt"
	"slices"
)

//go:embed input.txt
var input []byte

type Instructions struct {
	data []byte
	i    int
}

func (insts *Instructions) next() byte {
	ins := insts.data[insts.i]
	insts.i++
	if insts.i >= len(insts.data) {
		insts.i = 0
	}
	return ins
}

func (insts *Instructions) reset() {
	insts.i = 0
}

type Fork[T any] struct {
	left  T
	right T
}

func (f *Fork[T]) get(ins byte) T {
	switch ins {
	case 'L':
		return f.left
	case 'R':
		return f.right
	}
	panic("invalid instruction")
}

func parse(buf []byte) (Instructions, map[string]Fork[string], []string) {
	i := slices.Index(buf, '\n')
	insts := Instructions{
		data: buf[:i],
		i:    0,
	}
	i += 2

	graph := make(map[string]Fork[string])
	starts := make([]string, 0)
	for i < len(buf) {
		node := string(buf[i : i+3])
		i += 7
		left := string(buf[i : i+3])
		i += 5
		right := string(buf[i : i+3])
		i += 5
		graph[node] = Fork[string]{left, right}

		if node[2] == 'A' {
			starts = append(starts, node)
		}
	}

	return insts, graph, starts
}

func Prob1() int {
	insts, graph, _ := parse(input)
	cur := "AAA"
	steps := 0
	for cur != "ZZZ" {
		f := graph[cur]
		cur = f.get(insts.next())
		steps++
	}
	fmt.Println(steps)
	return steps
}

func lcm(ns []int) int {
	is_prime := func(p int) bool {
		if p == 1 {
			return false
		}

		yes := !(p % 2 == 0)
		for t := 3; yes && t < p/2; t += 2 {
			yes = !(p % t == 0)
		}
		return yes
	}

	prime_factors := make(map[int]int)
	for _, n := range ns {
		for p := 2; p <= n; p++ {
			if !is_prime(p) {
				continue
			}

			pow := 0
			for n % p == 0 {
				n = n / p
				pow++
			}
			prev_mult, _ := prime_factors[p]
			if pow > prev_mult {
				prime_factors[p] = pow
			}
		}
	}

	res := 1
	for p, pow := range prime_factors {
		for i := 0; i < pow; i++ {
			res *= p
		}
	}

	return res
}

func Prob2() int {
	insts, graph, starts := parse(input)
	periods := make([]int, len(starts))
	for i, start := range starts {
		insts.reset()
		cur := start
		steps := 0
		for cur[2] != 'Z' {
			f := graph[cur]
			cur = f.get(insts.next())
			steps++
		}
		periods[i] = steps
	}

	res := lcm(periods)
	fmt.Println(res)
	return res
}
