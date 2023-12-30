package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../inputs/day9.txt")
	defer file.Close()

	part1Total := 0
	part2Total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret1, ret2 := parse(scanner.Text())
		part1Total += ret1
		part2Total += ret2
	}

	fmt.Println("part 1:", part1Total)
	fmt.Println("part 2:", part2Total)
}

func parse(line string) (int, int) {
	numbers := make([]int, 0)
	for _, x := range strings.Split(line, " ") {
		number, _ := strconv.Atoi(x)
		numbers = append(numbers, number)
	}
	currNumbers := numbers
	lines := [][]int{currNumbers}
	for !allEqual(currNumbers) {
		currNumbers = nextLine(currNumbers)
		lines = append(lines, currNumbers)
	}

	forwardExtrapolated := lines[len(lines)-1][0]
	reverseExtrapolated := lines[len(lines)-1][0]
	for i := len(lines) - 2; i >= 0; i-- {
		forwardExtrapolated += lines[i][len(lines[i])-1]
		reverseExtrapolated = lines[i][0] - reverseExtrapolated
	}

	return forwardExtrapolated, reverseExtrapolated
}

func nextLine(oldLine []int) []int {
	line := make([]int, 0, len(oldLine)-1)
	for i, x := range oldLine[1:] {
		line = append(line, x-oldLine[i])
	}
	return line
}

func allEqual(numbers []int) bool {
	x := numbers[0]
	for _, y := range numbers[1:] {
		if y != x {
			return false
		}
	}
	return true
}
