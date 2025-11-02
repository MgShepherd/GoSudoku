package main

import (
	"fmt"
	"maps"
	"slices"
)

const MAX_VALUE = 9

func (s *Sudoku) Solve() {
	fmt.Println("Sudoku is now solved")
	s.fillSquare()
}

func (s *Sudoku) fillSquare() error {
	potentialValues := s.getPotentialValues(4, 2)
	fmt.Println(potentialValues)

	return nil
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
