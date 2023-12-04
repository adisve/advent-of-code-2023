package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var codebook = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func findFirstAndLastNumericPart(s string, includeWords bool) (string, string) {
	first := ""
	last := ""

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			if first == "" {
				first = string(s[i])
			}
			last = string(s[i])
		}

		if includeWords {
			for j := 3; j <= 5; j++ {
				if i+j <= len(s) {
					word := s[i : i+j]
					if num, ok := codebook[word]; ok {
						if first == "" {
							first = num
						}
						last = num
					}
				}
			}
		}
	}

	return first, last
}

func calculateSum(file *bufio.Scanner, includeWords bool) int {
	totalSum := 0

	for file.Scan() {
		calibration := strings.ToLower(file.Text())
		firstNumber, lastNumber := findFirstAndLastNumericPart(calibration, includeWords)

		if firstNumber != "" && lastNumber != "" {
			sum, err := strconv.Atoi(firstNumber + lastNumber)
			if err == nil {
				totalSum += sum
			}
		}
	}

	return totalSum
}

/** Read each row in input from left to find first number,
* when first number is found read from the end to find second number -
* alphabetical numbers are matched against a codebook
**/
func Run() {
	file, err := os.Open("src/inputs/day01.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1Sum := calculateSum(scanner, false)
	fmt.Println("Day 1 - Part 1: ", part1Sum)

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	part2Sum := calculateSum(scanner, true)
	fmt.Println("Day 1 - Part 2: ", part2Sum)
}
