package main

import "fmt"

func main() {
	sudoku := newSudoku()
	sudoku.Generate()
	fmt.Println("After Generation: ", sudoku)
}
