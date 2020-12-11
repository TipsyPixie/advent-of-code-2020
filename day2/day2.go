package day2

import (
	"errors"
	"fmt"
	aoc "github.com/TipsyPixie/advent-of-code-2020"
	"regexp"
	"strconv"
	"strings"
)

func parsePolicy(policy string) (rune, int, int, string, error) {
	formatOk, err := regexp.MatchString("^[0-9]+-[0-9]+ [a-z]: [a-z]+$", policy)
	if err != nil {
		return 0, 0, 0, "", err
	}
	if !formatOk {
		return 0, 0, 0, "", errors.New(fmt.Sprintf("invalid policy %s", policy))
	}

	splitPolicy := strings.SplitN(policy, ": ", 2)
	rule, password := splitPolicy[0], splitPolicy[1]
	splitRule := strings.SplitN(rule, " ", 2)
	frequencyRange, character := splitRule[0], []rune(splitRule[1])[0]
	splitFrequencyRange := strings.SplitN(frequencyRange, "-", 2)
	minFrequency, err := strconv.Atoi(splitFrequencyRange[0])
	if err != nil {
		return 0, 0, 0, "", err
	}
	maxFrequency, err := strconv.Atoi(splitFrequencyRange[1])
	if err != nil {
		return 0, 0, 0, "", err
	}
	return character, minFrequency, maxFrequency, password, nil
}

func solve(inputPath string, validate func(rune, int, int, string) bool) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	validPasswordCount := 0
	for policy, ok, err := input.ReadLine(); ok || err != nil; policy, ok, err = input.ReadLine() {
		if err != nil {
			return 0, err
		}
		ruleCharacter, minFrequency, maxFrequency, password, err := parsePolicy(policy)
		if err != nil {
			return 0, err
		}
		if validate(ruleCharacter, minFrequency, maxFrequency, password) {
			validPasswordCount++
		}
	}
	return validPasswordCount, nil
}

func solvePart1(inputPath string) (int, error) {
	return solve(
		inputPath,
		func(ruleCharacter rune, minFrequency int, maxFrequency int, password string) bool {
			ruleCharacterCount := 0
			for _, character := range []rune(password) {
				if character == ruleCharacter {
					ruleCharacterCount++
				}
			}
			return ruleCharacterCount <= maxFrequency && ruleCharacterCount >= minFrequency
		},
	)
}

func solvePart2(inputPath string) (int, error) {
	return solve(
		inputPath,
		func(ruleCharacter rune, position1 int, position2 int, password string) bool {
			matchedPositionCount, passwordCharacters := 0, []rune(password)
			for _, position := range []int{position1, position2} {
				if passwordCharacters[position-1] == ruleCharacter {
					matchedPositionCount++
				}
			}
			return matchedPositionCount == 1
		},
	)
}
