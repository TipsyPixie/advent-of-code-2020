package day5

import (
	"errors"
	aoc "github.com/TipsyPixie/advent-of-code-2020"
)

func getNumber(code string, minNumber int, maxNumber int) int {
	for _, sign := range []rune(code) {
		if sign == 'F' || sign == 'L' {
			maxNumber -= (maxNumber - minNumber + 1) / 2
		} else {
			minNumber += (maxNumber - minNumber + 1) / 2
		}
	}
	// assume that minNumber == maxNumber
	return minNumber
}

func getRow(rowCode string) int {
	return getNumber(rowCode, 0, 127)
}

func getColumn(columnCode string) int {
	return getNumber(columnCode, 0, 7)
}

func getSeatId(row int, column int) int {
	return row*8 + column
}

func solvePart1(inputPath string) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	highestSeatId := 0
	for seatCode, ok, err := input.ReadLine(); ok || err != nil; seatCode, ok, err = input.ReadLine() {
		if err != nil {
			return 0, err
		}
		row, column := getRow(seatCode[:7]), getColumn(seatCode[7:])
		if seatId := getSeatId(row, column); seatId > highestSeatId {
			highestSeatId = seatId
		}
	}
	return highestSeatId, nil
}

func solvePart2(inputPath string) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	seats := make([][]int, 128, 128)
	for rowNumber := range seats {
		seats[rowNumber] = make([]int, 8, 8)
	}

	seatIdsListed := map[int]bool{}
	for seatCode, ok, err := input.ReadLine(); ok || err != nil; seatCode, ok, err = input.ReadLine() {
		if err != nil {
			return 0, err
		}
		row, column := getRow(seatCode[:7]), getColumn(seatCode[7:])
		seatId := getSeatId(row, column)
		seats[row][column] = seatId
		seatIdsListed[seatId] = true
	}

	for rowNumber, row := range seats {
		for columnNumber, seatId := range row {
			if seatId == 0 {
				seatId = getSeatId(rowNumber, columnNumber)
				_, plusOneListed := seatIdsListed[seatId+1]
				_, minusOneListed := seatIdsListed[seatId-1]
				if plusOneListed && minusOneListed {
					return seatId, nil
				}
			}
		}
	}
	return 0, errors.New("seat not found")
}
