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

type SeedRange struct {
	Start int
	End   int
}

func parseRangeMap(s []string) RangeMap {
	sourceStart, _ := strconv.Atoi(s[1])
	rangeLen, _ := strconv.Atoi(s[2])
	destinationStart, _ := strconv.Atoi(s[0])
	return RangeMap{
		SourceStart:      sourceStart,
		SourceEnd:        sourceStart + rangeLen,
		DestinationStart: destinationStart,
	}
}

func transformValue(val int, maps []RangeMap) int {
	for _, m := range maps {
		if val >= m.SourceStart && val < m.SourceEnd {
			return m.DestinationStart + (val - m.SourceStart)
		}
	}
	return val
}

func readSeedsAsRange(path string) []SeedRange {
	file, _ := os.Open(path)
	defer file.Close()

	var seeds []SeedRange
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds: ") {
			seedStrs := strings.Fields(strings.TrimPrefix(line, "seeds: "))
			for i := 0; i < len(seedStrs); i += 2 {
				seed, _ := strconv.Atoi(seedStrs[i])
				count, _ := strconv.Atoi(seedStrs[i+1])
				upperBound := count + seed
				seeds = append(seeds, SeedRange{ seed, upperBound })
			}
		}
	}
	return seeds
}

func readSeeds(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	var seeds []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds: ") {
			seedStrs := strings.Fields(strings.TrimPrefix(line, "seeds: "))
			for _, s := range seedStrs {
				seed, _ := strconv.Atoi(s)
				seeds = append(seeds, seed)
			}
		}
	}
	return seeds
}

func readRangeMaps(path string) [][][]string {
	file, _ := os.Open(path)
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
			values := strings.Fields(line)
			currentMap = append(currentMap, values)
		}
	}
	if currentMap != nil {
		maps = append(maps, currentMap)
	}
	return maps
}

func findLowestTransformedSeedInRange(seedRanges []SeedRange, rangeMapSets [][][]string) int {
	lowest := math.MaxInt32
	for _, seedRange := range seedRanges {
		for seed := seedRange.Start; seed < seedRange.End; seed++ {
			lowest = transformSeed(seed, lowest, rangeMapSets)
		}
	}
	return lowest
}

func findLowestTransformedSeed(seeds []int, rangeMapSets [][][]string) int {
	lowest := math.MaxInt32
	for _, seed := range seeds {
		lowest = transformSeed(seed, lowest, rangeMapSets)
	}
	return lowest
}

func transformSeed(seed int, lowest int, rangeMapSets [][][]string) int  {
	transformedSeed := seed
	for _, rangeMaps := range rangeMapSets {
		var maps []RangeMap
		for _, rangeMap := range rangeMaps { // Read all the maps related to a specific transformation of a seed based on position
			rm := parseRangeMap(rangeMap)
			maps = append(maps, rm)
		}
		transformedSeed = transformValue(transformedSeed, maps)
	}
	if transformedSeed < lowest {
		return transformedSeed
	}
	return lowest
}

func Run() {
	dayPath := "src/inputs/day05.txt"

	seeds := readSeeds(dayPath)
	rangeMapSets := readRangeMaps(dayPath)

	lowestTransformedSeedOne := findLowestTransformedSeed(seeds, rangeMapSets)
	fmt.Println("Day 5 - Part 1:", lowestTransformedSeedOne)

	seedsAsRange := readSeedsAsRange(dayPath)
	lowestTransformedSeedTwo := findLowestTransformedSeedInRange(seedsAsRange, rangeMapSets)

	fmt.Println("Day 5 - Part 2:", lowestTransformedSeedTwo)
}