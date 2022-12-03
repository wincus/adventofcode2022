package main

import (
	"log"

	"github.com/wincus/adventofcode2022/internal/common"
	"github.com/wincus/adventofcode2022/internal/day3"
)

func main() {

	d, err := common.GetData(3)

	if err != nil {
		log.Panicf("no data, no game ... sorry!")
	}

	for _, p := range []common.Part{common.Part1, common.Part2} {
		log.Printf("Solution for Part %v: %v", p, day3.Solve(d, p))
	}
}
