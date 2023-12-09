package day09

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func predic_line(line string, reverse bool) int {
	strs := strings.Fields(line)
	numbers := make([]int, len(strs))
	for i := range strs {
		dest := i
		if reverse {
			dest = len(numbers)-i-1
		}
		numbers[dest], _ = strconv.Atoi(strs[i])
	}

	non_zero_found := true
	for n := len(numbers); n > 0 && non_zero_found; n-- {
		for i := 0; i+1 < n; i++ {
			cur := i
			next := i+1
			numbers[cur] = numbers[next] - numbers[cur]
			non_zero_found = non_zero_found || numbers[cur] != 0
		}
	}

	sum := 0
	for _, x := range numbers {
		sum += x
	}

	return sum
}

func Prob1() int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += predic_line(line, false)
	}

	fmt.Println(sum)
	return sum
}

func Prob2() int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += predic_line(line, true)
	}

	fmt.Println(sum)
	return sum
}
