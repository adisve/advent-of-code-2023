package day06

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func readInput(path string) ([]int, []int) {
    content, _ := os.ReadFile(path)
    lines := strings.Split(string(content), "\n")

    toIntSlice := func(s []string) []int {
        intSlice := make([]int, len(s))
        for i, v := range s {
            intSlice[i], _ = strconv.Atoi(v)
        }
        return intSlice
    }

    times := toIntSlice(strings.Fields(strings.TrimPrefix(lines[0], "Time: ")))
    distances := toIntSlice(strings.Fields(strings.TrimPrefix(lines[1], "Distance: ")))

    return times, distances
}


func calculateRaces(times []int, distances []int) int {
    product := 1
    for i, time := range times {
        makeProduct(time, distances[i], &product)
    }

    return product
}

func calculateRace(times []int, distances []int) int {
	raceTime := toInt(times)
	totalDistance := toInt(distances)
	product := 1

	makeProduct(raceTime, totalDistance, &product)

	return product
}

/** The calculation for the distance is essentially a quadratic,
 * so the key is to get the button hold times before and after 
 * the vertex where your possible distance in the race exceeds
 * the specific record distance.
 *
 * worst-case runtime complexity of the function is O(m x n)
*/
func makeProduct(time int, distance int, product *int) {
	vertex := time / 2

	lowerBound := 0
	for n := 0; n <= vertex; n++ {
		if n*(time-n) > distance {
			lowerBound = n
			break
		}
	}

	upperBound := time
	for n := time - 1; n >= vertex; n-- {
		if n*(time-n) > distance {
			upperBound = n
			break
		}
	}

	winnings := upperBound - lowerBound + 1
	if lowerBound > upperBound {
		winnings = 0
	}

	*product *= winnings
}


func toInt(list []int) int {
	var resultStr string
	for _, num := range list {
		resultStr += strconv.Itoa(num)
	}
	result, _ := strconv.Atoi(resultStr)
	return result
}


func Run() {
	times, distances := readInput("src/inputs/day06.txt")
	fmt.Println("Day 6 - Task 1: ", calculateRaces(times, distances))
	fmt.Println("Day 6 - Part 2: ", calculateRace(times, distances))
}