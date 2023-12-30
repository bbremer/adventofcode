package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../inputs/day5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seeds := scannerToSeeds(scanner)
	fn := scannerToFullMap(scanner)
	var minResult int64 = 1000000000
	for _, s := range seeds {
		result := fn(s)
		if result < minResult {
			minResult = result
		}
	}
	fmt.Println("part 1:", minResult)

	part2Seeds := seedsToRanges(seeds)
	fmt.Println("part 2:", part2(part2Seeds, fn))
}

func scannerToSeeds(scanner *bufio.Scanner) []int64 {
	scanner.Scan()
	seedsLine := scanner.Text()
	seedStrs := strings.Split(strings.Split(seedsLine, ": ")[1], " ")
	seeds := make([]int64, 0)
	for _, seedStr := range seedStrs {
		seed, _ := strconv.ParseInt(seedStr, 10, 64)
		seeds = append(seeds, seed)
	}
	scanner.Scan() // blank line
	return seeds
}

func scannerToFullMap(scanner *bufio.Scanner) func(int64) int64 {
	maps := scannerToMaps(scanner)
	return funcsToFunc(maps)
}

func scannerToMaps(scanner *bufio.Scanner) []func(int64) int64 {
	maps := make([]func(int64) int64, 0)

	currRanges := make([]*Range, 0)

	scanner.Scan() // get to first "map" line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			maps = append(maps, rangesToFunc(currRanges))
			currRanges = make([]*Range, 0)
			scanner.Scan() // get to next "map" line
			continue
		}
		currRanges = append(currRanges, NewRange(line))
	}
	maps = append(maps, rangesToFunc(currRanges))

	return maps
}

type Range struct {
	destStart int64
	srcStart  int64
	rangeLen  int64
}

func NewRange(line string) *Range {
	splitLine := strings.Split(line, " ")
	destStart, _ := strconv.ParseInt(splitLine[0], 10, 64)
	srcStart, _ := strconv.ParseInt(splitLine[1], 10, 64)
	rangeLen, _ := strconv.ParseInt(splitLine[2], 10, 64)
	return &Range{destStart, srcStart, rangeLen}
}

func (r *Range) Contains(x int64) (int64, bool) {
	if r.srcStart <= x && x < r.srcStart+r.rangeLen {
		return r.destStart + x - r.srcStart, true
	}
	return x, false
}

func rangesToFunc(ranges []*Range) func(int64) int64 {
	return func(x int64) int64 {
		for _, r := range ranges {
			if y, in := r.Contains(x); in {
				return y
			}
		}
		return x
	}
}

func funcsToFunc(fns []func(int64) int64) func(int64) int64 {
	return func(i int64) int64 {
		for _, fn := range fns {
			i = fn(i)
		}
		return i
	}
}

func seedsToRanges(seeds []int64) [][2]int64 {
	ranges := make([][2]int64, 0)
	for i := 0; i < len(seeds); i = i + 2 {
		ranges = append(ranges, [2]int64{seeds[i], seeds[i+1]})
	}
	return ranges
}

func part2(seedRanges [][2]int64, fn func(int64) int64) int64 {
	var min int64 = 100000000000
	for i, r := range seedRanges {
		fmt.Println(i, r)
		for j := r[0]; j < r[0]+r[1]; j++ {
			x := fn(j)
			if x < min {
				min = x
			}
		}
	}
	return min
}
