package day05

import (
	"testing"
)

func Test1(t *testing.T) {
	res := Prob1()
	correct := 3374647
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}

func Test2(t *testing.T) {
	res := Prob2()
	correct := 6082852
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}
