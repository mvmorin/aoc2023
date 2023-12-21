package day24

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var input_test string

func Prob1() int {
	fmt.Println(input)
	return 0
}

func Prob2() int {
	fmt.Println(input_test)
	return 0
}
