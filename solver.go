package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
)

const MAX_VALUE = 9

func (s *Sudoku) Solve() bool {
	startX, startY, foundUnfilled := s.getNextUnfilledSquare(0, 0)
	if !foundUnfilled {
		return true
	}
	solved, err := s.solveFromSquare(startX, startY)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Unable to solve Sudoku: %s\n", err)
		return false
	}
	return solved
}

func (s *Sudoku) solveFromSquare(x, y int) (bool, error) {
	potentialValues := s.getPotentialValues(x, y)

	if len(potentialValues) == 0 {
		return false, fmt.Errorf("No potential options at location X: %d, Y: %d", x, y)
	}

	s.grid[y][x] = potentialValues[0]
	newX, newY, foundUnfilled := s.getNextUnfilledSquare(x, y)
	if !foundUnfilled {
		return true, nil
	}
	solved, err := s.solveFromSquare(newX, newY)
	if err != nil {
		return false, err
	}
	return solved, nil
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
		for x := startX; x < len(s.grid[y]); x++ {
			if s.grid[y][x] == 0 {
				return x, y, true
			}
		}
		startX = 0
	}
	return -1, -1, false
}
