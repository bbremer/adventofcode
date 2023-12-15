package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day15.txt")
	defer file.Close()

	part1Total := 0

	b := make([]byte, 1)
	currHash := 0
	labelHash := 0
	currLabel := make([]byte, 0)
	inFocalLen := false
	currFocalLen := make([]byte, 0)
	var boxes [256][]*Lens
	for i := 0; i < 256; i++ {
		boxes[i] = make([]*Lens, 0)
	}
	for {
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}
		c := b[0]
		switch c {
		case '-':
			labelHash = currHash
			currHash = addToHash(c, currHash)
			boxes = removeFromBoxes(boxes, labelHash, currLabel)
		case '=':
			labelHash = currHash
			currHash = addToHash(c, currHash)
			inFocalLen = true
		case ',', '\n':
			part1Total += currHash
			if inFocalLen {
				boxes = addToBoxes(boxes, labelHash, currLabel, currFocalLen)
				inFocalLen = false
				currFocalLen = make([]byte, 0)
			}
			currHash = 0
			currLabel = make([]byte, 0)
		default:
			currHash = addToHash(c, currHash)
			if inFocalLen {
				currFocalLen = append(currFocalLen, c)
			} else {
				currLabel = append(currLabel, c)
			}
		}
	}

	part2Total := 0
	for i, b := range boxes {
		for j, l := range b {
			part2Total += (i + 1) * (j + 1) * l.focalLength
		}
	}

	fmt.Println("part 1:", part1Total)
	fmt.Println("part 2:", part2Total)
}

func addToHash(c byte, x int) int {
	x += int(c)
	x *= 17
	x %= 256
	return x
}

type Lens struct {
	label       string
	focalLength int
}

func addToBoxes(boxes [256][]*Lens, hash int, labelBytes, focalLengthBytes []byte) [256][]*Lens {
	focalLength, _ := strconv.Atoi(string(focalLengthBytes))
	lens := &Lens{string(labelBytes), focalLength}
	boxes[hash] = addToBox(boxes[hash], lens)
	return boxes
}

func addToBox(box []*Lens, lens *Lens) []*Lens {
	for i, l := range box {
		if l.label == lens.label {
			box[i] = lens
			return box
		}
	}
	return append(box, lens)
}

func removeFromBoxes(boxes [256][]*Lens, hash int, labelBytes []byte) [256][]*Lens {
	label := string(labelBytes)
	boxes[hash] = removeFromBox(boxes[hash], label)
	return boxes
}

func removeFromBox(box []*Lens, label string) []*Lens {
	if len(box) == 0 {
		return box
	}
	for i, l := range box {
		if l.label == label {
			if i == 0 {
				if len(box) == 1 {
					return []*Lens{}
				}
				return box[i+1:]
			}
			if i == len(box)-1 {
				return box[:i]
			}
			return append(box[:i], box[i+1:]...)
		}
	}
	return box
}
