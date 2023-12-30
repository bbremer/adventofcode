package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^Game (\d*): (.*)$`)
var max_per_color = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func drawPossible(draws_str string) bool {
	for _, draw := range strings.Split(draws_str, ", ") {
		draw_split := strings.Split(draw, " ")
		number, _ := strconv.Atoi(draw_split[0])
		color := draw_split[1]
		if number > max_per_color[color] {
			return false
		}
	}
	return true
}

func gamePossible(games_str string) bool {
	for _, game := range strings.Split(games_str, "; ") {
		if !drawPossible(game) {
			return false
		}
	}
	return true
}

func set_power(games_str string) int {
	m := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
	for _, draws_str := range strings.Split(games_str, "; ") {

		for _, draw := range strings.Split(draws_str, ", ") {
			draw_split := strings.Split(draw, " ")
			number, _ := strconv.Atoi(draw_split[0])
			color := draw_split[1]
			if number > m[color] {
				m[color] = number
			}
		}
	}
	prod := 1
	for _, number := range m {
		prod *= number
	}
	return prod
}

func line_to_id_and_games_str(line string) (id int, games_str string) {
	submatches := re.FindStringSubmatch(line)
	id, _ = strconv.Atoi(submatches[1])
	games_str = submatches[2]
	return
}

func main() {
	fmt.Println(max_per_color)

	file, _ := os.Open("../inputs/day2.txt")
	defer file.Close()

	part1_total := 0
	part2_total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id, games_str := line_to_id_and_games_str(scanner.Text())
		if gamePossible(games_str) {
			part1_total += id
		}
		part2_total += set_power(games_str)
	}

	fmt.Println("part 1:", part1_total)
	fmt.Println("part 2:", part2_total)
}
