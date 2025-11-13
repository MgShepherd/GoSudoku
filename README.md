# GoSudoku

A Sudoku Generator and Solver written in Golang

## Building the Project

This project can be run the same way as any normal Golang project (assuming you have Go installed on your system).

To build and run, you can use the command:

```
go run .
```

There are also additional optional command line arguments which can be provided:
- `-p` - Provide to enable parallel operations
- `-m` - The mode to use. 0 for Generating Sudokus, 1 for Solving Sudokus and 2 to interactively solve a generate puzzle in the terminal
- `-n` - The number of Sudokus to generate/solve - This does not apply to interactive mode
- `-f` - The file to output results to - defaults to stdout
