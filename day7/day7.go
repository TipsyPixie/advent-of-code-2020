package day7

import (
	aoc "github.com/TipsyPixie/advent-of-code-2020"
	"regexp"
	"strconv"
)

type bag struct {
	color              string
	outerBags          []*bag
	outerBagCapacities []int
	innerBags          []*bag
	innerBagCapacities []int
}

func parseRule(rule string) (outerBagColor string, innerBagColors []string, innerBagCounts []int, err error) {
	outerBagColorPattern := regexp.MustCompile("^[a-z]+ [a-z]+")
	outerBagColor = outerBagColorPattern.FindString(rule)

	innerBagPattern := regexp.MustCompile("[0-9]+ [a-z]+ [a-z]+ bags?")
	innerBags := innerBagPattern.FindAllString(rule, -1)
	countPattern, colorPattern := regexp.MustCompile("^[0-9]+"), regexp.MustCompile("[a-z]+ [a-z]+")
	innerBagColors, innerBagCounts = make([]string, 0, len(innerBags)), make([]int, 0, len(innerBags))
	for _, child := range innerBags {
		innerBagCount, err := strconv.Atoi(countPattern.FindString(child))
		if err != nil {
			return "", nil, nil, err
		}
		innerBagColor := colorPattern.FindString(child)
		innerBagColors, innerBagCounts = append(innerBagColors, innerBagColor), append(innerBagCounts, innerBagCount)
	}
	return
}

func appendBag(bags map[string]*bag, outerBagColor string, innerBagColor string, bagCapacity int) {
	innerBag, isInnerBagInTree := bags[innerBagColor]
	if !isInnerBagInTree {
		innerBag = &bag{
			color:              innerBagColor,
			outerBags:          make([]*bag, 0, 1),
			outerBagCapacities: make([]int, 0, 1),
		}
		bags[innerBagColor] = innerBag
	}

	outerBag, isOuterBagInTree := bags[outerBagColor]
	if !isOuterBagInTree {
		outerBag = &bag{
			color:              outerBagColor,
			outerBags:          make([]*bag, 0, 1),
			outerBagCapacities: make([]int, 0, 1),
		}
		bags[outerBagColor] = outerBag
	}

	innerBag.outerBags = append(innerBag.outerBags, outerBag)
	innerBag.outerBagCapacities = append(innerBag.outerBagCapacities, bagCapacity)
	outerBag.innerBags = append(outerBag.innerBags, innerBag)
	outerBag.innerBagCapacities = append(outerBag.innerBagCapacities, bagCapacity)
}

func countOuterColors(rootBag *bag, alreadyCounted map[string]bool) int {
	count := 0
	for _, outerBag := range rootBag.outerBags {
		if outerColorAlreadyCounted := alreadyCounted[outerBag.color]; !outerColorAlreadyCounted {
			count++
			alreadyCounted[outerBag.color] = true
		}
		count += countOuterColors(outerBag, alreadyCounted)
	}
	return count
}

func getCapacityTreeSum(rootBag *bag) int {
	sum := 0
	for index, innerBag := range rootBag.innerBags {
		sum += (1 + getCapacityTreeSum(innerBag)) * rootBag.innerBagCapacities[index]
	}
	return sum
}

func buildTree(inputPath string) (*bag, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = input.Close() }()

	rootBag := &bag{
		outerBags:          make([]*bag, 0, 2),
		outerBagCapacities: make([]int, 0, 2),
	}
	bags := map[string]*bag{
		"shiny gold": rootBag,
	}
	for rule, ok, err := input.ReadLine(); ok || err != nil; rule, ok, err = input.ReadLine() {
		if err != nil {
			return nil, err
		}
		parentColor, childrenColors, outerBagCapacities, err := parseRule(rule)
		if err != nil {
			return nil, err
		}
		for index, childColor := range childrenColors {
			appendBag(bags, parentColor, childColor, outerBagCapacities[index])
		}
	}
	return rootBag, nil
}

func solvePart1(inputPath string) (int, error) {
	rootBag, err := buildTree(inputPath)
	if err != nil {
		return 0, err
	}
	return countOuterColors(rootBag, map[string]bool{}), nil
}

func solvePart2(inputPath string) (int, error) {
	rootBag, err := buildTree(inputPath)
	if err != nil {
		return 0, err
	}
	return getCapacityTreeSum(rootBag), nil
}
