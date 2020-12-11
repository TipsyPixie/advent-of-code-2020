package day1

import (
	"fmt"
	aoc "github.com/TipsyPixie/advent-of-code-2020"
	"strconv"
)

func solve(input aoc.Input, n int) (int, error) {
	pairs := map[int]int{}
	for text, ok, err := input.ReadLine(); ok || err != nil; text, ok, err = input.ReadLine() {
		if err != nil {
			return 0, err
		}
		expense, err := strconv.Atoi(text)
		if err != nil {
			return 0, err
		}
		if pair, pairExists := pairs[expense]; pairExists {
			return expense * pair, nil
		}
		pairs[n-expense] = expense
	}
	return 0, fmt.Errorf("no 2 entries sum to %d", n)
}

func solvePart1(inputPath string) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()
	return solve(input, 2020)
}

func solvePart2(inputPath string) (int, error) {
	input1, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input1.Close() }()

	const sumTo int = 2020
	for text, ok, err := input1.ReadLine(); ok || err != nil; text, ok, err = input1.ReadLine() {
		expense, err := strconv.Atoi(text)
		if err != nil {
			return 0, err
		}
		input2, err := aoc.FromFile(inputPath)
		if err != nil {
			return 0, err
		}
		answer, err := solve(input2, sumTo-expense)
		_ = input2.Close()
		if err == nil {
			return answer * expense, nil
		}
	}
	return 0, fmt.Errorf("no 3 entries sum to %d", sumTo)
}
