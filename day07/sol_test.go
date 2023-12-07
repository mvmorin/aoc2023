package day07

import (
	"testing"
)

func Test1(t *testing.T) {
	res := Prob1()
	correct := 253205868
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}

func Test2(t *testing.T) {
	res := Prob2()
	correct := 253907829
	if res != correct {
		t.Errorf("Wrong result %d, should be %d", res, correct)
	}
}
