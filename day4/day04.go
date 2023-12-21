package day04

import (
	"fmt"
	"os"
	"strings"
)

type void struct{}

var member void

func solve() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	ans := 0
	stringInput := string(input)

	rows := strings.Split(stringInput, "\n")
	var rowSet []string

	for _, row := range rows {
		temp := strings.Split(row, ": ")
		numbers := strings.TrimSpace(temp[1])
		rowSet = append(rowSet, numbers)

	}

	prefixSum := make([]int, len(rowSet))

	for i := 0; i < len(rowSet); i++ {
		prefixSum[i] = 1
	}

	for cardIdx, row := range rowSet {
		numArr := strings.Split(row, " | ")

		set := make(map[string]void)

		winningNums := strings.Split(numArr[0], " ")

		for _, num := range winningNums {
			// formattedNum := strings.ReplaceAll(num, "\"", "")
			num = strings.TrimSuffix(num, "\n")
			if num == "" {
				continue
			}
			set[num] = member
		}

		numberWeHave := strings.Split(numArr[1], " ")
		count := 0
		for _, num := range numberWeHave {
			num = strings.TrimSuffix(num, "\n")
			// formattedNum := strings.ReplaceAll(num, "\"", "")
			if num == "" {
				continue
			}
			if _, ok := set[num]; ok {
				count++
			}
		}
		for i := 0; i < count; i++ {
			prefixSum[cardIdx+1+i] += prefixSum[cardIdx]
		}

	}

	for _, num := range prefixSum {
		ans += num
	}

	fmt.Printf("Ans: %v\n", ans)

}

func main() {
	solve()
}
