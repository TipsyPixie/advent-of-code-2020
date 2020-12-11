package day4

import (
	aoc "github.com/TipsyPixie/advent-of-code-2020"
	"regexp"
	"strconv"
	"strings"
)

func parsePassport(passport string) map[string]string {
	attributes := map[string]string{}
	keyValuePairs := strings.Split(passport, " ")
	for _, keyValuePair := range keyValuePairs {
		pair := strings.Split(keyValuePair, ":")
		attributes[pair[0]] = pair[1]
	}
	return attributes
}

func solve(inputPath string, validate func(string) bool) (int, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return 0, err
	}
	defer func() { _ = input.Close() }()

	validPassportCount := 0
	passportBuilder := strings.Builder{}
	composeAndValidate := func(builder *strings.Builder) bool {
		passport := builder.String()
		builder.Reset()
		return validate(passport)
	}
	for passportText, ok, err := input.ReadLine(); ok || err != nil; passportText, ok, err = input.ReadLine() {
		switch {
		case err != nil:
			return 0, err
		case passportText == "":
			if composeAndValidate(&passportBuilder) {
				validPassportCount++
			}
		default:
			if passportBuilder.Len() > 0 {
				_, err = passportBuilder.WriteString(" ")
				if err != nil {
					return 0, err
				}
			}
			_, err = passportBuilder.WriteString(passportText)
			if err != nil {
				return 0, err
			}
		}
	}

	if composeAndValidate(&passportBuilder) {
		validPassportCount++
	}
	return validPassportCount, nil
}

func solvePart1(inputPath string) (int, error) {
	return solve(
		inputPath,
		func(passport string) bool {
			requiredAttributes := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
			attributes := parsePassport(passport)
			for _, requiredAttribute := range requiredAttributes {
				if _, hasAttribute := attributes[requiredAttribute]; !hasAttribute {
					return false
				}
			}
			return true
		},
	)
}

func validateAttribute(key string, value string) bool {
	var err error = nil
	formatOk, valueOk := true, true
	comparableValue := 0
	switch key {
	case "byr":
		formatOk, err = regexp.MatchString("^[0-9]{4}$", value)
		comparableValue, err = strconv.Atoi(value)
		valueOk = comparableValue >= 1920 && comparableValue <= 2002
	case "iyr":
		formatOk, err = regexp.MatchString("^[0-9]{4}$", value)
		comparableValue, err = strconv.Atoi(value)
		valueOk = comparableValue >= 2010 && comparableValue <= 2020
	case "eyr":
		formatOk, err = regexp.MatchString("^[0-9]{4}$", value)
		comparableValue, err = strconv.Atoi(value)
		valueOk = comparableValue >= 2020 && comparableValue <= 2030
	case "hgt":
		formatOk, err = regexp.MatchString("^[0-9]+(cm|in)$", value)
		comparableValue, err = strconv.Atoi(value[:len(value)-2])
		valueOk = (value[len(value)-2:] == "cm" && comparableValue >= 150 && comparableValue <= 193) ||
			(value[len(value)-2:] == "in" && comparableValue >= 59 && comparableValue <= 76)
	case "hcl":
		formatOk, err = regexp.MatchString("^#[0-9a-f]{6}$", value)
	case "ecl":
		formatOk, err = regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", value)
	case "pid":
		formatOk, err = regexp.MatchString("^[0-9]{9}$", value)
	}
	return err == nil && formatOk && valueOk
}

func solvePart2(inputPath string) (int, error) {
	return solve(
		inputPath,
		func(passport string) bool {
			requiredAttributes := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
			attributes := parsePassport(passport)
			for _, requiredAttribute := range requiredAttributes {
				if value, hasAttribute := attributes[requiredAttribute]; !hasAttribute || !validateAttribute(requiredAttribute, value) {
					return false
				}
			}
			return true
		},
	)
}
