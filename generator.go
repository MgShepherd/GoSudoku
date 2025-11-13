package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"sync"
)

const GRID_SIZE = 3
const NUM_SQUARES = GRID_SIZE * GRID_SIZE

type Sudoku struct {
	grid [][]int
}

func GenerateSudokus(args *CmdArgs, f *os.File) {
	if args.parallel {
		generateParallel(args.numSudokus, f)
	} else {
		generateSequential(args.numSudokus, f)
	}
}

func generateParallel(numSudokus int, f *os.File) {
	var waitGroup sync.WaitGroup
	for range numSudokus {
		waitGroup.Go(func() { s := GenerateSudoku(); f.WriteString(s.String()) })
	}
	waitGroup.Wait()
}

func generateSequential(numSudokus int, f *os.File) {
	for range numSudokus {
		s := GenerateSudoku()
		f.WriteString(s.String())
	}
}

func GenerateSudoku() *Sudoku {
	grid := make([][]int, NUM_SQUARES)
	for i := range grid {
		grid[i] = make([]int, NUM_SQUARES)
	}
	s := Sudoku{grid}

	// All diagonal boxes are independant so generate these first
	s.generateBox(0, 0)
	s.generateBox(1, 1)
	s.generateBox(2, 2)

	s.generateConstrained(GRID_SIZE, 0, (NUM_SQUARES*NUM_SQUARES)-(GRID_SIZE*NUM_SQUARES))

	s.removeClues()
	return &s
}

func (s *Sudoku) generateBox(boxX, boxY int) {
	remainingOptions := sliceFromRange(1, 9)
	numRemaining := len(remainingOptions)

	for y := boxY * GRID_SIZE; y < (boxY*GRID_SIZE)+GRID_SIZE; y++ {
		for x := boxX * GRID_SIZE; x < (boxX*GRID_SIZE)+GRID_SIZE; x++ {
			randIdx := rand.IntN(numRemaining)
			s.grid[y][x] = remainingOptions[randIdx]
			numRemaining--
			remainingOptions[randIdx] = remainingOptions[numRemaining]
		}
	}
}

func (s *Sudoku) generateConstrained(x, y, remaining int) bool {
	if s.grid[y][x] != 0 {
		nextX, nextY := getNextPos(x, y)
		return s.generateConstrained(nextX, nextY, remaining)
	}

	potentialValues := s.getPotentialValues(x, y)

	if len(potentialValues) == 0 {
		return false
	}

	if remaining <= 1 {
		s.grid[y][x] = potentialValues[0]
		return true
	}

	success := false
	for _, el := range potentialValues {
		s.grid[y][x] = el
		nextX, nextY := getNextPos(x, y)
		success = s.generateConstrained(nextX, nextY, remaining-1)
		if success {
			break
		}
	}

	if !success {
		s.grid[y][x] = 0
	}
	return success
}

func (s *Sudoku) removeClues() {
	x, y := rand.IntN(NUM_SQUARES), rand.IntN(NUM_SQUARES)

	for s.grid[y][x] == 0 {
		x, y = rand.IntN(NUM_SQUARES), rand.IntN(NUM_SQUARES)
	}

	initialVal := s.grid[y][x]
	s.grid[y][x] = 0

	if s.getNumSolutions() > 1 {
		s.grid[y][x] = initialVal
		return
	}
	s.removeClues()
}

func getNextPos(currentX, currentY int) (int, int) {
	if currentX == NUM_SQUARES-1 {
		return 0, currentY + 1
	}
	return currentX + 1, currentY
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
