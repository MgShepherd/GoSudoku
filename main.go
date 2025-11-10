package main

import "fmt"

func main() {
	args, err := parseCmdArgs()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Successfully parsed arguments and got: %v\n", args)
	sudoku := newSudoku()
	sudoku.Generate()
	fmt.Println("After Generation: ", sudoku)
}
