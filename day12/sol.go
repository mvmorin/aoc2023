package day12

import (
	_ "embed"
	"fmt"
	"strings"
	"bufio"
	"strconv"
)

//go:embed input.txt
var input string

func arrangements(line string, mult int) int {
	parts := strings.FieldsFunc(line, func(r rune)bool {
		switch r {
		case ',', ' ':
			return true
		default:
			return false
		}
	})

	var sb strings.Builder
	for i := 0; i < mult; i++ {
		sb.WriteString(parts[0])
		if i < mult-1 {
			sb.WriteByte('?')
		}
	}
	row := sb.String()

	seg_str := parts[1:]
	offsets := make([]int, mult*len(seg_str))
	segs := make([]int, mult*len(seg_str))
	tot_segs_len := 0
	for i := 0; i < len(segs); i++ {
		segs[i], _ = strconv.Atoi(seg_str[i%len(seg_str)])
		tot_segs_len += segs[i]+1
	}
	tot_segs_len--
	// fmt.Println(row, segs, offsets, tot_segs_len)

	max_offset := len(row)-tot_segs_len
	sum := 0
	for {
		is_valid := true
		i := 0
		for j := range segs {
			seg_start := i+offsets[j]
			seg_end := seg_start + segs[j]

			is_valid = is_valid && seg_end <= len(row)
			for ; is_valid && i < seg_start; i++ {
				is_valid = is_valid && (row[i] == '.' || row[i] == '?')
			}
			for ; is_valid && i < seg_end; i++ {
				is_valid = is_valid && (row[i] == '#' || row[i] == '?')
			}
			if is_valid && i < len(row) {
				is_valid = is_valid && (row[i] == '.' || row[i] == '?')
				i++
			}
		}
		for ; is_valid && i < len(row); i++ {
			is_valid = is_valid && (row[i] == '.' || row[i] == '?')
		}

		if is_valid {
			sum++
		}

		i = 0
		offsets[0]++
		for ;i < len(offsets) && offsets[i] > max_offset; i++ {
			offsets[i] = 0
			if i+1 < len(offsets) {
				offsets[i+1]++
			}
		}
		if i == len(offsets) {
			break
		}
	}

	return sum
}


func Prob1() int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += arrangements(line, 1)
	}
	fmt.Println(sum)
	return sum
}

func Prob2() int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += arrangements(line,5)
	}
	fmt.Println(sum)
	return 0
}
