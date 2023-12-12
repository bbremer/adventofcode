package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func main() {
	file, _ := os.Open("../inputs/day11.txt")
	// file, _ := os.Open("../inputs/test.txt")
	defer file.Close()

	originalPoints := make([]image.Point, 0)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstRow := scanner.Text()

	shouldExpandRows := make([]bool, 0)
	shouldExpandCols := make([]bool, len(firstRow))
	allPeriods := true
	for j, x := range firstRow {
		if x == '.' {
			shouldExpandCols[j] = true
		} else {
			originalPoints = append(originalPoints, image.Point{j, 0})
			shouldExpandCols[j] = false
			allPeriods = false
		}
	}
	shouldExpandRows = append(shouldExpandRows, allPeriods)

	for i := 1; scanner.Scan(); i++ {
		row := scanner.Text()
		allPeriods := true
		for j, x := range row {
			if x != '.' {
				originalPoints = append(originalPoints, image.Point{j, i})
				shouldExpandCols[j] = false
				allPeriods = false
			}
		}
		shouldExpandRows = append(shouldExpandRows, allPeriods)
	}

	expandRows := make([]int, 0)
	for i, b := range shouldExpandRows {
		if b {
			expandRows = append(expandRows, i)
		}
	}
	expandCols := make([]int, 0)
	for j, b := range shouldExpandCols {
		if b {
			expandCols = append(expandCols, j)
		}
	}

	part1Total := run(originalPoints, expandRows, expandCols, 2)
	part2Total := run(originalPoints, expandRows, expandCols, 1000000)

	fmt.Println("part 1:", part1Total)
	fmt.Println("part 2:", part2Total)
}

func run(originalPoints []image.Point, expandRows, expandCols []int, expandedAmount int) (ret int) {
	expandAmount := expandedAmount - 1
	points := make([]image.Point, 0, len(originalPoints))
	for _, p := range originalPoints {
		newY := p.Y
		for _, i := range expandRows {
			if i < p.Y {
				newY += expandAmount
			} else {
				break
			}
		}
		newX := p.X
		for _, j := range expandCols {
			if j < p.X {
				newX += expandAmount
			} else {
				break
			}
		}
		points = append(points, image.Point{newX, newY})
	}

	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			sub := p1.Sub(p2)
			ret += absInt(sub.X) + absInt(sub.Y)
		}
	}

	return
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
