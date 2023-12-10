package main

import (
	"AoC2023/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed big_input.txt
var input string

type node struct {
	id    string
	left  *node
	right *node
}

func (n *node) String() string {
	return fmt.Sprintf("%s:(%s, %s)", n.id, n.left.id, n.right.id)
}

// Determines whether the node is initialised fully (has defined id, left, and right nodes)
func (n *node) isInitialised() bool {
	return n.id != "" && n.left != nil && n.right != nil
}

type nodeExistsError struct {
	n *node
}

func (e *nodeExistsError) Error() string {
	return fmt.Sprintf("node already exists:\n%#v", e.n)
}

type graph struct {
	nodes      map[string]*node
	startNodes []*node // Nodes that end in A
	endNodes   []*node // Nodes that end in Z
}

// NewGraph creates and returns a new graph with initialized nodes map.
func newGraph() *graph {
	return &graph{
		nodes: make(map[string]*node),
	}
}

func (g *graph) String() string {
	var sb strings.Builder
	for _, n := range g.nodes {
		sb.WriteString(n.String() + "\n")
	}
	return sb.String()
}

// Adds a new node to the graph.
// Returns a *nodeExistsError if a node with the same ID already exists in the graph.
// Otherwise, it adds the node and returns nil.
func (g *graph) addNode(id, leftID, rightID string) error {
	// Does the node exist already?
	newNode, okID := g.nodes[id]
	if okID {
		if newNode.isInitialised() {
			return &nodeExistsError{newNode}
		}
	} else {
		newNode = &node{id: id}
		g.nodes[id] = newNode
		switch id[2] {
		case 'A':
			g.startNodes = append(g.startNodes, newNode)

		case 'Z':
			g.endNodes = append(g.endNodes, newNode)

		}
	}

	nLeft, okLeft := g.nodes[leftID]
	if !okLeft {
		nLeft = &node{id: leftID}
		g.nodes[leftID] = nLeft
	}
	newNode.left = nLeft

	nRight, okRight := g.nodes[rightID]
	if !okRight {
		nRight = &node{id: rightID}
		g.nodes[rightID] = nRight
	}
	newNode.right = nRight

	return nil
}

// Given an input line extract and return the three relevant strings
func parseLine(line string) (string, string, string) {
	id := line[:3]
	leftID := line[7:10]
	rightID := line[12:15]

	return id, leftID, rightID
}

// In this version of the problem there are multiple starting nodes (all nodes ending in A instead of just AAA)
// We must simultaneously follow the paths from each starting node and find the timestep at which ALL nodes end in Z
func main() {
	lines := utils.SplitLines(input)
	instructions := lines[0] // RL list that loops
	g := newGraph()

	// Populate graph
	for _, line := range lines[2:] {
		err := g.addNode(parseLine(line))
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(instructions)

	currentNodes := make([]*node, len(g.startNodes))
	copy(currentNodes, g.startNodes)

	// Print the starting nodes
	for _, n := range currentNodes {
		fmt.Printf("%s\t", n.id)
	}
	fmt.Println()

	allCurrentNodesEnd := func(cn []*node) bool {
		for _, n := range cn {
			if n.id[2] != 'Z' {
				return false
			}
		}
		return true
	}
	//printStep := func(move string) {
	//	//for _ = range currentNodes {
	//	//	fmt.Printf(" |\t")
	//	//}
	//	//fmt.Println()
	//	for _ = range currentNodes {
	//		fmt.Printf(" %s\t", move)
	//	}
	//	//fmt.Println()
	//	//for _ = range currentNodes {
	//	//	fmt.Print(" v\t")
	//	//}
	//	fmt.Println()
	//	for _, n := range currentNodes {
	//		fmt.Printf("%s\t", n.id)
	//	}
	//	fmt.Println()
	//}

	for i := 0; ; i++ {
		if allCurrentNodesEnd(currentNodes) {
			fmt.Println("\n", i)
			break
		}
		move := instructions[i%len(instructions)]
		switch move {
		case 'L':
			for idx := range currentNodes {
				currentNodes[idx] = currentNodes[idx].left
			}
		case 'R':
			for idx := range currentNodes {
				currentNodes[idx] = currentNodes[idx].right
			}
		}

		//printStep(string(move))
	}

	//// Walk graph
	//currentNode := g.nodes["AAA"]
	//fmt.Print(currentNode.id)
	//for i := 0; ; i++ {
	//	if currentNode.id == "ZZZ" {
	//		fmt.Println("\n", i)
	//		break
	//	}
	//	move := instructions[i%len(instructions)]
	//	switch move {
	//	case 'L':
	//		currentNode = currentNode.left
	//	case 'R':
	//		currentNode = currentNode.right
	//	}
	//	fmt.Printf("-%s>%s", string(move), currentNode.id)
	//}
}
