package common

import "strconv"

type Part int

const (
	Part0 Part = iota
	Part1
	Part2
)

func (p Part) String() string {
	return strconv.Itoa(int(p))
}
