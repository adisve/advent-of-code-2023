package day08

import (
	"adventofcode/src/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Reads the file and returns the directions, adjacency list, and the first node key
func readFile(path string) ([]rune, map[string][2]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	directions := []rune(scanner.Text())

	adjList := make(map[string][2]string)
	firstNodeKey := ""
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		if firstNodeKey == "" {
			firstNodeKey = key
		}

		values := strings.Trim(parts[1], "()")
		nodes := strings.Split(values, ", ")
		if len(nodes) != 2 {
			continue
		}
		adjList[key] = [2]string{nodes[0], nodes[1]}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return directions, adjList, nil
}

func getSteps(directions []rune, graph *utils.Graph, startNode string) int {
	currentNode := startNode
	steps := 0

	for !strings.HasSuffix(currentNode, "Z") {
		direction := string(directions[steps%len(directions)])
		currentNode = graph.GetNextNode(currentNode, direction).Val
		steps++
	}

	return steps
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func solveFirst(directions []rune, graph *utils.Graph) int {
	currentNodeKey := "AAA"
	count := 0

	for i := 0; currentNodeKey != "ZZZ"; i = (i + 1) % len(directions) {
		direction := string(directions[i])
		currentNodeKey = graph.GetNextNode(currentNodeKey, direction).Val
		count++
	}

	return count
}

func solveSecond(directions []rune, graph *utils.Graph) int {
	startNodes := graph.GetStartingNodeKeys()
	var steps []int

	for _, node := range startNodes {
		steps = append(steps, getSteps(directions, graph, node))
	}

	result := steps[0]
	for _, step := range steps[1:] {
		result = lcm(result, step)
	}

	return result
}

func Run() {
	filePath := "src/inputs/day08.txt"
	directions, adjacencyList, _ := readFile(filePath)
	graph := utils.NewGraph(adjacencyList)

	fmt.Printf("Day 8 - Task 1: %d\n", solveFirst(directions, graph))
	fmt.Printf("Day 8 - Task 2: %d\n", solveSecond(directions, graph))
}
