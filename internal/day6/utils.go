package day6

import (
	"log"

	"github.com/wincus/adventofcode2022/internal/common"
)

const (
	START_OF_PACKET_LENGHT  = 4
	START_OF_MESSAGE_LENGTH = 14
)

// Solve returns the solutions for day 6
func Solve(s []string, p common.Part) int {

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		// look for start of packet markers
		if p == common.Part1 {

			for i := 0; i < len(line); i++ {

				if i < START_OF_PACKET_LENGHT {
					continue
				}

				if check(line[i-START_OF_PACKET_LENGHT : i]) {
					return i
				}
			}
		}

		// look for start of message markers
		if p == common.Part2 {

			for i := 0; i < len(line); i++ {

				if i < START_OF_MESSAGE_LENGTH {
					continue
				}

				if check(line[i-START_OF_MESSAGE_LENGTH : i]) {
					return i
				}
			}
		}

	}

	log.Printf("no solution found")

	return 0
}

func check(s string) bool {

	var m = make(map[rune]int)

	for i := 0; i < len(s); i++ {
		m[rune(s[i])]++
	}

	for _, v := range m {
		if v > 1 {
			return false
		}
	}

	return true

}
