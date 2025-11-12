package main

import (
	"flag"
	"fmt"
)

const MODE_USAGE = "Mode - 0 to generate, 1 to output puzzles with solutions, 2 to interactively solve a generated puzzle"
const NUM_SUDOKUS_USAGE = "Number of sudokus to perform selected operation on (maximum 100)"
const PARALLEL_USAGE = "Whether to perform operations in parallel - does not apply to interactive mode"
const MAX_SUDOKUS = 1000

const NUM_MODES = 3

type Mode int

const (
	ModeGenerate Mode = iota
	ModeSolve
	ModeInteractive
)

type CmdArgs struct {
	mode       Mode
	numSudokus int
	parallel   bool
}

func parseCmdArgs() (CmdArgs, error) {
	modePtr := flag.Int("m", 0, MODE_USAGE)
	numSudokusPtr := flag.Int("n", 1, NUM_SUDOKUS_USAGE)
	parallelPtr := flag.Bool("p", false, PARALLEL_USAGE)

	flag.Parse()

	if *modePtr < 0 || *modePtr >= NUM_MODES {
		return CmdArgs{}, fmt.Errorf("[ERROR]: Invalid mode argument provided\nMode Usage: %s", MODE_USAGE)
	}

	if *numSudokusPtr <= 0 || *numSudokusPtr > MAX_SUDOKUS {
		return CmdArgs{}, fmt.Errorf("[ERROR]: Invalid number of sudokus provided, value must be in range 1-100")
	}

	return CmdArgs{
		mode:       Mode(*modePtr),
		numSudokus: *numSudokusPtr,
		parallel:   *parallelPtr,
	}, nil
}
