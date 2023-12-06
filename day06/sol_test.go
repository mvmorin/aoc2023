package day06

import (
	"testing"
)

func Test1(t *testing.T) {
	res := Prob1()
	correct := 449820
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}

func Test2(t *testing.T) {
	res := Prob2()
	correct := 42250895
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}
