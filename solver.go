package main

import (
	"fmt"
	"maps"
	"slices"
	"sync"
)

const MAX_VALUE = 9

func SolveSudokus(args *CmdArgs) {
	if args.parallel {
		solveParallel(args.numSudokus)
	} else {
		solveSequential(args.numSudokus)
	}
}

func solveParallel(numSudokus int) {
	var waitGroup sync.WaitGroup
	for i := range numSudokus {
		waitGroup.Go(func() { solveSingle(i) })
	}
	waitGroup.Wait()
}

func solveSequential(numSudokus int) {
	for i := range numSudokus {
		solveSingle(i)
	}
}

func solveSingle(i int) {
	s := GenerateSudoku()
	fmt.Printf("Puzzle %d:\n%s\n", i+1, s)
	s.solve()
	fmt.Printf("Solution: %s\n", s)
}

func (s *Sudoku) isSolved() bool {
	for y := range len(s.grid) {
		for x := range len(s.grid[y]) {
			if s.grid[y][x] == 0 {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) solve() error {
	startX, startY, foundUnfilled := s.getNextUnfilledSquare(-1, 0)
	if !foundUnfilled {
		return nil
	}
	if !s.solveFromSquare(startX, startY) {
		return fmt.Errorf("[ERROR]: Unable to find any solutions for given Sudoku")
	}
	return nil
}

func (s *Sudoku) getNumSolutions() int {
	startX, startY, foundUnfilled := s.getNextUnfilledSquare(-1, 0)
	if !foundUnfilled {
		return 1
	}
	return s.getNumSolutionsFromSquare(startX, startY)
}

func (s *Sudoku) getNumSolutionsFromSquare(x, y int) int {
	potentialValues := s.getPotentialValues(x, y)

	numSolutions := 0
	if len(potentialValues) == 0 {
		return numSolutions
	}

	newX, newY, foundUnfilled := s.getNextUnfilledSquare(x, y)
	if !foundUnfilled {
		numSolutions++
	}

	for i := range potentialValues {
		s.grid[y][x] = potentialValues[i]

		if foundUnfilled {
			numSolutions += s.getNumSolutionsFromSquare(newX, newY)
		}

	}

	s.grid[y][x] = 0

	return numSolutions
}

func (s *Sudoku) solveFromSquare(x, y int) bool {
	potentialValues := s.getPotentialValues(x, y)

	if len(potentialValues) == 0 {
		return false
	}

	newX, newY, foundUnfilled := s.getNextUnfilledSquare(x, y)
	if !foundUnfilled {
		s.grid[y][x] = potentialValues[0]
		return true
	}

	for i := range potentialValues {
		s.grid[y][x] = potentialValues[i]

		if s.solveFromSquare(newX, newY) {
			return true
		}

	}

	s.grid[y][x] = 0
	return false
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
