package day04

import (
	"adventofcode/src/utils"
	"bufio"
	"fmt"
	"math"
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
	elfSet := make(map[int]struct{})
	for _, num := range card.ElfNumbers {
		elfSet[num] = struct{}{}
	}

	intersectionCount := 0
	for _, num := range card.WinningNumbers {
		if _, exists := elfSet[num]; exists {
			intersectionCount++
		}
	}

	if intersectionCount == 0 {
		return 0
	}

	return int(math.Pow(2, float64(intersectionCount-1)))
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

/** Iteratively cascades through the cards, starting at the given offset, which
 * depends on the card's position in the input list of Cards. Using a stack and while loop
 * proved to be significantly faster than using recursion.
 */
func cascadeAllCards(startIndex int, scratchCards []Card, intersections []int) int {
    stack := []int{startIndex}
    totalCards := 0

    for len(stack) > 0 {
        index := stack[len(stack)-1]
        stack = stack[:len(stack)-1] // New stack of card positions without the card we will process in current iter

        cardMatches := intersections[index]
        if cardMatches == 0 {
            continue
        }

        for i := 0; i < cardMatches; i++ {
            nextCardIndex := index + i + 1
            if nextCardIndex < len(scratchCards) {
                totalCards++
                stack = append(stack, nextCardIndex)
            }
        }
    }

    return totalCards
}

/** Counts the number of cards iteratively by cascading through the cards,
 * pre-computing the number of intersections between each card first.
 */
func countCards(scratchCards []Card) int {
    totalCards := len(scratchCards)

    intersections := make([]int, len(scratchCards))
    for i, card := range scratchCards {
        intersections[i] = len(utils.Intersection(card.ElfNumbers, card.WinningNumbers))
    }

    for i := range scratchCards {
        totalCards += cascadeAllCards(i, scratchCards, intersections)
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
