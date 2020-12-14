package day11

import (
	aoc "github.com/TipsyPixie/advent-of-code-2020"
)

type state string

const (
	EMPTY       = state("L")
	OCCUPIED    = state("#")
	NONEXISTENT = state(".")
)

type board struct {
	alternativeCounting bool
	states              [][]state
	occupationCount     int
}

func parseLine(signs string) []state {
	states := make([]state, 0, len(signs))
	for _, character := range signs {
		states = append(states, state(character))
	}
	return states
}

func (thisBoard *board) countAround(rowIndex int, columnIndex int) int {
	maxRowIndex, maxColumnIndex := len(thisBoard.states)-1, len(thisBoard.states[0])-1
	getCoords := func(r int, rChange int, c int, cChange int) (int, int) {
		return r + rChange, c + cChange
	}
	if thisBoard.alternativeCounting {
		getCoords = func(r int, rChange int, c int, cChange int) (int, int) {
			for {
				rNext, cNext := r+rChange, c+cChange
				if rNext < 0 || rNext > maxRowIndex || cNext < 0 || cNext > maxColumnIndex {
					break
				}
				r, c = rNext, cNext
				if thisBoard.states[r][c] != NONEXISTENT {
					break
				}
			}
			return r, c
		}
	}
	oneIfOccupied := func(r int, c int) int {
		if thisBoard.states[r][c] == OCCUPIED {
			return 1
		}
		return 0
	}
	count := 0
	if rowIndex > 0 {
		count += oneIfOccupied(getCoords(rowIndex, -1, columnIndex, 0))
	}
	if rowIndex < maxRowIndex {
		count += oneIfOccupied(getCoords(rowIndex, +1, columnIndex, 0))
	}
	if columnIndex > 0 {
		count += oneIfOccupied(getCoords(rowIndex, 0, columnIndex, -1))
	}
	if columnIndex < maxColumnIndex {
		count += oneIfOccupied(getCoords(rowIndex, 0, columnIndex, +1))
	}
	if rowIndex > 0 && columnIndex > 0 {
		count += oneIfOccupied(getCoords(rowIndex, -1, columnIndex, -1))
	}
	if rowIndex > 0 && columnIndex < maxColumnIndex {
		count += oneIfOccupied(getCoords(rowIndex, -1, columnIndex, +1))
	}
	if rowIndex < maxRowIndex && columnIndex > 0 {
		count += oneIfOccupied(getCoords(rowIndex, +1, columnIndex, -1))
	}
	if rowIndex < maxRowIndex && columnIndex < maxColumnIndex {
		count += oneIfOccupied(getCoords(rowIndex, +1, columnIndex, +1))
	}
	return count
}

func (thisBoard *board) getNextState(rowIndex int, columnIndex int) state {
	emptyThreshold := 4
	if thisBoard.alternativeCounting {
		emptyThreshold = 5
	}

	switch {
	case thisBoard.states[rowIndex][columnIndex] == EMPTY && thisBoard.countAround(rowIndex, columnIndex) == 0:
		return OCCUPIED
	case thisBoard.states[rowIndex][columnIndex] == OCCUPIED && thisBoard.countAround(rowIndex, columnIndex) >= emptyThreshold:
		return EMPTY
	default:
		return thisBoard.states[rowIndex][columnIndex]
	}
}

func (thisBoard *board) proceed() (changed bool) {
	nextStates := make([][]state, len(thisBoard.states), len(thisBoard.states))
	for rowIndex := range nextStates {
		nextStates[rowIndex] = make([]state, len(thisBoard.states[rowIndex]), len(thisBoard.states[rowIndex]))
		for columnIndex := range nextStates[rowIndex] {
			nextStates[rowIndex][columnIndex] = thisBoard.getNextState(rowIndex, columnIndex)
			if nextStates[rowIndex][columnIndex] != thisBoard.states[rowIndex][columnIndex] {
				changed = true
				if nextStates[rowIndex][columnIndex] == OCCUPIED {
					thisBoard.occupationCount++
				} else {
					thisBoard.occupationCount--
				}
			}
		}
	}

	thisBoard.states = nextStates
	return
}

func solve(inputPath string, alternativeCounting bool) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	states := make([][]state, 0, 64)
	for line, ok, err := input.ReadLine(); ok || err != nil; line, ok, err = input.ReadLine() {
		if err != nil {
			return 0, err
		}
		states = append(states, parseLine(line))
	}

	tempBoard := board{
		alternativeCounting: alternativeCounting,
		states:              states,
		occupationCount:     0,
	}
	for changed := tempBoard.proceed(); changed; changed = tempBoard.proceed() {
	}
	return tempBoard.occupationCount, nil
}

func solvePart1(inputPath string) (int, error) {
	return solve(inputPath, false)
}

func solvePart2(inputPath string) (int, error) {
	return solve(inputPath, true)
}
