package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var m = map[byte]byte{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func main() {
	file, _ := os.Open("../inputs/day7.txt")
	defer file.Close()

	part1_total := 0
	part2_total := 0

	scanner := bufio.NewScanner(file)
	hands := make([]*Hand, 0)
	for scanner.Scan() {
		hands = append(hands, lineToHand(scanner.Text()))
	}

	sort.Slice(hands, makeSortFn(hands))

	for i, h := range hands {
		part1_total += (i + 1) * h.Bet
	}

	sort.Slice(hands, makeSortFn2(hands))

	for i, h := range hands {
		part2_total += (i + 1) * h.Bet
	}

	fmt.Println("part 1:", part1_total)
	fmt.Println("part 2:", part2_total)
}

type Hand struct {
	Value  [7]byte
	Value2 [7]byte
	Bet    int
}

func lineToHand(line string) *Hand {
	splitLine := strings.Split(line, " ")
	bet, _ := strconv.Atoi(splitLine[1])
	value, value2 := stringToValue(splitLine[0])
	return &Hand{value, value2, bet}
}

func stringToValue(hand string) ([7]byte, [7]byte) {
	valueMap := make(map[byte]byte, 5)
	for i := 0; i < len(hand); i++ {
		x := m[hand[i]]
		if _, ok := valueMap[x]; ok {
			valueMap[x]++
		} else {
			valueMap[x] = 1
		}
	}
	return value1(valueMap, hand), value2(valueMap, hand)
}

func value1(valueMap map[byte]byte, hand string) [7]byte {
	var maxCard, maxCount byte
	for c, v := range valueMap {
		if v > maxCount {
			maxCard, maxCount = c, v
		}
	}
	var secCount byte
	for c, v := range valueMap {
		if v > secCount && c != maxCard {
			secCount = v
		}
	}
	return [7]byte{maxCount, secCount,
		m[hand[0]], m[hand[1]], m[hand[2]], m[hand[3]], m[hand[4]]}

}

func value2(valueMap map[byte]byte, hand string) [7]byte {
	hand = strings.ReplaceAll(hand, "J", "1")
	jCount := valueMap[11]
	var maxCard, maxCount byte
	for c, v := range valueMap {
		if v > maxCount && c != 11 {
			maxCard, maxCount = c, v
		}
	}
	var secCount byte
	for c, v := range valueMap {
		if v > secCount && c != maxCard && c != 11 {
			secCount = v
		}
	}
	return [7]byte{maxCount + jCount, secCount,
		m[hand[0]], m[hand[1]], m[hand[2]], m[hand[3]], m[hand[4]]}
}

func makeSortFn(hands []*Hand) func(int, int) bool {
	return func(i, j int) bool {
		for x := 0; x < 7; x++ {
			diff := int(hands[i].Value[x]) - int(hands[j].Value[x])
			if diff != 0 {
				return diff < 0
			}
		}
		return true
	}
}

func makeSortFn2(hands []*Hand) func(int, int) bool {
	return func(i, j int) bool {
		for x := 0; x < 7; x++ {
			diff := int(hands[i].Value2[x]) - int(hands[j].Value2[x])
			if diff != 0 {
				return diff < 0
			}
		}
		return true
	}
}
