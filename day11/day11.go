package day11

import (
	"fmt"

	"github.com/sidkhuntia/advent-of-code-2023/utils"
)

type Point struct {
	x int
	y int
}

func lowerBound(arr []int, target int) int {
	low := 0
	high := len(arr)

	for low < high {
		mid := (low + high) / 2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func upperBound(arr []int, target int) int {
	low := 0
	high := len(arr)

	for low < high {
		mid := (low + high) / 2
		if arr[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low - 1
}

func Part01(input string) {
	lines := utils.SplitLines(input)

	var emptyPointsX []int
	var emptyPointsY []int

	rowSize := len(lines)
	colSize := len(lines[0])

	filledPointsY := make([]int, rowSize)
	filledPointsX := make([]int, colSize)

	var starPoints []Point

	for rowNo, line := range lines {
		for colNo, char := range line {
			if char == '#' {
				filledPointsY[rowNo]++
				filledPointsX[colNo]++
				starPoints = append(starPoints, Point{colNo, rowNo})
			}
		}
	}

	for i := 0; i < rowSize; i++ {
		if filledPointsY[i] == 0 {
			emptyPointsY = append(emptyPointsY, i)
		}
	}

	for i := 0; i < colSize; i++ {
		if filledPointsX[i] == 0 {
			emptyPointsX = append(emptyPointsX, i)
		}
	}

	ans := 0

	for curr, starPoint := range starPoints {
		xCor := starPoint.x
		yCor := starPoint.y

		for _, endPoint := range starPoints[curr+1:] {
			xCor2 := endPoint.x
			yCor2 := endPoint.y

			dist := utils.AbsInt(xCor2-xCor) + utils.AbsInt(yCor2-yCor)

			// fmt.Printf("Actual dist: %d\n", dist)

			emptyPointsXLowerBound := lowerBound(emptyPointsX, utils.MinInt(xCor, xCor2))
			emptyPointsXUpperBound := upperBound(emptyPointsX, utils.Max(xCor, xCor2))

			emptyPointsYLowerBound := lowerBound(emptyPointsY, utils.MinInt(yCor, yCor2))
			emptyPointsYUpperBound := upperBound(emptyPointsY, utils.Max(yCor, yCor2))

			// fmt.Printf("X lower bound: %d\n", emptyPointsXLowerBound)
			// fmt.Printf("X upper bound: %d\n", emptyPointsXUpperBound)
			// fmt.Printf("Y lower bound: %d\n", emptyPointsYLowerBound)
			// fmt.Printf("Y upper bound: %d\n", emptyPointsYUpperBound)

			addtinalY := emptyPointsYUpperBound - emptyPointsYLowerBound + 1
			addtinalX := emptyPointsXUpperBound - emptyPointsXLowerBound + 1

			dist += addtinalX + addtinalY

			ans += dist

		}
	}

	fmt.Println(ans)

}

func Part02(input string) {

}
