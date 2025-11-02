package main

import "fmt"

func main() {
	sudoku := newSudoku()
	fmt.Println(sudoku)
	sudoku.Solve()
}
