package main

import (
	"strings"
	"testing"
)

var testInput string = `
1 _ 3 _ _ 6 _ 8 _
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
_ 4 _ _ 7 8 _ 1 _
6 _ 8 _ _ 2 _ 4 _
_ 1 2 _ 4 5 _ 7 8
`

func TestReadPuzzle(t *testing.T) {
	s, err := readPuzzle(strings.NewReader(testInput))
	if err != nil {
		t.Fatalf("error reading puzzle from string: %v", err)
	}
	expectAt := func(col, row, val int) {
		if s[col][row] != val {
			t.Errorf("pos (%d,%d): expected %d, got %d", col, row, val, s[col][row])
		}
	}
	expectAt(0, 0, 1)
	expectAt(1, 0, 0)
	expectAt(7, 0, 8)
	expectAt(8, 8, 8)

}
