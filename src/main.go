package main

import (
	"adventofcode/src/solutions/day01"
	"adventofcode/src/solutions/day02"
	"adventofcode/src/solutions/day03"
	"adventofcode/src/solutions/day04"
	"adventofcode/src/solutions/day05"
	"adventofcode/src/solutions/day06"
	"adventofcode/src/solutions/day07"
	"adventofcode/src/solutions/day08"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Choose solution to display >> ")

	for scanner.Scan() {

		solution := scanner.Text()

		switch solution {
		case "1":
			day01.Run()
		case "2":
			day02.Run()
		case "3":
			day03.Run()
		case "4":
			day04.Run()
		case "5":
			day05.Run()
		case "6":
			day06.Run()
		case "7":
			day07.Run()
		case "8":
			day08.Run()
		case "q", "exit", "quit":
			os.Exit(0)
		default:
			fmt.Println("Invalid solution number.")
		}

		fmt.Print("\nChoose solution to display >> ")
	}
}
