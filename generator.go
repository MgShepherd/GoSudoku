package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const GRID_SIZE = 3
const MIN_CLUES = 25
const NUM_SQUARES = GRID_SIZE*GRID_SIZE

type Sudoku struct {
	grid [][]int
}

func newSudoku() *Sudoku {
	grid := make([][]int, NUM_SQUARES)
	for i := range grid {
		grid[i] = make([]int, NUM_SQUARES)
	}


	return &Sudoku{grid}
}

func (s *Sudoku) generateRandomGrid() error {
	numClues := 0
	for numClues < MIN_CLUES {
		x, y := rand.IntN(NUM_SQUARES), rand.IntN(NUM_SQUARES)
		if s.grid[y][x] != 0 {
			continue
		}
		potentialValues := s.getPotentialValues(x, y)

		if len(potentialValues) == 0 {
			return fmt.Errorf("Generated sudoku has no valid solutions")
		}

		s.grid[y][x] = potentialValues[rand.IntN(len(potentialValues))]
		numClues += 1
	}

	return nil
}

func generateFixedGrid() [][]int {
	return [][]int{
		{0, 0, 0, 8, 3, 0, 0, 5, 7},
		{0, 0, 8, 5, 0, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 2, 0, 8, 0},
		{8, 0, 2, 3, 9, 0, 7, 0, 0},
		{6, 0, 0, 1, 0, 0, 0, 3, 2},
		{0, 5, 7, 2, 0, 4, 0, 9, 0},
		{0, 6, 0, 4, 1, 0, 3, 7, 0},
		{0, 7, 3, 9, 0, 8, 0, 6, 0},
		{0, 0, 0, 7, 6, 0, 4, 0, 0},
	}
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
