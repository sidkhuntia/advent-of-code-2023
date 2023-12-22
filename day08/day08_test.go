package day08_test

import (
	"log"
	"testing"

	"github.com/sidkhuntia/advent-of-code-2023/day08"
	"github.com/sidkhuntia/advent-of-code-2023/utils"
)

func BenchmarkPart01(b *testing.B) {
	filePath := "../data/day08/input.txt"
	input, err := utils.ReadFileIntoString(filePath)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}
	for i := 0; i < b.N; i++ {
		day08.Part01(input)
	}
}
func BenchmarkPart02(b *testing.B) {
	filePath := "../data/day08/input.txt"
	input, err := utils.ReadFileIntoString(filePath)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}
	for i := 0; i < b.N; i++ {
		day08.Part02(input)
	}
}

