package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	args, err := parseCmdArgs()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	switch args.mode {
	case ModeGenerate:
		fallthrough
	case ModeSolve:
		SolveSudokus(&args)
	case ModeInteractive:
		SolveInteractive()
	}

	duration := time.Since(startTime)
	fmt.Printf("Program ran in %d microseconds (%d milliseconds)\n", duration.Microseconds(), duration.Milliseconds())
}
