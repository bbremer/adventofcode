package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func line_to_number(line string) int {
	var first, last rune
	for _, char := range line {
		if '0' <= char && char <= '9' {
			if first == 0 {
				first = char
			}
			last = char
		}
	}
	num_str := fmt.Sprintf("%c%c", first, last)
	num, _ := strconv.Atoi(num_str)
	return num
}

func main() {
	file, _ := os.Open("inputs/day1.txt")
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += line_to_number(scanner.Text())
	}

	fmt.Println(total)
}
