package main

import "fmt"

func main() {
	sudoku := newSudoku()
	fmt.Println("Before: ", sudoku)
	solved := sudoku.Solve()
	fmt.Println("After: ", sudoku)
	if solved {
		fmt.Println("Successfully Solved")
	} else {
		fmt.Println("Unable to solve Sudoku")
	}
}
