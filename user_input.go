package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var INVALID_INPUT_MSG = fmt.Sprintf("Invalid Input, please ensure you enter a single integer value between 1 and %d\n", NUM_SQUARES)

func SolveInteractive() {
	s := GenerateSudoku()
	reader := bufio.NewReader(os.Stdin)
	for !s.isSolved() {
		performTurn(s, reader)
	}
	fmt.Println("Congratulations, you successfully solved the puzzle")
}

func performTurn(s *Sudoku, reader *bufio.Reader) {
	fmt.Printf("Current Puzzle State %s\n", s)
	col := getUserInput(reader, "Please enter the column number you would like to input: ")
	row := getUserInput(reader, "Please enter the row number you would like to input: ")
	if s.grid[row-1][col-1] != 0 {
		fmt.Printf("Selected Square already has value %d, please try again\n", s.grid[row-1][col-1])
		return
	}

	inputVal := getUserInput(reader, "Please enter the value you would like to enter at that location: ")
	possibleValues := s.getPotentialValues(col-1, row-1)
	if !slices.Contains(possibleValues, inputVal) {
		fmt.Printf("Value %d is not valid for column %d, row %d\n", inputVal, col, row)
		return
	}

	s.grid[row-1][col-1] = inputVal
}

func getUserInput(reader *bufio.Reader, msg string) int {
	var userInput string
	var err error
	var intVal int

	for {
		fmt.Println(msg)
		userInput, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf(INVALID_INPUT_MSG)
			continue
		}
		intVal, err = strconv.Atoi(strings.TrimSpace(userInput))
		if err != nil || intVal <= 0 || intVal > NUM_SQUARES {
			fmt.Printf(INVALID_INPUT_MSG)
			continue
		}

		return intVal
	}
}
