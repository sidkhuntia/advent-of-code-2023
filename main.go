package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sidkhuntia/advent-of-code-2023/utils"

	// "github.com/sidkhuntia/advent-of-code-2023/day1"
	// "github.com/sidkhuntia/advent-of-code-2023/day2"
	// "github.com/sidkhuntia/advent-of-code-2023/day3"
	// "github.com/sidkhuntia/advent-of-code-2023/day4"
	"github.com/sidkhuntia/advent-of-code-2023/day05"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("You need to provide both the day and the path to input!\n\t%s <day> <path/to/sample.txt>\n", os.Args[0])
	}

	day := os.Args[1]
	filePath := os.Args[2]
	debugMode := ""

	if len(os.Args) > 3 {
		debugMode = os.Args[3]
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File not found: %s\n", filePath)
	}

	input, err := utils.ReadFileIntoString(filePath)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}

	switch day {
	case "day05":
		runDay(day, day05.Part01, day05.Part02, input, debugMode)

	default:
		log.Fatalf("Unknown day: %s\n", day)
	}
}

func runDay(day string, part01, part02 func(string), input string, debugMode string) {
	fmt.Println()

	switch debugMode {
	case "01":
		fmt.Println("Debug mode 01")
		fmt.Println("Running Part01")
		fmt.Println()
		part01(input)

	case "02":
		fmt.Println("Debug mode 02")
		fmt.Println("Running Part02")
		fmt.Println()
		part02(input)
	default:
		break
	}
	if debugMode != "" {
		return
	}
	fmt.Printf("===== %s =====\n", day)

	if part01 != nil {
		start := time.Now()
		fmt.Printf("Part01: ")
		part01(input)
		elapsed := time.Since(start)
		fmt.Printf("\t(took %s)\n", elapsed)
	}

	if part02 != nil {
		start := time.Now()
		fmt.Printf("Part02: ")
		part02(input)
		elapsed := time.Since(start)
		fmt.Printf("\t(took %s)\n", elapsed)
	}

	fmt.Println()
}
