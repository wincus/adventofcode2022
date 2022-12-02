package day2

import (
	"testing"

	"github.com/wincus/adventofcode2022/internal/common"
)

type Test struct {
	input []string
	p     common.Part
	want  int
}

func TestSolver(t *testing.T) {

	tests := []Test{
		{
			input: []string{
				"A Y",
				"B X",
				"C Z",
			},
			p:    common.Part1,
			want: 15,
		},
		{
			input: []string{
				"A Y",
				"B X",
				"C Z",
			},
			p:    common.Part2,
			want: 12,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
