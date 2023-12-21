package day01

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var DIGITS []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var value map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func calculateValue(row string) int {
	left := 0
	right := 0
	len := len(row)
	for _, char := range row {
		if char >= '0' && char <= '9' {
			left = int(char - '0')
			break
		}
	}

	for i := len - 1; i >= 0; i-- {
		if char := row[i]; char >= '0' && char <= '9' {
			right = int(char - '0')
			break
		}
	}

	return left*10 + right
}

func calculateValue2(row string) int {
	left := 0
	right := 0
	leftIdx := math.MaxInt
	rightIdx := math.MinInt
	len := len(row)

	for _, digit := range DIGITS {
		idx := strings.Index(row, digit)
		if idx != -1 && idx < leftIdx {
			left = value[digit]
			leftIdx = idx
		}
	}

	for _, digit := range DIGITS {
		idx := strings.LastIndex(row, digit)
		if idx != -1 && idx > rightIdx {
			right = value[digit]
			rightIdx = idx
		}
	}

	for i, char := range row {
		if char >= '0' && char <= '9' && i < leftIdx {
			left = int(char - '0')
			break
		}
	}

	for i := len - 1; i >= 0; i-- {
		if char := row[i]; char >= '0' && char <= '9' && i > rightIdx {
			right = int(char - '0')
			break
		}
	}

	return left*10 + right

}

func main() {
	input, err := os.ReadFile("input2.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	stringInput := string(input)

	rows := strings.Split(stringInput, "\n")

	// fmt.Printf("%T", rows)
	ans := 0
	for _, row := range rows {
		ans += calculateValue2(row)
		// fmt.Println(val)
		// ans += val
	}

	fmt.Println(ans)
}
