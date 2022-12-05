package day5

import (
	"bytes"
	"fmt"
	"unicode"

	"github.com/wincus/adventofcode2022/internal/common"
)

type stack []rune

type rule struct {
	from int
	to   int
	qty  int
}

// Solve returns the solutions for day 5
func Solve(s []string, p common.Part) string {

	st := getStacks(s)
	r := getRules(s)

	if p == common.Part1 {
		process(st, r)
	}

	if p == common.Part2 {
		processInBatches(st, r)
	}

	return getTop(st)
}

func getTop(s map[int]stack) string {

	var top bytes.Buffer

	for i := 1; i <= len(s); i++ {
		_, fromItem := popN(s[i], 1)
		top.WriteRune(fromItem[0])
	}

	return top.String()

}

func process(st map[int]stack, r []rule) {

	for _, rule := range r {
		for i := 1; i <= rule.qty; i++ {
			var m []rune
			st[rule.from], m = popN(st[rule.from], 1)
			st[rule.to] = pushN(st[rule.to], m)
		}
	}
}

func processInBatches(st map[int]stack, r []rule) {

	for _, rule := range r {
		var m []rune
		st[rule.from], m = popN(st[rule.from], rule.qty)
		st[rule.to] = pushN(st[rule.to], m)
	}
}

func getRules(s []string) []rule {

	var rules []rule

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		if line[0] == 'm' {

			var r rule

			fmt.Sscanf(line, "move %d from %d to %d", &r.qty, &r.from, &r.to)

			rules = append(rules, r)
		}
	}

	return rules
}

func getStacks(s []string) map[int]stack {

	st := make(map[int]stack)

	for _, line := range s {

		if len(line) == 0 {
			break
		}

		for i := 0; i < len(line); i++ {
			if (i-1)%4 == 0 {
				if unicode.IsLetter(rune(line[i])) {
					st[(i/4)+1] = append(st[(i/4)+1], rune(line[i]))
				}
			}
		}
	}

	// reverse is needed cause we are parsing the file from top
	// to bottom .... and the crates are stacked from bootom to top
	// ( that is the firt one we see is the one in the top )
	for _, v := range st {
		v.reverse()
	}

	return st
}

func (s *stack) reverse() {
	for i := 0; i < len(*s)/2; i++ {
		j := len(*s) - i - 1
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func pushN(s stack, r []rune) stack {
	return append(s, r...)
}

func popN(s stack, n int) (stack, []rune) {

	r := s[len(s)-n:]

	s = s[:len(s)-n]

	return s, r
}
