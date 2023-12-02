package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var m = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var m_reversed = make(map[string]string, len(m))
var re *regexp.Regexp
var re_reversed *regexp.Regexp
var replace func(string) string
var replace_reverse func(string) string

func init() {
	re = make_re_from_map(m)

	for str, num := range m {
		m_reversed[reverse(str)] = num
	}
	re_reversed = make_re_from_map(m_reversed)

	replace = make_replace(m)
	replace_reverse = make_replace(m_reversed)
}

func make_re_from_map(m map[string]string) (re *regexp.Regexp) {
	str_numbers := make([]string, 0, len(m)+1)
	for str := range m {
		str_numbers = append(str_numbers, str)
	}
	numbers := append(str_numbers, `\d`)
	numbers_pattern := strings.Join(numbers, "|")
	re = regexp.MustCompile(numbers_pattern)
	return
}

func make_replace(m map[string]string) func(string) string {
	return func(in string) string {
		match, ok := m[in]
		if !ok {
			return in
		}
		return match
	}
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func line_to_number_part1(line string) int {
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

func replace_first(str string) string {
	found := re.FindString(str)
	return strings.Replace(str, found, replace(found), 1)
}

func replace_last(str string) string {
	reversed_str := reverse(str)
	found := re_reversed.FindString(reversed_str)
	new_str := strings.Replace(reversed_str, found, replace_reverse(found), 1)
	return reverse(new_str)
}

func line_to_number_part2(line string) int {
	line = strings.TrimSuffix(line, "\n")
	first_replaced := replace_first(line)
	last_replaced := replace_last(first_replaced)
	return line_to_number_part1(last_replaced)
}

func main() {
	file, _ := os.Open("inputs/day1.txt")
	defer file.Close()

	part1_total := 0
	part2_total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		part1_total += line_to_number_part1(scanner.Text())
		part2_total += line_to_number_part2(scanner.Text())
	}

	fmt.Println("part 1:", part1_total)
	fmt.Println("part 2:", part2_total)
}
