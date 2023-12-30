package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../inputs/day6.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeNumbers := parseLine(scanner.Text())
	time := parseLine2(scanner.Text())
	scanner.Scan()
	distanceNumbers := parseLine(scanner.Text())
	distance := parseLine2(scanner.Text())
	part1_total := part1(timeNumbers, distanceNumbers)
	part2_total := numWays(time, distance)

	fmt.Println("part 1:", part1_total)
	fmt.Println("part 2:", part2_total)
}

func parseLine(line string) (numbers []int) {
	numbersStr := strings.Split(line, ":")[1]
	numbersStrs := strings.Fields(numbersStr)
	for _, numStr := range numbersStrs {
		number, _ := strconv.Atoi(numStr)
		numbers = append(numbers, number)
	}
	return
}

func parseLine2(line string) int {
	numbersStr := strings.Split(line, ":")[1]
	numbersStr2 := strings.ReplaceAll(numbersStr, " ", "")
	number, _ := strconv.Atoi(numbersStr2)
	return number
}

func part1(times, records []int) int {
	ways := 1
	for i := 0; i < len(times); i++ {
		ways *= numWays(times[i], records[i])
	}
	return ways
}

func numWays(totalTime, record int) (count int) {
	for i := 1; i < totalTime; i++ {
		distance := i * (totalTime - i)
		if distance > record {
			count++
		}
	}
	return
}
