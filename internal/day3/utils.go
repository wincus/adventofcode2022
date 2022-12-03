package day3

import (
	"fmt"
	"log"

	"github.com/wincus/adventofcode2022/internal/common"
)

const GROUPSIZE = 3

// Solve returns the solutions for day 3
func Solve(s []string, p common.Part) int {

	var sum int
	var group []string

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		if p == common.Part1 {
			r, err := getRepeatedItem(line[:len(line)/2], line[len(line)/2:])

			if err != nil {
				log.Printf("could not find repeated item for part %v", p)
				return 0
			}

			sum += getPriority(r)
		}

		if p == common.Part2 {

			group = append(group, line)

			if len(group) == GROUPSIZE {
				r, err := getRepeatedItem(group...)

				if err != nil {
					log.Printf("could not find repeated item for part %v", p)
					return 0
				}
				sum += getPriority(r)

				group = nil
			}

		}

	}

	return sum
}

func getRepeatedItem(s ...string) (rune, error) {

	var maps []map[rune]bool

	for _, line := range s {
		m := getMap(line)
		maps = append(maps, m)
	}

	keys := getMapKeys(maps...)

	for _, k := range keys {
		if isItemInAllMaps(k, maps...) {
			return k, nil
		}
	}

	return 0, fmt.Errorf("could not find the repetead item")

}

func isItemInAllMaps(k rune, maps ...map[rune]bool) bool {

	for _, m := range maps {
		if !m[k] {
			return false
		}
	}

	return true
}

func getMap(s string) map[rune]bool {

	m := make(map[rune]bool)

	for _, r := range s {
		m[r] = true
	}

	return m

}

func getMapKeys(m ...map[rune]bool) []rune {

	var r []rune

	for _, v := range m {
		for k := range v {
			r = append(r, k)
		}
	}

	return r
}

func getPriority(r rune) int {

	if r >= 'a' && r <= 'z' {
		return int(r - '`')
	}

	if r >= 'A' && r <= 'Z' {
		return int(r - '&')
	}

	return 0

}
