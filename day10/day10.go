package day10

import (
	"fmt"
	"strings"

	"github.com/sidkhuntia/advent-of-code-2023/utils"
)

var PIPES = map[string][][]int{
	"-": {
		{0, 0},
		{1, 0},
		{0, 0},
		{-1, 0},
	},
	"|": {
		{0, -1},
		{0, 0},
		{0, 1},
		{0, 0},
	},
	"7": {
		{-1, 0},
		{0, 1},
		{0, 0},
		{0, 0},
	},
	"J": {
		{0, 0},
		{0, -1},
		{-1, 0},
		{0, 0},
	},
	"L": {
		{0, 0},
		{0, 0},
		{1, 0},
		{0, -1},
	},
	"F": {
		{1, 0},
		{0, 0},
		{0, 0},
		{0, 1},
	},
}

func detectDirection(targetX, targetY, startX, startY int) int {
	if targetY < startY {
		return 0
	}
	if targetX > startX {
		return 1
	}
	if targetY > startY {
		return 2
	}
	if targetX < startX {
		return 3
	}
	return -1
}

func checkNewCoordinates(grid []string, newX, newY int) bool {
	if newX < 0 || newY < 0 {
		return false
	}
	if newY >= len(grid) || newX >= len(grid[0]) {
		return false
	}
	if string(grid[newY][newX]) == "." {
		return false
	}
	return true
}

func dfs(grid []string, currX, currY, direction, currSteps int, stepsGrid [][]int) {
	char := string(grid[currY][currX])

	if char == "S" {
		return
	}

	if stepsGrid[currY][currX] != 0 {
		stepsGrid[currY][currX] = utils.MinInt(stepsGrid[currY][currX], currSteps)
	} else {
		stepsGrid[currY][currX] = currSteps
	}

	newX := currX + PIPES[char][direction][0]
	newY := currY + PIPES[char][direction][1]

	if !checkNewCoordinates(grid, newX, newY) {
		return
	}

	if newX == currX && newY == currY {
		return
	}

	newDirection := detectDirection(newX, newY, currX, currY)

	dfs(grid, newX, newY, newDirection, currSteps+1, stepsGrid)

}

func Part01(input string) {
	lines := utils.SplitLines(input)

	startX := -1
	startY := -1

	for i, line := range lines {
		for j, char := range strings.Split(line, "") {
			if char == "S" {
				startX = j
				startY = i
				break
			}
		}
		if startX != -1 || startY != -1 {
			break
		}
	}

	ans := 0

	var dir = []int{0, 1, 0, -1, 0}
	stepsGrid := make([][]int, len(lines))

	for i := range lines {
		stepsGrid[i] = make([]int, len(lines[0]))
	}

	for i := 0; i < 4; i++ {
		newX := startX + dir[i]
		newY := startY + dir[i+1]

		if !checkNewCoordinates(lines, newX, newY) {
			continue
		}

		newDirection := detectDirection(newX, newY, startX, startY)

		dfs(lines, newX, newY, newDirection, 1, stepsGrid)

	}

	for _, line := range stepsGrid {
		for _, val := range line {
			ans = utils.Max(ans, val)
		}
	}

	fmt.Printf("Answer: %d\n", ans)

}

func Part02(input string) {

}
