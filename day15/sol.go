package day15

import (
	_ "embed"
	"fmt"
	"strings"
	"slices"
)

//go:embed input.txt
var input string

func hash(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h += int(s[i])
		h *= 17
		h %= 256
	}
	return h
}

type Entry struct {
	focal int
	label string
}

type HashMap struct {
	data [256][]Entry
}

func (hm *HashMap) add_step(s string) {
	var cmd byte
	var label string
	var focal int
	if s[len(s)-2] == '=' {
		label = s[:len(s)-2]
		cmd = '='
		focal = int(s[len(s)-1] - '0')
	} else {
		label = s[:len(s)-1]
		cmd = '-'
		focal = 0
	}

	h := hash(label)
	if hm.data[h] == nil {
		hm.data[h] = make([]Entry, 0, 10)
	}

	bi := slices.IndexFunc(hm.data[h], func (e Entry) bool {
		return e.label == label
	})

	switch cmd {
	case '-':
		if bi >= 0 {
			hm.data[h] = slices.Delete(hm.data[h], bi, bi+1)
		}
	case '=':
		e := Entry{focal, label}
		if bi >= 0 {
			hm.data[h][bi] = e
		} else {
			hm.data[h] = append(hm.data[h], e)
		}
	}
}

func (hm *HashMap) focusing_power() int {
	sum := 0
	for i := range hm.data {
		for j := range hm.data[i] {
			sum += (i+1)*(j+1)*hm.data[i][j].focal
		}
	}
	return sum
}

func Prob1() int {
	parts := strings.Split(strings.TrimSpace(input), ",")
	sum := 0
	for _, s := range parts {
		sum += hash(s)
	}
	fmt.Println(sum)
	return sum
}

func Prob2() int {
	parts := strings.Split(strings.TrimSpace(input), ",")
	hm := HashMap{}
	for _, s := range parts {
		hm.add_step(s)
	}
	fp := hm.focusing_power()
	fmt.Println(fp)
	return fp

}
