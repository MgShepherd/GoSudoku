package main

import (
	"maps"
	"slices"
)

const MAX_VALUE = 9

func (s *Sudoku) Solve() bool {
	startX, startY, foundUnfilled := s.getNextUnfilledSquare(-1, 0)
	if !foundUnfilled {
		return true
	}
	return s.solveFromSquare(startX, startY)
}

func (s *Sudoku) solveFromSquare(x, y int) bool {
	potentialValues := s.getPotentialValues(x, y)

	if len(potentialValues) == 0 {
		return false
	}

	newX, newY, foundUnfilled := s.getNextUnfilledSquare(x, y)
	solved := !foundUnfilled

	for i := range potentialValues {
		s.grid[y][x] = potentialValues[i]

		if foundUnfilled {
			solved = s.solveFromSquare(newX, newY)
			if solved {
				return true
			}
		}

	}

	if !solved {
		s.grid[y][x] = 0
	}
	
	return solved
}

func (s *Sudoku) getPotentialValues(x, y int) []int {
	potentialValues := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
	for i := 0; i < len(s.grid); i++ {
		delete(potentialValues, s.grid[i][x])
		delete(potentialValues, s.grid[y][i])
	}
	boxRow, boxCol := (y/GRID_SIZE)*GRID_SIZE, (x/GRID_SIZE)*GRID_SIZE
	for y := boxRow; y < boxRow+GRID_SIZE; y++ {
		for x := boxCol; x < boxCol+GRID_SIZE; x++ {
			delete(potentialValues, s.grid[y][x])
		}
	}
	return slices.Collect(maps.Keys(potentialValues))
}

func (s *Sudoku) getNextUnfilledSquare(startX, startY int) (int, int, bool) {
	for y := startY; y < len(s.grid); y++ {
		for x := startX + 1; x < len(s.grid[y]); x++ {
			if s.grid[y][x] == 0 {
				return x, y, true
			}
		}
		startX = -1
	}
	return -1, -1, false
}
