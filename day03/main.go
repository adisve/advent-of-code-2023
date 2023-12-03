package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isAdjacentToSymbol(symbolMatrix [][]rune, startX, endX, y int) (bool, [2]int) {
	rows := len(symbolMatrix)
	cols := len(symbolMatrix[0])

	for dx := startX; dx <= endX; dx++ {
		for _, dy := range []int{-1, 1} {
			if newY := y + dy; 0 <= newY && newY < rows && isSymbol(symbolMatrix[newY][dx]) {
				return true, [2]int{newY, dx}
			}
		}
	}

	for dy := -1; dy <= 1; dy++ {
		if newY := y + dy; 0 <= newY && newY < rows {
			if startX > 0 && isSymbol(symbolMatrix[newY][startX-1]) {
				return true, [2]int{newY, startX - 1}
			}
			if endX < cols-1 && isSymbol(symbolMatrix[newY][endX+1]) {
				return true, [2]int{newY, endX + 1}
			}
		}
	}

	return false, [2]int{}
}

func isSymbol(char rune) bool {
	return !isNumeric(char) && char != '.'
}

func isNumeric(char rune) bool {
	return '0' <= char && char <= '9'
}

func calculateGearRatios(numberIndices [][3]int, schematic [][]rune) int {
	adjacencies := make(map[[2]int][]int)
	for _, indices := range numberIndices {
		startX, endX, y := indices[0], indices[1], indices[2]
		adjacent, symbolPos := isAdjacentToSymbol(schematic, startX, endX, y)
		if adjacent && schematic[symbolPos[0]][symbolPos[1]] == '*' {
			number, _ := strconv.Atoi(string(schematic[y][startX : endX+1]))
			adjacencies[symbolPos] = append(adjacencies[symbolPos], number)
		}
	}

	sum := 0
	for _, values := range adjacencies {
		if len(values) > 1 {
			product := 1
			for _, val := range values {
				product *= val
			}
			sum += product
		}
	}
	return sum
}

func calculateSumParts(numberIndices [][3]int, schematic [][]rune) int {
	sum := 0
	for _, indices := range numberIndices {
		startX, endX, y := indices[0], indices[1], indices[2]
		if adjacent, _ := isAdjacentToSymbol(schematic, startX, endX, y); adjacent {
			number, _ := strconv.Atoi(string(schematic[y][startX : endX+1]))
			sum += number
		}
	}
	return sum
}

func getNumIndices(schematic [][]rune) [][3]int {
	var numberIndices [][3]int
	for y, row := range schematic {
		startIndex := -1
		for x, char := range row {
			if isNumeric(char) {
				if startIndex == -1 {
					startIndex = x
				}
				if x == len(row)-1 || !isNumeric(row[x+1]) {
					endIndex := x
					numberIndices = append(numberIndices, [3]int{startIndex, endIndex, y})
					startIndex = -1
				}
			} else {
				startIndex = -1
			}
		}
	}
	return numberIndices
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var schematic [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
	}

	numberIndices := getNumIndices(schematic)

	sumParts := calculateSumParts(numberIndices, schematic)
	gearRatios := calculateGearRatios(numberIndices, schematic)

	fmt.Println("Day 3 - Part 1:", sumParts)
	fmt.Println("Day 3 - Part 2:", gearRatios)
}
