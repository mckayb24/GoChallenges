package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const size = 9
const square = 3

type puzzle [size][size]uint8

type tester map[uint8]struct{}

func main() {
	p, err := readPuzzle()
	if err != nil {
		log.Fatalln("Read puzzle error", err)
	}
	valid := make(chan puzzle)
	go p.try(valid)
	log.Println(<-valid)
}

func readPuzzle() (puzzle, error) {
	r := bufio.NewReader(os.Stdin)
	var p puzzle
	for row := range p {
		for col := range p[row] {
			val, err := readEntry(r)
			if err != nil {
				return p, err
			}
			p[row][col] = val
		}
	}
	return p, nil
}

func readEntry(r *bufio.Reader) (uint8, error) {
	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		entry := strings.TrimSpace(string(b))
		switch entry {
		case "_":
			return 0, nil
		case "":
			continue
		default:
			val, err := strconv.ParseUint(entry, 10, 8)
			return uint8(val), err
		}
	}

}

func (p puzzle) String() string {
	var ps []string
	for _, row := range p {
		var rs []string
		for _, val := range row {
			rs = append(rs, strconv.FormatUint(uint64(val), 10))
		}
		ps = append(ps, strings.Join(rs, " "))
	}
	return strings.Join(ps, "\n")
}

func (p puzzle) try(valid chan puzzle) {
	for val := 1; val <= size; val++ {
		newP, full := p.add(uint8(val))
		if !newP.test() {
			continue
		}
		if full {
			valid <- newP
		}
		go newP.try(valid)
	}
}

// add creates a copy of puzzle with val added to next open spot
// returns true when puzzle is full
func (p puzzle) add(val uint8) (puzzle, bool) {
	for row := range p {
		for col, v := range p[row] {
			if v != 0 {
				continue
			}
			p[row][col] = val
			return p, false
		}
	}
	return p, true
}

func (p puzzle) test() bool {
	switch {
	case !p.testRows():
		return false
	case !p.testColumns():
		return false
	case !p.testBoxes():
		return false
	default:
		return true
	}
}

// testRows returns false if there is a problem in the rows
func (p puzzle) testRows() bool {
	for row := range p {
		t := tester{}
		for _, val := range p[row] {
			if !t.add(val) {
				return false
			}
		}
	}
	return true
}

func (p puzzle) testColumns() bool {
	for col := 0; col < size; col++ {
		t := tester{}
		for row := 0; row < size; row++ {
			if !t.add(p[row][col]) {
				return false
			}
		}
	}
	return true
}

func (p puzzle) testBoxes() bool {
	for box := 0; box < size; box++ {
		rowMod := (box / square) * square
		colMod := (box % square) * square
		t := tester{}
		for row := 0; row < square; row++ {
			for col := 0; col < square; col++ {
				if !t.add(p[rowMod+row][colMod+col]) {
					return false
				}
			}
		}
	}
	return true
}

// add returns false if val has already been added to tester
func (t tester) add(val uint8) bool {
	if val == 0 {
		return true
	}
	if _, ok := t[val]; ok {
		return false
	}
	t[val] = struct{}{}
	return true
}
