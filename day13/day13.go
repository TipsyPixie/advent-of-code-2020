package day13

import (
	"fmt"
	aoc "github.com/TipsyPixie/advent-of-code-2020"
	"strconv"
	"strings"
)

func solvePart1(inputPath string) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	line, _, err := input.ReadLine()
	if err != nil {
		return 0, err
	}
	earliestDeparture, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}

	line, _, err = input.ReadLine()
	if err != nil {
		return 0, err
	}
	shortestWaitTime, matchedBusId := -1, -1
	for _, busID := range strings.Split(line, ",") {
		if busID != "x" {
			IDNumber, err := strconv.Atoi(busID)
			if err != nil {
				return 0, err
			}
			if shortestWaitTime == -1 || shortestWaitTime > IDNumber-earliestDeparture%IDNumber {
				shortestWaitTime = IDNumber - earliestDeparture%IDNumber
				matchedBusId = IDNumber
			}
		}
	}

	fmt.Println(shortestWaitTime, matchedBusId)
	return shortestWaitTime * matchedBusId, nil
}
