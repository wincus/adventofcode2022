package day6

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
				"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			p:    common.Part1,
			want: 7,
		},
		{
			input: []string{
				"bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			p:    common.Part1,
			want: 5,
		},
		{
			input: []string{
				"nppdvjthqldpwncqszvftbrmjlhg",
			},
			p:    common.Part1,
			want: 6,
		},
		{
			input: []string{
				"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			p:    common.Part1,
			want: 10,
		},
		{
			input: []string{
				"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			p:    common.Part1,
			want: 11,
		},
		{
			input: []string{
				"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			p:    common.Part2,
			want: 19,
		},
		{
			input: []string{
				"bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			p:    common.Part2,
			want: 23,
		},
		{
			input: []string{
				"nppdvjthqldpwncqszvftbrmjlhg",
			},
			p:    common.Part2,
			want: 23,
		},
		{
			input: []string{
				"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			p:    common.Part2,
			want: 29,
		},
		{
			input: []string{
				"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			p:    common.Part2,
			want: 26,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
