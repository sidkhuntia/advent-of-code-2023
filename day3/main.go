package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DIR = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

func getNum(matrix []string, visited [][]bool, i int, j int) int {
	num := ""

	for k := j; k < len(matrix[i]); k++ {
		charAt := (matrix[i][k])
		if charAt >= '0' && charAt <= '9' {
			num = num + string(charAt)
			visited[i][k] = true
		} else {
			break
		}
	}
	// fmt.Println(j)
	for k := j - 1; k >= 0; k-- {
		charAt := (matrix[i][k])
		if charAt >= '0' && charAt <= '9' {
			num = string(charAt) + num
			visited[i][k] = true
		} else {
			break
		}
	}
	ans, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("error converting string to int")
	}
	return ans

}

func dfs(matrix []string, visited [][]bool, i int, j int) int {
	// if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[i]) {
	// 	return 0
	// }
	ans := 1
	count := 0
	
	for _, dir := range DIR {
		newI := i + dir[0]
		newJ := j + dir[1]
		// TODO: use benchmark to test which is faster
		if newI < 0 || newJ < 0 || newI >= len(matrix) || newJ >= len(matrix[i]) {
			continue
		}
		charAt := (matrix[newI][newJ])

		if charAt >= '0' && charAt <= '9' && !visited[newI][newJ] {
			num := getNum(matrix, visited, newI, newJ)
			ans *= num
			count += 1
		}
	}
	
	if count == 2 && matrix[i][j] == '*' {
		return ans
	}

	return 0
}

func main() {
	input, err := os.ReadFile("input2.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	stringInput := string(input)
	ans := 0
	matrix := strings.Split(stringInput, "\n")
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			charAt := matrix[i][j]
			if charAt == '.' || (charAt >= '0' && charAt <= '9') || charAt == 13 {
				continue
			} else {
				// fmt.Println(matrix[i])
				ans += dfs(matrix, visited, i, j)
			}
		}
	}

	fmt.Printf("Ans: %v\n", ans)

}
