package utils

import (
	"fmt"
	"strings"
)

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph(adjList map[string][2]string) *Graph {
	graph := &Graph{Nodes: make(map[string]*Node)}

	for key := range adjList {
		graph.Nodes[key] = &Node{Val: key}
	}

	for key, edges := range adjList {
		graph.Nodes[key].Left = graph.Nodes[edges[0]]
		graph.Nodes[key].Right = graph.Nodes[edges[1]]
	}

	return graph
}

func (g *Graph) GetNextNode(currentNodeKey, direction string) *Node {
	currentNode := g.Nodes[currentNodeKey]
	if direction == "L" {
		return currentNode.Left
	} else {
		return currentNode.Right
	}
}

func (g *Graph) GetStartingNodeKeys() []string {
	keys := make([]string, 0, len(g.Nodes))
	for key := range g.Nodes {
		if strings.HasSuffix(key, "A") {
			keys = append(keys, key)
		}
	}
	return keys
}

func (g *Graph) PrintGraph() {
	for key, node := range g.Nodes {
		left := ""
		right := ""

		if node.Left != nil {
			left = node.Left.Val
		}
		if node.Right != nil {
			right = node.Right.Val
		}

		fmt.Printf("%s = (%s, %s)\n", key, left, right)
	}
}
