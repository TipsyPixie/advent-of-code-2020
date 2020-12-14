package day10

import (
	aoc "github.com/TipsyPixie/advent-of-code-2020"
	"sort"
	"strconv"
)

func toSlice(inputPath string) ([]int, error) {
	result := make([]int, 0, 256)
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = input.Close() }()

	for line, ok, err := input.ReadLine(); ok || err != nil; line, ok, err = input.ReadLine() {
		if err != nil {
			return nil, err
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		result = append(result, number)
	}
	return result, nil
}

func solvePart1(inputPath string) (int, error) {
	adapters, err := toSlice(inputPath)
	if err != nil {
		return 0, err
	}

	sort.Ints(adapters)
	lastAdapter, joltDifferences := 0, map[int]int{3: 1}
	for _, adapter := range adapters {
		joltDifferences[adapter-lastAdapter] += 1
		lastAdapter = adapter
	}
	return joltDifferences[1] * joltDifferences[3], nil
}

func countArrangements(numbers []int, lastNumber int, arrangements map[[2]int]int) int {
	if len(numbers) <= 1 {
		return 1
	}
	if cached, ok := arrangements[[2]int{len(numbers), lastNumber}]; ok {
		return cached
	}
	currentNumber, nextNumber := numbers[0], numbers[1]
	count := countArrangements(numbers[1:], currentNumber, arrangements)
	if nextNumber-lastNumber <= 3 {
		count += countArrangements(numbers[1:], lastNumber, arrangements)
	}
	arrangements[[2]int{len(numbers), lastNumber}] = count
	return count
}

func solvePart2(inputPath string) (int, error) {
	adapters, err := toSlice(inputPath)
	if err != nil {
		return 0, err
	}

	sort.Ints(adapters)
	return countArrangements(adapters, 0, map[[2]int]int{}), nil
}
