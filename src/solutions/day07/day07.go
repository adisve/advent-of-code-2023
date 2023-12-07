package day07

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
	rank  int
}

func parseHand(handStr string) Hand {
	parts := strings.Fields(handStr)
	bid, _ := strconv.Atoi(parts[1])
	return Hand{cards: parts[0], bid: bid}
}

func classify(hand string) int {
	counts := make(map[rune]int)
	for _, card := range hand {
		counts[card]++
	}

	if contains(counts, 5) {
		return 6
	}
	if contains(counts, 4) {
		return 5
	}
	if contains(counts, 3) {
		if contains(counts, 2) {
			return 4
		}
		return 3
	}
	if countPairs(counts) == 2 {
		return 2
	}
	if contains(counts, 2) {
		return 1
	}
	return 0
}

func contains(m map[rune]int, value int) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func countPairs(m map[rune]int) int {
	pairs := 0
	for _, v := range m {
		if v == 2 {
			pairs++
		}
	}
	return pairs
}

func mapCards(hand string) string {
	mapped := ""
	for _, card := range hand {
		switch card {
		case 'T':
			mapped += "A"
		case 'J':
			mapped += "B"
		case 'Q':
			mapped += "C"
		case 'K':
			mapped += "D"
		case 'A':
			mapped += "E"
		default:
			mapped += string(card)
		}
	}
	return mapped
}

func strength(hand string) (int, string) {
	return classify(hand), mapCards(hand)
}

func readInput(path string) []string {
	content, _ := os.ReadFile(path)
	lines := strings.Split(string(content), "\n")

	handsAndBids := make([]string, 0)

	for _, line := range lines {
		handsAndBids = append(handsAndBids, line)
	}

	return handsAndBids
}

func solve(handsAndBids []string) int {
	hands := make([]Hand, len(handsAndBids))
	for i, handStr := range handsAndBids {
		hands[i] = parseHand(handStr)
	}

	sort.Slice(hands, func(i, j int) bool {
		classI, mappedI := strength(hands[i].cards)
		classJ, mappedJ := strength(hands[j].cards)

		if classI == classJ {
			return mappedI < mappedJ
		}
		return classI < classJ
	})

	totalWinnings := 0
	for i, hand := range hands {
		rank := i + 1
		totalWinnings += rank * hand.bid
	}

	return totalWinnings
}

func Run() {
	handsAndBids := readInput(("src/inputs/day07.txt"))
	winnings := solve(handsAndBids)

	fmt.Printf("Day 7 - Task 1: %d\n", winnings)
}
