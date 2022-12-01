package day1

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
				"1000",
				"2000",
				"3000",
				"",
				"4000",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000",
			},
			p:    common.Part1,
			want: 24000,
		},
		{
			input: []string{
				"1000",
				"2000",
				"3000",
				"",
				"4000",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000",
			},
			p:    common.Part2,
			want: 45000,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
