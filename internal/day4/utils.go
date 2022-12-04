package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2022/internal/common"
)

// Solve returns the solutions for day 4
func Solve(s []string, p common.Part) int {

	var count int

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		first, second, err := getPairs(line)

		if err != nil {
			fmt.Println(err)
			return 0
		}

		if p == common.Part1 {
			if pairsFullOverlap(first, second) {
				count++
			}
		}

		if p == common.Part2 {
			if pairsPartialOverlap(first, second) {
				count++
			}
		}
	}

	return count
}

func pairsPartialOverlap(first, second [2]int) bool {

	if first[0] <= second[0] && first[1] >= second[0] {
		return true
	}

	if second[0] <= first[0] && second[1] >= first[0] {
		return true
	}

	return false
}

func pairsFullOverlap(first, second [2]int) bool {

	if first[0] >= second[0] && first[1] <= second[1] {
		return true
	}

	if second[0] >= first[0] && second[1] <= first[1] {
		return true
	}

	return false
}

func getPairs(s string) ([2]int, [2]int, error) {

	pairs := strings.Split(s, ",")

	if len(pairs) != 2 {
		return [2]int{}, [2]int{}, fmt.Errorf("could not parse pair")
	}

	firstPair := strings.Split(pairs[0], "-")

	if len(firstPair) != 2 {
		return [2]int{}, [2]int{}, fmt.Errorf("could not parse pair")
	}

	secondPair := strings.Split(pairs[1], "-")

	if len(secondPair) != 2 {
		return [2]int{}, [2]int{}, fmt.Errorf("could not parse pair")
	}

	firstStart, err := strconv.Atoi(firstPair[0])

	if err != nil {
		return [2]int{}, [2]int{}, fmt.Errorf("could not parse pair")
	}

	firstEnd, err := strconv.Atoi(firstPair[1])

	if err != nil {
		return [2]int{}, [2]int{}, fmt.Errorf("could not parse pair")
	}

	secondStart, err := strconv.Atoi(secondPair[0])

	if err != nil {
		return [2]int{}, [2]int{}, fmt.Errorf("could not parse pair")
	}

	secondEnd, err := strconv.Atoi(secondPair[1])

	if err != nil {
		return [2]int{}, [2]int{}, fmt.Errorf("could not parse pair")
	}

	return [2]int{firstStart, firstEnd}, [2]int{secondStart, secondEnd}, nil

}
