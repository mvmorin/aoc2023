package day19

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var input_test string

type Part struct {
	x int
	m int
	a int
	s int
}

type Condition struct {
	field  byte
	less   bool
	val    int
	target string
}

func (c *Condition) eval(p Part) (string, bool) {
	if c.field == 't' {
		return c.target, true
	}

	var val int
	switch c.field {
	case 'x':
		val = p.x
	case 'm':
		val = p.m
	case 'a':
		val = p.a
	case 's':
		val = p.s
	}

	var valid bool
	if c.less {
		valid = val < c.val
	} else {
		valid = val > c.val
	}

	return c.target, valid
}

type Workflow struct {
	label string
	conds []Condition
}

func parse_part(s string) Part {
	part := Part{}

	for _, f := range strings.Split(s[1:len(s)-1], ",") {
		field := f[0]
		val, _ := strconv.Atoi(f[2:])

		switch field {
		case 'x':
			part.x = val
		case 'm':
			part.m = val
		case 'a':
			part.a = val
		case 's':
			part.s = val
		}
	}

	return part
}

func parse_workflow(s string) Workflow {
	i := 0
	for ; s[i] != '{'; i++ {
	}
	label := s[:i]

	conds := make([]Condition, 0)
	for _, line := range strings.Split(s[i+1:len(s)-1], ",") {
		j := strings.IndexByte(line, ':')

		if j < 0 {
			conds = append(conds, Condition{
				field:  't',
				target: line,
			})
		} else {
			val, _ := strconv.Atoi(line[2:j])
			conds = append(conds, Condition{
				field:  line[0],
				less: line[1] == '<',
				val: val,
				target: line[j+1:],
			})
		}
	}

	return Workflow{
		label,
		conds,
	}
}

func parse(s string) (map[string]Workflow, []Part) {
	scanner := bufio.NewScanner(strings.NewReader(s))

	workflows := make(map[string]Workflow)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		wf := parse_workflow(line)
		workflows[wf.label] = wf
	}

	parts := make([]Part, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		part := parse_part(line)
		parts = append(parts, part)
	}


	return workflows, parts
}

func Prob1() int {
	workflows, parts := parse(input)

	accepted := make([]Part,0)
	for _, part := range parts {
		label := "in"
		for label != "R" && label != "A" {
			wf := workflows[label]
			for _, c := range wf.conds {
				new_label, valid := c.eval(part)

				if valid {
					label = new_label
					break
				}
			}
		}

		if label == "A" {
			accepted = append(accepted, part)
		}
	}

	sum := 0
	for _, part := range accepted {
		sum += part.x
		sum += part.m
		sum += part.a
		sum += part.s
	}
	fmt.Println(sum)
	return sum
}

func Prob2() int {
	fmt.Println(input_test)
	return 0
}
