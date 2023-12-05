package day05

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type RangeMap struct {
	SourceStart      int
	SourceEnd        int
	DestinationStart int
}

func parseRangeMap(s []string) (RangeMap, error) {
	if len(s) < 3 {
		return RangeMap{}, fmt.Errorf("invalid range map")
	}
	sourceStart, err := strconv.Atoi(s[1])
	if err != nil {
		return RangeMap{}, err
	}
	rangeLen, err := strconv.Atoi(s[2])
	if err != nil {
		return RangeMap{}, err
	}
	destinationStart, err := strconv.Atoi(s[0])
	if err != nil {
		return RangeMap{}, err
	}
	return RangeMap{
		SourceStart:      sourceStart,
		SourceEnd:        sourceStart + rangeLen,
		DestinationStart: destinationStart,
	}, nil
}

func transformValue(val int, maps []RangeMap) int {
	for _, m := range maps {
		if val >= m.SourceStart && val < m.SourceEnd {
			return m.DestinationStart + (val - m.SourceStart)
		}
	}
	return val
}

func processSeeds(path string, rangeMapSets [][][]string, isRange bool) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	lowest := math.MaxInt32
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds: ") {
			seedStrs := strings.Fields(strings.TrimPrefix(line, "seeds: "))
			for i := 0; i < len(seedStrs); i++ {
				var seed, upperBound int

				if isRange {
					seed, _ = strconv.Atoi(seedStrs[i])
					seedRange, _ := strconv.Atoi(seedStrs[i+1])
					upperBound = seedRange + seed
					i++
				} else {
					seed, _ = strconv.Atoi(seedStrs[i])
					upperBound = seed + 1
				}

				for seed < upperBound {
					transformedSeed := seed
					for _, rangeMaps := range rangeMapSets {
						var maps []RangeMap
						for _, rangeMap := range rangeMaps {
							rm, err := parseRangeMap(rangeMap)
							if err != nil {
								return 0, fmt.Errorf("error parsing range map: %w", err)
							}
							maps = append(maps, rm)
						}
						transformedSeed = transformValue(transformedSeed, maps)
					}

					if transformedSeed < lowest {
						lowest = transformedSeed
					}
					seed++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lowest, nil
}

func readRangeMaps(path string) ([][][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var maps [][][]string
	var currentMap [][]string
	scanner := bufio.NewScanner(file)
	isInMapSection := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "map:") {
			isInMapSection = true
			if currentMap != nil {
				maps = append(maps, currentMap)
				currentMap = nil
			}
		} else if line != "" && isInMapSection {
			if !strings.HasPrefix(line, "seeds:") {
				values := strings.Fields(line)
				currentMap = append(currentMap, values)
			}
		}
	}
	if currentMap != nil {
		maps = append(maps, currentMap)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return maps, nil
}

func Run() {
	dayPath := "src/inputs/day05.txt"
	rangeMapSets, err := readRangeMaps(dayPath)
	if err != nil {
		fmt.Println("Error reading range maps:", err)
		return
	}

	lowestTransformedSeedOne, err := processSeeds(dayPath, rangeMapSets, false)
	if err != nil {
		fmt.Println("Error processing seeds:", err)
		return
	}
	fmt.Println("Day 5 - Part 1:", lowestTransformedSeedOne)

	lowestTransformedSeedTwo, err := processSeeds(dayPath, rangeMapSets, true)
	if err != nil {
		fmt.Println("Error processing seeds as range:", err)
		return
	}
	fmt.Println("Day 5 - Part 2:", lowestTransformedSeedTwo)
}
