package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day13.txt")
	defer file.Close()

	part1Total := 0
	part2Total := 0

	scanner := bufio.NewScanner(file)

	currPattern := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			x, _ := part1(currPattern)
			part1Total += x
			part2Total += part2(currPattern, x)
			currPattern = make([]string, 0)
			continue
		}
		currPattern = append(currPattern, line)
	}
	if len(currPattern) > 0 {
		x, _ := part1(currPattern)
		part1Total += x
		part2Total += part2(currPattern, x)
	}

	fmt.Println("part 1:", part1Total)
	fmt.Println("part 2:", part2Total)
}

func part1(pattern []string) (int, bool) {
	if count, ok := searchRows(pattern); ok {
		return count * 100, true
	}
	patternTransposed := transpose(pattern)
	if count, ok := searchRows(patternTransposed); ok {
		return count, true
	}
	return 0, false
}

func searchRows(pattern []string) (int, bool) {
	for i := 1; i < len(pattern); i++ {
		if pattern[i-1] == pattern[i] {
			for j := 1; ; j++ {
				if i+j >= len(pattern) {
					return i, true
				}
				if i-1-j < 0 {
					return i, true
				}
				if pattern[i+j] != pattern[i-1-j] {
					break
				}
			}
		}
	}
	return 0, false
}

func transpose(pattern []string) []string {
	transposed := make([]string, 0, len(pattern))
	for i := range pattern[0] {
		var newString bytes.Buffer
		for _, s := range pattern {
			newString.WriteByte(s[i])
		}
		transposed = append(transposed, newString.String())
	}
	return transposed
}

func part2(pattern []string, part1Count int) int {
	for i := range pattern {
		originalString := pattern[i]
		byteSlice := []byte(originalString)
		for j := range pattern[0] {
			byteSlice[j] = flipChar(byteSlice[j])
			pattern[i] = string(byteSlice)
			if count, ok := part1Wrapper(pattern, part1Count); ok {
				return count
			}
			byteSlice[j] = flipChar(byteSlice[j])
		}
		pattern[i] = originalString
	}

	fmt.Println("problem")
	os.Exit(1)
	return 0
}

func flipChar(x byte) byte {
	if x == '.' {
		return '#'
	} else if x == '#' {
		return '.'
	}
	fmt.Println("flipChar problem")
	os.Exit(1)
	return 0
}

func part1Wrapper(pattern []string, part1Count int) (int, bool) {
	if count, ok := searchRows(pattern); ok {
		if count*100 != part1Count {
			return count * 100, true
		}
		if part1Count >= 100 {
			cutPattern := pattern[part1Count/100:]
			if count, ok := searchRows(cutPattern); ok {
				return (count + part1Count/100) * 100, true
			}
		}
	}
	patternTransposed := transpose(pattern)
	if count, ok := searchRows(patternTransposed); ok {
		if count != part1Count {
			return count, true
		}
		if part1Count < 100 {
			cutPattern := patternTransposed[part1Count:]
			if count, ok := searchRows(cutPattern); ok {
				return count + part1Count, true
			}
		}
	}
	return 0, false
}
