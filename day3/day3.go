package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type PartNumber struct {
	ID        int
	row       int
	start_col int
	end_col   int
}

func NewPartNumber(lines []string, row, start_col int) *PartNumber {
	end_col := start_col
	id_bytes := []byte{lines[row][start_col]}
	for c := start_col + 1; c < len(lines[row]); c++ {
		if !isDigit(lines[row][c]) {
			break
		}
		end_col = c
		id_bytes = append(id_bytes, lines[row][c])
	}
	id, _ := strconv.Atoi(string(id_bytes))
	return &PartNumber{id, row, start_col, end_col}
}

func (pn PartNumber) isAdjacentToSymbol(lines []string) bool {
	return symbolInRowRange(lines, pn.row-1, pn.start_col-1, pn.end_col+1) ||
		symbolInRowRange(lines, pn.row+1, pn.start_col-1, pn.end_col+1) ||
		isSymbolBounded(lines, pn.row, pn.start_col-1) ||
		isSymbolBounded(lines, pn.row, pn.end_col+1)
}

func (pn PartNumber) adjacentToLocation(row, col int) bool {
	return pn.row-1 <= row && row <= pn.row+1 &&
		pn.start_col-1 <= col && col <= pn.end_col+1
}

func (pn PartNumber) asterisksAdjacencies(lines []string) {
	for point := range asterisks {
		if pn.adjacentToLocation(point[0], point[1]) {
			asterisks[point] = append(asterisks[point], pn.ID)
		}
	}
}

func symbolInRowRange(lines []string, row, start_col, end_col int) bool {
	if !inBoundsRow(lines, row) {
		return false
	}
	for c := start_col; c <= end_col; c++ {
		if isSymbolBounded(lines, row, c) {
			return true
		}
	}
	return false
}

func isDigit(x byte) bool {
	return '0' <= x && x <= '9'
}

func isPeriod(x byte) bool {
	return x == '.'
}

func isSymbol(x byte) bool {
	return !(isDigit(x) || isPeriod(x))
}

func inBoundsRow(lines []string, row int) bool {
	return 0 <= row && row < len(lines)
}

func inBounds(lines []string, row, col int) bool {
	if !inBoundsRow(lines, row) {
		return false
	}
	maxCol := len(lines[row]) - 1
	if col < 0 || maxCol < col {
		return false
	}
	return true
}

func isSymbolBounded(lines []string, row, col int) bool {
	return inBounds(lines, row, col) && isSymbol(lines[row][col])
}

func sumLine(lines []string, row int) (sum int) {
	line := lines[row]
	c := 0
	for c < len(line) {
		if !isDigit(line[c]) {
			c++
			continue
		}
		partNumber := NewPartNumber(lines, row, c)
		if partNumber.isAdjacentToSymbol(lines) {
			sum += partNumber.ID
			c = partNumber.end_col
			partNumber.asterisksAdjacencies(lines)
		}
		c++
	}
	return
}

func assertASCII(line string) {
	for i := 0; i < len(line); i++ {
		if line[i] > unicode.MaxASCII {
			fmt.Fprintf(os.Stderr, "Line has non-ASCII characters: %s", line)
			os.Exit(1)
		}
	}
}

var asterisks = make(map[[2]int][]int)

func addLineAsterisks(line string, row int) {
	for c, char := range line {
		if char == '*' {
			asterisks[[2]int{row, c}] = make([]int, 0)
		}
	}
}

func main() {
	file, _ := os.Open("../inputs/day3.txt")
	defer file.Close()

	part1Total := 0
	part2Total := 0

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	r := 0
	for scanner.Scan() {
		line := scanner.Text()
		assertASCII(line)
		lines = append(lines, line)
		addLineAsterisks(line, r)
		r++
	}

	for r := 0; r < len(lines); r++ {
		part1Total += sumLine(lines, r)
	}

	for _, a := range asterisks {
		if len(a) == 2 {
			part2Total += a[0] * a[1]
		}
	}

	fmt.Println("part 1:", part1Total)
	fmt.Println("part 2:", part2Total)
}
