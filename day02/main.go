package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var caseLimits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func readGamesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var games []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		games = append(games, scanner.Text())
	}
	return games, scanner.Err()
}

func parseGame(gameStr string) []map[string]int {
	parts := strings.Split(gameStr, ":")
	turns := strings.Split(parts[len(parts)-1], ";")
	var turnCounts []map[string]int
	for _, turn := range turns {
		turnCounts = append(turnCounts, processTurn(turn))
	}
	return turnCounts
}

func processTurn(turnStr string) map[string]int {
	counts := map[string]int{"red": 0, "green": 0, "blue": 0}
	cases := strings.Split(turnStr, ",")
	for _, c := range cases {
		parts := strings.Fields(c)
		if len(parts) == 2 {
			amount, err := strconv.Atoi(parts[0])
			color := parts[1]
			if err == nil && amount > counts[color] {
				counts[color] = amount
			}
		}
	}
	return counts
}

func isGameValid(turnCounts []map[string]int) bool {
	for _, turn := range turnCounts {
		for color, amount := range turn {
			if amount > caseLimits[color] {
				return false
			}
		}
	}
	return true
}

func getLeastAmount(turnCounts []map[string]int) int {
	mins := map[string]int{ "red": 0, "green": 0, "blue": 0 }
	for _, turn := range turnCounts {
		for color, amount := range turn {
			if amount > mins[color] {
				mins[color] = amount
			}
		}
	}
	product := 1
	for _, amount := range mins {
		product *= amount
	}
	return product
}

func main() {
	games, err := readGamesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	validGamesSum := 0
	powerSetsSum := 0
	for gameID, gameStr := range games {
		turnCounts := parseGame(gameStr)
		if isGameValid(turnCounts) {
			validGamesSum += gameID + 1
		}
		powerSetsSum += getLeastAmount(turnCounts)
	}

	fmt.Println("Sum of valid games:", validGamesSum)
	fmt.Println("Sum of the power of sets of least amount of cubes:", powerSetsSum)
}