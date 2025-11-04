package main

import "fmt"

func main() {
	sudoku := newSudoku()
	fmt.Println("Before: ", sudoku)
	sudoku.generateRandomGrid()
	fmt.Println("After Generation: ", sudoku)
	numSolutions := sudoku.Solve()
	fmt.Printf("Number of solutions for puzzle: %d\n", numSolutions)
}
