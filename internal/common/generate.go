package common

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

const UTILS = `package day{{.Day}}

import (
	"github.com/wincus/adventofcode2022/internal/common"
)

// Solve returns the solutions for day {{.Day}}
func Solve(s []string, p common.Part) int {
	return 0
}
`

const TEST = `package day{{.Day}}

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
			input: []string{},
			p:     common.Part1,
			want:  0,
		},
		{
			input: []string{},
			p:     common.Part2,
			want:  0,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
`

const MAIN = `package main

import (
	"log"

	"github.com/wincus/adventofcode2022/internal/common"
	"github.com/wincus/adventofcode2022/internal/day{{.Day}}"
)

func main() {

	d, err := common.GetData({{.Day}})

	if err != nil {
		log.Panicf("no data, no game ... sorry!")
	}

	for _, p := range []common.Part{common.Part1, common.Part2} {
		log.Printf("Solution for Part %v: %v", p, day{{.Day}}.Solve(d, p))
	}
}
`

type data struct {
	Day string
}

func Generate(day string) error {

	if day == "" {
		panic("day number is required")
	}

	u, err := template.New("utils").Parse(UTILS)

	if err != nil {
		return err
	}

	t, err := template.New("test").Parse(TEST)

	if err != nil {
		return err
	}

	m, err := template.New("main").Parse(MAIN)

	if err != nil {
		return err
	}

	d := fmt.Sprintf("internal/day%s", day)

	_, err = os.Stat(d)

	if err == nil {
		log.Fatal("Folder already exists.")
	}

	err = os.MkdirAll(d, 0755)

	if err != nil {
		return err
	}

	e := fmt.Sprintf("solutions/day%s", day)

	err = os.MkdirAll(e, 0755)

	if err != nil {
		return err
	}

	fMain, _ := os.Create(fmt.Sprintf("%s/main.go", e))
	defer fMain.Close()

	err = m.Execute(fMain, data{day})

	if err != nil {
		return err
	}

	fUtils, _ := os.Create(fmt.Sprintf("%s/utils.go", d))
	defer fUtils.Close()

	err = u.Execute(fUtils, data{day})

	if err != nil {
		return err
	}

	fTest, _ := os.Create(fmt.Sprintf("%s/utils_test.go", d))
	defer fTest.Close()

	err = t.Execute(fTest, data{day})

	if err != nil {
		return err
	}

	return nil

}
