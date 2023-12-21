package day05_test

import (
	"github.com/sidkhuntia/advent-of-code-2023/day05"
	"github.com/sidkhuntia/advent-of-code-2023/utils"
	"log"
	"os"
	"testing"
)

func BenchmarkPart01(b *testing.B) {
	filePath := os.Args[2]
	input, err := utils.ReadFileIntoString(filePath)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}
	for i := 0; i < b.N; i++ {
		day05.Part01(input)
	}

}

func BenchmarkPart02(b *testing.B) {
	filePath := os.Args[2]
	input, err := utils.ReadFileIntoString(filePath)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}
	for i := 0; i < b.N; i++ {
		day05.Part02(input)
	}

}
