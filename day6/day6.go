package day6

import (
	aoc "github.com/TipsyPixie/advent-of-code-2020"
)

func getBinaryRepresentation(form string) int {
	binaryRepresentation := 0
	for _, answer := range []rune(form) {
		binaryRepresentation += 1 << int(answer-'a')
	}
	return binaryRepresentation
}

func or(values ...int) int {
	result := 0
	for _, value := range values {
		result |= value
	}
	return result
}

func and(values ...int) int {
	result := 0b11111111111111111111111111
	for _, value := range values {
		result &= value
	}
	return result
}

func countOnes(value int) int {
	count := 0
	for value > 0 {
		if value&1 == 1 {
			count++
		}
		value = value >> 1
	}
	return count
}

func solve(inputPath string, combine func(values ...int) int) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	totalSum := 0
	buffer := make([]int, 0, 4)
	for form, ok, err := input.ReadLine(); ok || err != nil; form, ok, err = input.ReadLine() {
		if err != nil {
			return 0, nil
		}
		if form != "" {
			buffer = append(buffer, getBinaryRepresentation(form))
		} else {
			totalSum += countOnes(combine(buffer...))
			buffer = buffer[:0]
		}
	}
	return totalSum + countOnes(combine(buffer...)), nil
}

func solvePart1(inputPath string) (int, error) {
	return solve(inputPath, or)
}

func solvePart2(inputPath string) (int, error) {
	return solve(inputPath, and)
}
