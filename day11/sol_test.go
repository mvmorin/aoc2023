package day11

import (
	"testing"
)

func Test1(t *testing.T) {
	res := Prob1()
	correct := 9536038
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}

func Test2(t *testing.T) {
	res := Prob2()
	correct := 447744640566
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}
