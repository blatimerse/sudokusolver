/*
The Goal of the challenge

The goal of this challenge is to implement a Sudoku solver.

Requirements of the challenge

Your program should read a puzzle of this form from standard input:

1 _ 3 _ _ 6 _ 8 _
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
_ 4 _ _ 7 8 _ 1 _
6 _ 8 _ _ 2 _ 4 _
_ 1 2 _ 4 5 _ 7 8
And it should write the solution to standard output:

1 2 3 4 5 6 7 8 9
4 5 6 7 8 9 1 2 3
7 8 9 1 2 3 4 5 6
2 3 4 5 6 7 8 9 1
5 6 7 8 9 1 2 3 4
8 9 1 2 3 4 5 6 7
3 4 5 6 7 8 9 1 2
6 7 8 9 1 2 3 4 5
9 1 2 3 4 5 6 7 8
It should reject malformed or invalid inputs and recognize and report puzzles that cannot be solved.

(Incidentally, the puzzle above makes a nice test case, because the solution is easy to validate by sight.)

Bonus features:

Print a rating of the puzzle's difficulty (easy, medium, hard). This rating should roughly coincide with the ratings of shown
by sites like Web Sudoku.

Implement a puzzle generator that produces a puzzle of the given difficulty rating.
Maximize the efficiency of your program. (Write a benchmark.)
Write test cases, and use the cover tool to make sure your tests are thorough. (I found a bug in both my implementation and a test case when I checked the coverage.)
Use a non-obvious technique, like Knuth's "Dancing Links" or something of your own invention.


Hints

For an elegant and efficient representation of the puzzle, try using an array (not a slice).
Recursion can dramatically simplify your implementation.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Sudoku [9][9]int

func (s Sudoku) isValid(col, row, val int) bool {
	for i := 0; i < 9; i++ {
		if s[i][row] == val {
			return false
		}

		if s[col][i] == val {
			return false
		}
	}
	leftX, topY := col/3*3, row/3*3
	for i := leftX; i < leftX+3; i++ {
		for j := topY; j < topY+3; j++ {
			if s[i][j] == val {
				return false
			}
		}
	}
	return true
}

var iterations int

func (s *Sudoku) solve(col, row int) bool {
	defer func() { iterations++ }()
	nextpos := (col + row*9) + 1
	if nextpos > 9*9 {
		return true
	}

	nextcol, nextrow := nextpos%9, nextpos/9
	if s[col][row] != 0 {
		return s.solve(nextcol, nextrow)
	}

	for v := 1; v <= 9; v++ {
		if !s.isValid(col, row, v) {
			continue
		}
		s[col][row] = v
		if s.solve(nextcol, nextrow) {
			return true
		}
		s[col][row] = 0
	}
	return false
}

func (s Sudoku) Solve() (Sudoku, bool) {
	// make a copy to return as the solution
	solution := s
	return solution, solution.solve(0, 0)
}

func (s *Sudoku) Read(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	var col, row int
	for scanner.Scan() {
		line := scanner.Text()
		ws := bufio.NewScanner(strings.NewReader(line))
		ws.Split(bufio.ScanWords)
		for ws.Scan() {
			if col >= 9 {
				return fmt.Errorf("Line '%s' too long", line)
			}
			d := ws.Text()
			if len(d) != 1 {
				return fmt.Errorf("Entry '%s' too long", d)
			}
			if d != "_" {
				val, err := strconv.Atoi(d)
				if err != nil {
					return err
				}
				s[col][row] = val
			}

			col++
		}
		if col == 0 {
			// empty line
			continue
		}
		if err := ws.Err(); err != nil {
			return err
		}
		row++
		col = 0
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (s Sudoku) String() string {
	var str string
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			e := s[col][row]
			switch {
			case e == 0:
				str += "-"
			case e > 0 && e <= 9:
				str += string('0' + e)
			default:
				panic(fmt.Sprintf("byte %x is not a valid digit", e))
			}
			str += " "
		}
		str += "\n"
	}
	return str
}

func main() {
	for _, f := range os.Args[1:] {
		iterations = 0
		var puzzle Sudoku
		fh, err := os.Open(f)
		if err != nil {
			fmt.Printf("Error reading puzzle file %s: %v\n", f, err)
			os.Exit(-1)
		}
		err = puzzle.Read(fh)
		if err != nil {
			fmt.Printf("Error reading puzzle: %v\n", err)
			os.Exit(-1)
		}

		fmt.Printf("Original:\n%s\n", puzzle)
		solution, _ := puzzle.Solve()
		fmt.Println(f)
		fmt.Printf("Solution in %d iterations:\n%s\n", iterations, solution)
	}
}
