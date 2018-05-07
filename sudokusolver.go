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
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Sudoku [9][9]int

func (s Sudoku) isValid(col, row, val int) (valid bool) {
	defer func() {
		fmt.Printf("(%d,%d): %d -- %v\n", col, row, val, valid)
	}()

	for i := 0; i < 9; i++ {
		if s[i][row] == val || s[col][i] == val {
			return
		}
	}
	topX, topY := col/3*3, row/3*3
	for i := topX; i < topX+3; i++ {
		for j := topY; j < topY+3; j++ {
			if s[i][j] == val {
				return
			}
		}
	}
	valid = true
	return
}

func (s Sudoku) solve(col, row int) (Sudoku, error) {
	if row >= 9 {
		return s, nil
	}
	if col >= 9 {
		return s.solve(0, row+1)
	}
	if s[row][col] != 0 {
		return s.solve(col, row+1)
	}

	for v := 1; v <= 9; v++ {
		if !s.isValid(col, row, v) {
			continue
		}
		s[row][col] = v
		new, err := s.solve(col, row)
		if err == nil {
			return new, err
		}
	}
	return s, errors.New("no solution")
}

func (s Sudoku) Solve() (Sudoku, error) {
	return s.solve(0, 0)
}

func readPuzzle(r io.Reader) (*Sudoku, error) {
	scanner := bufio.NewScanner(r)
	var col, row int
	var puzzle Sudoku
	for scanner.Scan() {
		line := scanner.Text()
		ws := bufio.NewScanner(strings.NewReader(line))
		ws.Split(bufio.ScanWords)
		for ws.Scan() {
			if col >= 9 {
				return nil, fmt.Errorf("Line '%s' too long", line)
			}
			d := ws.Text()
			if len(d) != 1 {
				return nil, fmt.Errorf("Entry '%s' too long", d)
			}
			if d != "_" {
				val, err := strconv.Atoi(d)
				if err != nil {
					return nil, err
				}
				puzzle[col][row] = val
			}

			col++
		}
		if col == 0 {
			// empty line
			continue
		}
		if err := ws.Err(); err != nil {
			return nil, err
		}
		row++
		col = 0
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &puzzle, nil
}

func main() {
	s, err := readPuzzle(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading puzzle: %v\n", err)
		os.Exit(-1)
	}

	solution, err := s.Solve()
	for _, line := range solution {
		for _, e := range line {
			if e == 0 {
				fmt.Print(" _")
			} else {
				fmt.Printf(" %d", e)
			}
		}
		fmt.Println()
	}
}
