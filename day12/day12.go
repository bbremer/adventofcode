package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../inputs/day12.txt")
	// file, _ := os.Open("../inputs/test.txt")
	defer file.Close()

	part1Total := 0
	part2Total := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		part1Total += parseLine(line)
		part2Total += parseLine(unfold(line))
	}

	fmt.Println("part 1:", part1Total)
	fmt.Println("part 2:", part2Total)
}

func parseLine(line string) (count int) {
	cache = make(map[string]int, 10000)

	splitLine := strings.Split(line, " ")
	conditionStrings := strings.Split(splitLine[1], ",")
	conditions := make([]int, 0)
	for _, x := range conditionStrings {
		y, _ := strconv.Atoi(x)
		conditions = append(conditions, y)
	}

	springs := splitLine[0]

	count = countValid2(springs, conditions, 0)

	return
}

var cache map[string]int

func countValid2(springs string, condition []int, numDefective int) (count int) {
	cacheKey := stringify(condition, numDefective, springs)
	if cachedCount, ok := cache[cacheKey]; ok {
		return cachedCount
	}

	if len(springs) == 0 {
		if len(condition) == 0 {
			if numDefective == 0 {
				return 1
			}
			return 0
		}
		if len(condition) == 1 && condition[0] == numDefective {
			return 1
		}
		return 0
	}

	if len(condition) == 0 {
		for i := 0; i < len(springs); i++ {
			if springs[i] == '#' {
				return 0
			}
		}
		return 1
	}

	originalCondition := condition

	var newSprings string
	if len(springs) > 1 {
		newSprings = springs[1:]
	}
	switch springs[0] {
	case '.':
		if numDefective > 0 {
			if numDefective != condition[0] {
				return 0
			}
			if len(condition) > 1 {
				condition = condition[1:]
			} else {
				condition = []int{}
			}
		}
		count = countValid2(newSprings, condition, 0)
	case '#':
		count = countValid2(newSprings, condition, numDefective+1)
	case '?':
		count = countValid2("."+newSprings, condition, numDefective) + countValid2("#"+newSprings, condition, numDefective)
	}
	cacheKey = stringify(originalCondition, numDefective, springs)
	cache[cacheKey] = count
	return
}

func unfold(line string) string {
	splitLine := strings.Split(line, " ")
	springs := splitLine[0]
	conditionStrings := splitLine[1]
	springs2 := make([]string, 0, 5)
	conditions2 := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
		springs2 = append(springs2, springs)
		conditions2 = append(conditions2, conditionStrings)
	}
	springs = strings.Join(springs2, "?")
	conditions := strings.Join(conditions2, ",")
	return springs + " " + conditions
}

func stringify(condition []int, numDefective int, originalSprings string) string {
	return fmt.Sprintf("%v%d%s", condition, numDefective, originalSprings)
}
