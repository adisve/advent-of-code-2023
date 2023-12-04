package day04

import (
	"adventofcode/src/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/** A Card struct contains the elf numbers and the winning numbers.
 */
type Card struct {
	ElfNumbers     []int
	WinningNumbers []int
}

func readCardsFromFile(path string) ([]Card, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cards []Card
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		card := parseCard(scanner.Text())
		cards = append(cards, card)
	}

	return cards, nil
}

/** Parses a line of the input file into a Card struct.
 */
func parseCard(line string) Card {
	parts := strings.Split(line, ":")
	numberParts := strings.Split(parts[1], "|")

	elfNumbers := parseNumbers(strings.TrimSpace(numberParts[0]))
	winningNumbers := parseNumbers(strings.TrimSpace(numberParts[1]))

	return Card{
		ElfNumbers:     elfNumbers,
		WinningNumbers: winningNumbers,
	}
}

/** Parses a string of numbers separated by spaces into a slice of ints.
 */
func parseNumbers(numberStr string) []int {
	strNums := strings.Fields(numberStr)
	nums := make([]int, len(strNums))
	for i, str := range strNums {
		num, _ := strconv.Atoi(str)
		nums[i] = num
	}
	return nums
}

/** Calculates the value of the card by doubling the value for each matching
 * number between the elf numbers and the winning numbers.
 */
func getCardValue(card Card) int {
	cardValue := 0
	for _, elfNum := range card.ElfNumbers {
		for _, winNum := range card.WinningNumbers {
			if elfNum == winNum {
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue = 2 * cardValue
				}
			}
		}
	}
	return cardValue
}

/** Calculates the sum of the card values.
 */
func getCardSums(scratchCards []Card) int {
	cardSums := 0
	for _, card := range scratchCards {
		cardSums += getCardValue(card)
	}
	return cardSums
}

/** Recursively cascades through the cards, starting at the given offset which
 * depends on the card's position in the input list of Cards.
 */
func cascade(card Card, offset int, scratchCards []Card) int {
	cardMatches := len(utils.Intersection(card.ElfNumbers, card.WinningNumbers))
	if cardMatches == 0 {
		return 0
	}

	totalCopies := 0
	for i := 0; i < cardMatches; i++ {
		nextCardIndex := offset + i + 1
		if nextCardIndex < len(scratchCards) {
			nextCard := scratchCards[nextCardIndex]
			totalCopies += 1 + cascade(nextCard, nextCardIndex, scratchCards)
		}
	}
	return totalCopies
}

/** Counts the number of cards recursively by cascading through the cards.
 */
func countCards(scratchCards []Card) int {
	totalCards := 0

	for offset, card := range scratchCards {
		totalCards += 1
		totalCards += cascade(card, offset, scratchCards)
	}

	return totalCards
}

func Run() {
	cards, err := readCardsFromFile("src/inputs/day04.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Day 4 - Part 1:", getCardSums(cards))
	fmt.Println("Day 4 - Part 2:", countCards(cards))
}
