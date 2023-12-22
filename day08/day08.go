package day08

import (
	"fmt"
	"strings"

	"github.com/sidkhuntia/advent-of-code-2023/utils"
)

type MapNode struct {
	left  string
	right string
}

func Part01(input string) {
	lines := utils.SplitLines(input)
	walkingPath := lines[0]
	lines = lines[1:]

	graph := make(map[string]MapNode)

	for _, line := range lines {
		currNode, options, _ := strings.Cut(line, " = ")
		options = options[1 : len(options)-1]
		left, right, _ := strings.Cut(options, ", ")

		graph[currNode] = MapNode{left, right}

	}

	currNode := "AAA"
	targetNode := "ZZZ"

	steps := 0

	for currNode != targetNode {
		currStep := walkingPath[(steps % len(walkingPath))]

		switch currStep {
		case 'L':
			currNode = graph[currNode].left
		case 'R':
			currNode = graph[currNode].right
		}
		steps++
	}

	fmt.Printf("Steps: %d\n", steps)

}

func Part02(input string) {
	lines := utils.SplitLines(input)
	walkingPath := lines[0]
	lines = lines[1:]

	graph := make(map[string]MapNode)

	startingNodes := make([]string, 0)
	endingNodes := make([]string, 0)

	for _, line := range lines {
		currNode, options, _ := strings.Cut(line, " = ")
		options = options[1 : len(options)-1]
		left, right, _ := strings.Cut(options, ", ")

		if currNode[2] == 'A' {
			startingNodes = append(startingNodes, currNode)
		}
		if currNode[2] == 'Z' {
			endingNodes = append(endingNodes, currNode)
		}

		graph[currNode] = MapNode{left, right}

	}

	// fmt.Printf("Length of Starting Nodes: %v\n", len(startingNodes))
	// fmt.Printf("Length of Ending Nodes: %v\n", len(endingNodes))

	minStepsToReachEnd := make([]int64, 0)

	for _, currNode := range startingNodes {
		steps := 0
		for currNode[2] != 'Z' {
			currStep := walkingPath[(steps % len(walkingPath))]

			// fmt.Printf("Prev Node: %s\n", startingNodes[nodeIdx])
			switch currStep {
			case 'L':
				currNode = graph[currNode].left
			case 'R':
				currNode = graph[currNode].right
			}
			steps++

			// fmt.Printf("Updated Node: %s\n", startingNodes[nodeIdx])
		}
		minStepsToReachEnd = append(minStepsToReachEnd, int64(steps))
	}

	ans := minStepsToReachEnd[0]

	for i := 1; i < len(minStepsToReachEnd); i++ {
		gcd := utils.GCD64(ans, minStepsToReachEnd[i])
		ans = (ans * minStepsToReachEnd[i]) / gcd
	}

	fmt.Printf("Steps: %v\n", ans)

}
