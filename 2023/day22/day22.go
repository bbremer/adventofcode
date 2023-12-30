package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day22.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bricks := make([]*Brick, 0)
	for scanner.Scan() {
		bricks = append(bricks, NewBrick(scanner.Text()))
	}
	slices.SortFunc(bricks, func(a, b *Brick) int {
		return cmp.Compare(a.zs.min, b.zs.max)
	})
	for i, b := range bricks {
		b.Fall(bricks[:i])
	}
	part1Total := 0
	part2Total := 0
	for _, b := range bricks {
		if b.Disintegratable() {
			part1Total++
		}
		part2Total += b.NumWouldFall()
	}
	fmt.Println("Part 1:", part1Total)
	fmt.Println("Part 2:", part2Total)
}

type Brick struct {
	xs          Interval
	ys          Interval
	zs          *Interval
	SupportedBy []*Brick
	Supports    []*Brick
}

func NewBrick(line string) *Brick {
	splitLine := strings.Split(line, "~")
	p1 := ParsePoint(strings.Split(splitLine[0], ","))
	p2 := ParsePoint(strings.Split(splitLine[1], ","))
	xs := Interval{p1[0], p2[0]}
	ys := Interval{p1[1], p2[1]}
	zs := Interval{p1[2], p2[2]}
	return &Brick{xs, ys, &zs, make([]*Brick, 0), make([]*Brick, 0)}
}

func (b *Brick) Fall(bricks []*Brick) {
	conflicts := make([]*Brick, 0)
	for _, b2 := range bricks {
		if b.Intersects(b2) {
			conflicts = append(conflicts, b2)
		}
	}
	for {
		if b.zs.min <= 0 || b.IsSupported(conflicts) {
			break
		}
		b.zs.Fall()
	}
}

func (b *Brick) IsSupported(bricks []*Brick) bool {
	targetZ := b.zs.min - 1
	ret := false
	for _, b2 := range bricks {
		if b2.zs.max == targetZ {
			b.SupportedBy = append(b.SupportedBy, b2)
			b2.Supports = append(b2.Supports, b)
			ret = true
		}
	}
	return ret
}

func (b *Brick) Intersects(b2 *Brick) bool {
	return b.xs.Intersects(b2.xs) && b.ys.Intersects(b2.ys)
}

func (b *Brick) Disintegratable() bool {
	for _, b2 := range b.Supports {
		if len(b2.SupportedBy) < 2 {
			return false
		}
	}
	return true
}

func (b *Brick) NumWouldFall() int {
	wouldFall := map[*Brick]bool{b: true}
	stack := make([]*Brick, len(b.Supports))
	copy(stack, b.Supports)
	for {
		l := len(stack)
		if l == 0 {
			break
		}
		b2 := stack[l-1]
		stack = stack[:l-1]

		if _, ok := wouldFall[b2]; ok {
			continue
		}
		thisWouldFall := true
		for _, b3 := range b2.SupportedBy {
			if _, ok := wouldFall[b3]; !ok {
				thisWouldFall = false
				break
			}
		}
		if thisWouldFall {
			wouldFall[b2] = true
			stack = append(stack, b2.Supports...)
		}
	}
	return len(wouldFall) - 1
}

func ParsePoint(s []string) [3]int {
	p1, _ := strconv.Atoi(s[0])
	p2, _ := strconv.Atoi(s[1])
	p3, _ := strconv.Atoi(s[2])
	return [3]int{p1, p2, p3}
}

type Interval struct {
	min int
	max int
}

func (i Interval) Intersects(i2 Interval) bool {
	return i.max >= i2.min && i.min <= i2.max
}

func (i *Interval) Fall() {
	i.min--
	i.max--
}
