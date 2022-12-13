package day7

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2022/internal/common"
)

type code int

const FS_SIZE = 70000000
const UPDATE_SIZE = 30000000
const DIR_THRESHOLD = 100000

const (
	CD code = iota
	LS
	DIR
	FILE
)

type fs struct {
	root *dir
}

type file struct {
	sizeBytes int
	name      string
}

type dir struct {
	name  string
	files []*file
	dirs  []*dir
}

// Solve returns the solutions for day 7
func Solve(s []string, p common.Part) int {

	fs := getNewFs()

	proc := getProcessor(fs)

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		err := proc(line)

		if err != nil {
			log.Printf("could not process line")
			return 0
		}
	}

	m := make(map[string]int)

	fs.root.getSizeMap("", m)

	if p == common.Part1 {
		return getDirSizeThreshold(DIR_THRESHOLD, m)
	}

	if p == common.Part2 {
		s, err := getDirCandidateSize(FS_SIZE, UPDATE_SIZE, m)

		if err != nil {
			log.Printf("could not get dir candidate size")
			return 0
		}

		return s
	}

	return 0

}

func getNewFs() *fs {
	return &fs{
		root: &dir{
			name: "/",
		},
	}
}

func getProcessor(fs *fs) func(line string) error {

	wd := fs.root

	return func(line string) error {

		code, o, err := getOpCode(line)

		if err != nil {
			return err
		}

		switch code {
		case FILE:
			f := o.(*file)
			wd.files = append(wd.files, f)
		case DIR:
			d := o.(*dir)
			dd := &dir{
				name: "..",
				dirs: []*dir{wd},
			}

			d.dirs = append(d.dirs, dd)
			wd.dirs = append(wd.dirs, d)

		case CD:
			d := o.(string)

			if d == "/" {
				wd = fs.root
				break
			}

			for _, dir := range wd.dirs {
				if dir.name == d {
					wd = dir
				}
			}

			if d == ".." {
				wd = wd.dirs[0]
			}

		case LS:
			// ignore ls commands
		}

		return nil
	}
}

func getOpCode(line string) (code, interface{}, error) {

	// CD command
	if strings.HasPrefix(line, "$ cd ") {

		s := strings.Split(line, " ")

		if len(s) != 3 {
			return 0, nil, fmt.Errorf("invalid cd command: %s", line)
		}

		return CD, s[2], nil
	}

	// LS Command
	if strings.HasPrefix(line, "$ ls") {
		return LS, nil, nil
	}

	// DIR Command
	if strings.HasPrefix(line, "dir ") {

		s := strings.Split(line, " ")

		if len(s) != 2 {
			return 0, nil, fmt.Errorf("invalid dir command: %s", line)
		}

		d := &dir{
			name: s[1],
		}

		return DIR, d, nil
	}

	// FILE Command
	s := strings.Split(line, " ")

	if len(s) != 2 {
		return 0, nil, fmt.Errorf("invalid file: %s", line)
	}

	size, err := strconv.Atoi(s[0])

	if err == nil {
		f := &file{
			sizeBytes: size,
			name:      s[1],
		}

		return FILE, f, nil
	}

	return 0, nil, fmt.Errorf("unknown op code: %s", line)
}

func (d *dir) getSize() int {
	size := 0

	if d.name == ".." {
		return size
	}

	for _, f := range d.files {
		size += f.sizeBytes
	}

	for _, dd := range d.dirs {
		size += dd.getSize()
	}

	return size
}

// getSizeMap returns a map of directory name to size
func (d *dir) getSizeMap(prefix string, m map[string]int) map[string]int {

	if d.name == ".." {
		return m
	}

	if len(d.dirs) == 0 {
		return m
	}

	p := fmt.Sprintf("%s/%s", prefix, d.name)

	size := d.getSize()

	m[p] = size

	for _, dd := range d.dirs {

		m = dd.getSizeMap(p, m)

	}

	return m

}

// getDirSizeThreshold returns the sum of all directory
// smaller than the threshold t
func getDirSizeThreshold(t int, m map[string]int) int {

	var sum int

	for _, size := range m {
		if size < t {
			sum += size
		}
	}

	return sum

}

// getDirCandidateSize gets the size of the fs fsSize and the size
// of the update: updateSize. It will check all directories in the
// m map to find what is the smallest directory that, if deleted,
// will make enough space for the update
func getDirCandidateSize(fsSize, updateSize int, m map[string]int) (int, error) {

	var a []int

	for _, size := range m {
		a = append(a, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	used := a[0]

	var b int

	for i, size := range a {

		if (fsSize-used)+size > updateSize {
			b = i
		}
	}

	if b == 0 {
		return 0, fmt.Errorf("no candidate found")
	}

	return a[b], nil
}
