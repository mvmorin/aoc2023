package day08

import (
	"testing"
)

func Test1(t *testing.T) {
	res := Prob1()
	correct := 16697
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}

func Test2(t *testing.T) {
	res := Prob2()
	correct := 10668805667831
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}
