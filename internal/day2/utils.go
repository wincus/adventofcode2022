package day2

import (
	"fmt"
	"strings"

	"github.com/wincus/adventofcode2022/internal/common"
)

type shape int

type result int

const (
	draw result = iota
	win
	lose
)

const (
	rock shape = iota
	paper
	scissors
)

// Solve returns the solutions for day 2
func Solve(s []string, p common.Part) int {

	var points int
	var shape1, shape2 shape
	var err error

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		shapes := strings.Split(line, " ")

		if len(shapes) != 2 {
			fmt.Printf("invalid input: %v", line)
			return 0
		}

		shape1, err = getShape(shapes[0])

		if err != nil {
			fmt.Printf("invalid input: %v", line)
			return 0
		}

		if p == common.Part1 {

			shape2, err = getShape(shapes[1])

			if err != nil {
				fmt.Printf("invalid input: %v", line)
				return 0
			}
		} else if p == common.Part2 {
			result, err := getResult(shapes[1])

			if err != nil {
				fmt.Printf("invalid input: %v", line)
				return 0
			}

			shape2, err = getShapeForResult(shape1, result)

			if err != nil {
				fmt.Printf("invalid input: %v", line)
				return 0
			}
		}

		turn, err := getPoints(shape1, shape2)

		if err != nil {
			fmt.Printf("invalid input: %v", line)
			return 0
		}

		points += turn
	}

	return points
}

// getShapeForResult gets a shape and a result. It returns the shape
// that would satisfy that result
func getShapeForResult(s shape, r result) (shape, error) {

	if r == win {
		switch s {
		case rock:
			return paper, nil
		case paper:
			return scissors, nil
		case scissors:
			return rock, nil
		default:
			return 0, fmt.Errorf("invalid shape: %v", s)
		}
	}

	if r == lose {
		switch s {
		case rock:
			return scissors, nil
		case paper:
			return rock, nil
		case scissors:
			return paper, nil
		default:
			return 0, fmt.Errorf("invalid shape: %v", s)
		}
	}

	// draw
	return s, nil

}

// getPoints returns the points for a round of
// rick, paper & scissors where a is the first player
// and b is the second player ( you )
func getPoints(a, b shape) (int, error) {

	var points int

	switch b {
	case rock:
		points = 1
	case paper:
		points = 2
	case scissors:
		points = 3
	default:
		return 0, fmt.Errorf("invalid shape: %v", b)
	}

	// draw equals to 3 points
	if a == b {
		points += 3
	}

	// a win on my side equals to 6 points
	if a == rock && b == paper || a == paper && b == scissors || a == scissors && b == rock {
		points += 6
	}

	return points, nil
}

func getShape(s string) (shape, error) {
	switch s {
	case "A", "X":
		return rock, nil
	case "B", "Y":
		return paper, nil
	case "C", "Z":
		return scissors, nil
	default:
		return 0, fmt.Errorf("invalid shape: %v", s)
	}
}

func getResult(s string) (result, error) {
	switch s {
	case "X":
		return lose, nil
	case "Y":
		return draw, nil
	case "Z":
		return win, nil
	default:
		return 0, fmt.Errorf("invalid result: %v", s)
	}
}
