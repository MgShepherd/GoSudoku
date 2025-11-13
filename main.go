package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	args, err := parseCmdArgs()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	var f *os.File
	f, err = os.Create(args.fileName)
	if err != nil {
		fmt.Printf("[ERROR] Could not open file: %s\n", err)
		return
	}
	defer f.Close()

	switch args.mode {
	case ModeGenerate:
		GenerateSudokus(&args, f)
	case ModeSolve:
		SolveSudokus(&args, f)
	case ModeInteractive:
		SolveInteractive()
	}

	duration := time.Since(startTime)
	fmt.Printf("Program ran in %d microseconds (%d milliseconds)\n", duration.Microseconds(), duration.Milliseconds())
}
