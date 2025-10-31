package main

import (
	"fmt"
	"strings"
)

const GRID_SIZE = 3

type Sudoku struct {
	grid [][]int
}

func newSudoku() *Sudoku {
	grid := make([][]int, GRID_SIZE*GRID_SIZE)
	for i := range grid {
		grid[i] = make([]int, GRID_SIZE*GRID_SIZE)
	}
	grid[2][2] = 4
	grid[4][4] = 9
	return &Sudoku{grid}
}

func (s *Sudoku) String() string {
	var sb strings.Builder
	sb.WriteRune('\n')
	for i := range GRID_SIZE {
		sb.WriteString("+-------+-------+-------+\n")
		for j := range GRID_SIZE {
			writeElementLine(&sb, s.grid[(i*GRID_SIZE)+j])
		}
	}
	sb.WriteString("+-------+-------+-------+\n")
	return sb.String()
}

func writeElementLine(sb *strings.Builder, line []int) {
	startPos := 0
	for range GRID_SIZE {
		sb.WriteString("| ")
		for i := range GRID_SIZE {
			el := line[startPos+i]
			if el == 0 {
				sb.WriteString(". ")
			} else {
				fmt.Fprintf(sb, "%d ", el)
			}
		}
		startPos += GRID_SIZE
	}
	sb.WriteString("|\n")
}
