package main

import (
	"log"

	"github.com/wincus/adventofcode2022/internal/common"
	"github.com/wincus/adventofcode2022/internal/day1"
)

func main() {

	d, err := common.GetData(1)

	if err != nil {
		log.Panicf("no data, no game: %v", err)
	}

	for _, p := range []common.Part{common.Part1, common.Part2} {
		log.Printf("Solution for Part %v: %v", p, day1.Solve(d, p))
	}
}
