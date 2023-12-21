package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sidkhuntia/advent-of-code-2023/utils"
)

func Part01(input string) {
	lines := strings.Split(input, "\n")

	var times []int
	var distances []int

	for lineNumber, line := range lines {
		fields := strings.Fields(line)
		for _, field := range fields[1:] {
			num := utils.ParseInt(field)

			if lineNumber == 0 {
				times = append(times, num)
			} else {
				distances = append(distances, num)
			}

		}
	}

	ans := 0

	for timeIdx, time := range times {
		count := 0
		recordDistance := distances[timeIdx]
		for seconds := 0; seconds <= time; seconds++ {
			myDistance := seconds * (time - seconds)
			if myDistance > recordDistance {
				count++
			}
		}

		if ans == 0 {
			ans = count
		} else {
			ans *= count
		}
	}

	fmt.Printf("Answer: %d\n", ans)
}

func Part02(input string) {
	lines := strings.Split(input, "\n")

	recordTime := ""
	recordDistance := ""
	for lineNumber, line := range lines {
		fields := strings.Fields(line)
		for _, field := range fields[1:] {

			if lineNumber == 0 {
				recordTime = recordTime + field
			} else {
				recordDistance = recordDistance + field
			}

		}
	}

	recordTimeint, _ := strconv.ParseInt(recordTime, 10, 64)
	recordDistanceint, _ := strconv.ParseInt(recordDistance, 10, 64)

	ans := 0
	for seconds := int64(0); seconds <= (recordTimeint); seconds++ {
		myDistance := seconds * (recordTimeint - seconds)
		if myDistance > recordDistanceint {
			ans++
		}
	}

	fmt.Printf("Answer: %d\n", ans)
}
