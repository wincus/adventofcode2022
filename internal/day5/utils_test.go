package day5

import (
	"testing"

	"github.com/wincus/adventofcode2022/internal/common"
)

type Test struct {
	input []string
	p     common.Part
	want  string
}

func TestSolver(t *testing.T) {

	tests := []Test{
		{
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3 ",
				"",
				"move 1 from 2 to 1",
				"move 3 from 1 to 3",
				"move 2 from 2 to 1",
				"move 1 from 1 to 2",
			},
			p:    common.Part1,
			want: "CMZ",
		},
		{
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3 ",
				"",
				"move 1 from 2 to 1",
				"move 3 from 1 to 3",
				"move 2 from 2 to 1",
				"move 1 from 1 to 2",
			},
			p:    common.Part2,
			want: "MCD",
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
