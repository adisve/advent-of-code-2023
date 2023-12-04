package day02

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

/** Opens a file at the given path and reads its content line by line.
* Each line is assumed to be a game, and the function returns a slice of these games as strings.
**/
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

/** Takes a string representing a game and splits it into turns.
* Each turn is processed to extract counts of colors and returns a slice of maps,
* where each map represents the counts of colors in a turn.
**/
func parseGame(gameStr string) []map[string]int {
	parts := strings.Split(gameStr, ":")
	turns := strings.Split(parts[len(parts)-1], ";")
	var turnCounts []map[string]int
	for _, turn := range turns {
		turnCounts = append(turnCounts, processTurn(turn))
	}
	fmt.Println(turnCounts)
	return turnCounts
}

/** Takes a string representing a turn and splits it into cases of colors.
* Parse the counts of each color and updates the counts if the new count is higher than the existing one.
**/
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

/** Iterate over each turn's color count map.
* Check if any color count exceeds the predefined limit, returning false if so.
**/
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

/** Calculate the product of the highest count of each color across all turns.
**/
func getLeastAmount(turnCounts []map[string]int) int {
	mins := map[string]int{"red": 0, "green": 0, "blue": 0}
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

/** Reads games from a file, processes each game,
* and calculates the sum of valid game IDs and the sum of
* least amounts of colors across all games.
**/
func Run() {
	games, err := readGamesFromFile("src/inputs/day02.txt")
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

	fmt.Println("Day 2 - Part 1: ", validGamesSum)
	fmt.Println("Day 2 - Part 2: ", powerSetsSum)
}
