package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, _ := os.Open("test.txt")
	file, _ := os.Open("../inputs/day4.txt")
	defer file.Close()

	part1_total := 0

	cards := []int{0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cardMatches := parseCard(scanner.Text())
		cards = append(cards, cardMatches)
		part1_total += cardValue(cardMatches)
	}

	part2_total := process(cards)

	fmt.Println("part 1:", part1_total)
	fmt.Println("part 2:", part2_total)
}

func parseCard(line string) int {
	numbers := strings.ReplaceAll(strings.Split(line, ": ")[1], "  ", " ")
	winning_have := strings.Split(numbers, " | ")
	return countWinners(winning_have[0], winning_have[1])
}

func cardValue(numWinning int) int {
	if numWinning == 0 {
		return 0
	}
	return int(math.Pow(2, float64(numWinning-1)))
}

func countWinners(winning string, have string) (count int) {
	for _, h_str := range strings.Split(have, " ") {
		h, _ := strconv.Atoi(h_str)
		if checkWinning(h, winning) {
			count++
		}
	}
	return
}

func checkWinning(h int, winning string) bool {
	for _, w_str := range strings.Split(winning, " ") {
		w, _ := strconv.Atoi(w_str)
		if h == w {
			return true
		}
	}
	return false
}

func process(cards []int) (count int) {
	m := make(map[int]int, len(cards)-1)
	for i := 1; i < len(cards); i++ {
		m[i] = 1
	}
	for i := 1; i < len(cards); i++ {
		count += m[i]
		for j := 1; j <= cards[i]; j++ {
			m[i+j] += m[i]
		}
	}
	return
}
