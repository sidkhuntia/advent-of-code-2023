package day02

import (
	"fmt"
	"math"
	// "math"
	"os"
	"strconv"
	"strings"
)

const RED int = 12
const GREEN int = 13
const BLUE int = 14

func main() {
	input, err := os.ReadFile("sample.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	stringInput := string(input)

	rows := strings.Split(stringInput, "\n")
	var rowSet []string

	for _, row := range rows {
		temp := strings.Split(row, ": ")
		rowSet = append(rowSet, temp[1])
	}

	var sets [][]string

	for _, row := range rowSet {
		temp := strings.Split(row, "; ")

		sets = append(sets, temp)
	}

	ans := 0

	for _, row := range sets {
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, col := range row {
			colorsSet := strings.Split(col, ", ")
			// fmt.Println(col)
			for _, color := range colorsSet {
				checkSet := strings.Split(color, " ")

				quantity, err := strconv.Atoi(checkSet[0])
				if err != nil {
					fmt.Println("error converting string to int")
				}
				currColor := checkSet[1]
				fmt.Println(color)
				fmt.Printf("color: %v, quantity: %v\n", checkSet[1], len(checkSet))

				// TODO: its not printing the last color name 
				if currColor == "blue" {
					maxBlue = int(math.Max(float64(maxBlue), float64(quantity)))
				} else if currColor == "red" {
					maxRed = int(math.Max(float64(maxRed), float64(quantity)))
					
				} else if currColor == "green" {
					maxGreen = int(math.Max(float64(maxGreen), float64(quantity)))
				}
			}

		}

		power := (maxBlue * maxGreen * maxRed)
		ans += power
		// fmt.Println("----------------------------")
	}

	fmt.Printf("Ans: %v\n", ans)

}
