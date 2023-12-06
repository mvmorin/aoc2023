package day06

import (
	_"embed"
	"fmt"
	"strings"
	"strconv"
	"math"
)

//go:embed input.txt
var input string

func Prob1() int {
	parts := strings.Fields(input)

	times := make([]int, 0)
	distances := make([]int, 0)
	cur := &times
	for _, s := range parts[1:] {
		val, err := strconv.Atoi(s)
		if err != nil {
			cur = &distances
		} else {
			*cur = append(*cur, val)
		}
	}

	// t_hold*(time - t_hold) - dist > 0
	// t_hold*time - t_hold^2 - dist > 0
	// 0 > t_hold^2 + (-time)*t_hold + dist
	// t_hold_0 = time/2 +- ( time^2/4 - dist )^(1/2)
	// t_hold > time/2 - ( time^2/4 - dist )^(1/2)
	// t_hold < time/2 + ( time^2/4 - dist )^(1/2)

	total_n_possible := 1
	for i := 0; i < len(times); i++ {
		time := float64(times[i])
		dist := float64(distances[i])

		sq := math.Sqrt(time*time/4 - dist)
		t_hold_min_f := time/2 - sq
		t_hold_max_f := time/2 + sq
		t_hold_min_r := math.Ceil(t_hold_min_f)
		t_hold_max_r := math.Floor(t_hold_max_f)
		t_hold_min := int(t_hold_min_r)
		t_hold_max := int(t_hold_max_r)
		if !(t_hold_min_r - t_hold_min_f > 0) {
			t_hold_min++
		}
		if !(t_hold_max_f - t_hold_max_r > 0) {
			t_hold_max--
		}

		n_possible := t_hold_max - t_hold_min + 1
		total_n_possible *= n_possible
	}

	fmt.Println(total_n_possible)
	return total_n_possible
}

func Prob2() int {
	lines := strings.Split(input, "\n")

	time := 0
	for _, r := range lines[0] {
		if '0' <= r && r <= '9' {
			time = 10*time + int(r-'0')
		}
	}

	dist := 0
	for _, r := range lines[1] {
		if '0' <= r && r <= '9' {
			dist = 10*dist + int(r-'0')
		}
	}

	time_f := float64(time)
	dist_f := float64(dist)

	sq := math.Sqrt(time_f*time_f/4 - dist_f)
	t_hold_min_f := time_f/2 - sq
	t_hold_max_f := time_f/2 + sq
	t_hold_min_r := math.Ceil(t_hold_min_f)
	t_hold_max_r := math.Floor(t_hold_max_f)
	t_hold_min := int(t_hold_min_r)
	t_hold_max := int(t_hold_max_r)
	if !(t_hold_min_r - t_hold_min_f > 0) {
		t_hold_min++
	}
	if !(t_hold_max_f - t_hold_max_r > 0) {
		t_hold_max--
	}

	n_possible := t_hold_max - t_hold_min + 1

	fmt.Println(n_possible)
	return n_possible
}
