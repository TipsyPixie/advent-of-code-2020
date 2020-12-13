package day9

import (
	"errors"
	aoc "github.com/TipsyPixie/advent-of-code-2020"
	"strconv"
)

func hasSummationPair(sum int, values ...int) bool {
	pairs := map[int]bool{}
	for _, value := range values {
		if _, ok := pairs[value]; ok {
			return true
		}
		pairs[sum-value] = true
	}
	return false
}

func findSequence(sum int, values ...int) int {
	sequence, sequenceSum := make([]int, 0, len(values)), 0
	valueIndex := 0
	for sequenceSum != sum {
		switch {
		case valueIndex >= len(values):
			return 0
		case sequenceSum < sum:
			sequence = append(sequence, values[valueIndex])
			sequenceSum += values[valueIndex]
			valueIndex++
		case sequenceSum > sum:
			sequenceSum -= sequence[0]
			sequence = sequence[1:]
		}
	}

	min, max := sequence[0], sequence[len(sequence)-1]
	for _, number := range sequence {
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	return min + max
}

func solvePart1(inputPath string) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	preambleLength, index := 25, 0
	preamble := make([]int, 25, 25)
	for line, ok, err := input.ReadLine(); ok || err != nil; line, ok, err = input.ReadLine() {
		if err != nil {
			return 0, err
		}

		number, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		if index >= preambleLength {
			if !hasSummationPair(number, preamble...) {
				return number, nil
			}
		}
		preamble[index%preambleLength] = number
		index++
	}
	return 0, errors.New("not found")
}

func solvePart2(inputPath string) (int, error) {
	invalidNumber, err := solvePart1(inputPath)
	if err != nil {
		return 0, nil
	}

	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	numbers := make([]int, 0, 256)
	for line, ok, err := input.ReadLine(); ok || err != nil; line, ok, err = input.ReadLine() {
		if err != nil {
			return 0, err
		}

		number, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, number)
	}

	return findSequence(invalidNumber, numbers...), nil
}
