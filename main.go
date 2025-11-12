package main

import "fmt"

func main() {
	args, err := parseCmdArgs()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	err = SolveSudokus(args.numSudokus)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
