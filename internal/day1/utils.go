package day1

import (
	"log"
	"sort"
	"strconv"

	"github.com/wincus/adventofcode2022/internal/common"
)

// Solve returns the solutions for day 1
func Solve(s []string, p common.Part) int {

	var calories []int
	var sum int

	for i, line := range s {

		if len(line) == 0 {
			calories = append(calories, sum)
			sum = 0
			continue
		}

		v, err := strconv.Atoi(line)

		if err != nil {
			log.Printf("could not convert string: %v to int", line)
			return 0
		}

		// ensures last line is counted
		if i == len(s)-1 {
			calories = append(calories, v)
			continue
		}

		sum += v
	}

	switch p {
	case common.Part1:
		return getMaxSum(calories, 1)
	case common.Part2:
		return getMaxSum(calories, 3)
	default:
		log.Printf("Part %v not supported", p)
		return 0
	}
}

func getMaxSum(calories []int, top int) int {

	var sum int

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	for i := 0; i < top; i++ {
		sum += calories[i]
	}

	return sum
}
