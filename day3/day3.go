package day3

import aoc "github.com/TipsyPixie/advent-of-code-2020"

func solve(inputPath string, stepRight int, stepDown int) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	treeCount := 0
	goRight, goDown := stepRight, stepDown+1
	currentRight := 0
	for terrain, ok, err := input.ReadLine(); ok || err != nil; terrain, ok, err = input.ReadLine() {
		if err != nil {
			return 0, err
		}
		goDown--
		if goDown > 0 {
			continue
		}
		currentRight = (currentRight + goRight) % len(terrain)
		if terrain[currentRight:currentRight+1] == "#" {
			treeCount++
		}
		goDown = stepDown
	}
	return treeCount, nil
}

func solvePart1(inputPath string) (int, error) {
	return solve(inputPath, 3, 1)
}

func solvePart2(inputPath string) (int, error) {
	treeCountProduction := 1
	steps := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, step := range steps {
		treeCount, err := solve(inputPath, step[0], step[1])
		if err != nil {
			return 0, err
		}
		treeCountProduction *= treeCount
	}
	return treeCountProduction, nil
}
