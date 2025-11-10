package main

import (
	"flag"
	"fmt"
)

const MODE_USAGE = "Mode - 0 to generate, 1 to solve"
const NUM_SUDOKUS_USAGE = "Number of sudokus to perform selected operation on (maximum 100)"

type Mode int

const (
	ModeGenerate Mode = iota
	ModeSolve
)

type CmdArgs struct {
	mode       Mode
	numSudokus int
}

func parseCmdArgs() (CmdArgs, error) {
	modePtr := flag.Int("m", 0, MODE_USAGE)
	numSudokusPtr := flag.Int("n", 1, NUM_SUDOKUS_USAGE)

	flag.Parse()

	if *modePtr < 0 || *modePtr > 1 {
		return CmdArgs{}, fmt.Errorf("[ERROR]: Invalid mode argument provided\nMode Usage: %s", MODE_USAGE)
	}

	if *numSudokusPtr <= 0 || *numSudokusPtr > 100 {
		return CmdArgs{}, fmt.Errorf("[ERROR]: Invalid number of sudokus provided, value must be in range 1-100")
	}

	return CmdArgs{
		mode:       Mode(*modePtr),
		numSudokus: *numSudokusPtr,
	}, nil
}
