package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sidkhuntia/advent-of-code-2023/utils"
)

type MapItem struct {
	src      int64
	dest     int64
	rangeLen int64
}

type Range struct {
	start int64
	end   int64
}

func newMapItem(mapStr string) MapItem {
	numbers := strings.Fields(mapStr)

	dest, _ := strconv.ParseInt(numbers[0], 10, 64)
	src, _ := strconv.ParseInt(numbers[1], 10, 64)
	rangeLen, _ := strconv.ParseInt(numbers[2], 10, 64)

	return MapItem{
		src:      (src),
		dest:     (dest),
		rangeLen: (rangeLen),
	}
}

func GetDest(seed int64, conversionMap []MapItem) int64 {
	for _, mapItem := range conversionMap {
		if seed >= mapItem.src && seed < mapItem.src+mapItem.rangeLen {
			seed = mapItem.dest + (seed - mapItem.src)
			return seed
		}
	}

	return seed
}

func GetDestRange(seedRange Range, conversionMap []MapItem) Range {
	for _, mapItem := range conversionMap {
		if mapItem.src > seedRange.end || mapItem.src+mapItem.rangeLen < seedRange.start {
			continue
		}

		offset := mapItem.dest - mapItem.src

		return Range{
			start: seedRange.start + offset,
			end:   seedRange.end + offset,
		}
	}

	return seedRange
}

func solve() {
	input, err := utils.ReadFileIntoString("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	re := regexp.MustCompile(`\r?\n\r?\n`)
	mapContent := re.Split(input, -1)

	var conversionMaps [][]MapItem

	for _, element := range mapContent[1:] {

		mapStrs := utils.SplitLines(element)
		var conversionMap []MapItem

		for _, mapStr := range mapStrs[1:] {
			if len(mapStr) == 0 {
				continue
			}

			conversionMap = append(conversionMap, newMapItem(mapStr))
		}

		conversionMaps = append(conversionMaps, conversionMap)
	}

	ans := int64(math.MaxInt64)

	_, seeds, _ := strings.Cut(mapContent[0], ": ")
	// seedsStr := strings.Fields(seeds)
	for _, seedStr := range strings.Fields(seeds) {
		seed, _ := (strconv.ParseInt(seedStr, 10, 64))
		for _, conversionMap := range conversionMaps {
			seed = GetDest(seed, conversionMap)
		}
		// fmt.Printf("Seed: %d\n", seed)
		// fmt.Println()
		ans = utils.MinInt64(ans, seed)
		// fmt.Println("----------------------------")
	}

	fmt.Printf("Answer: %d\n", ans)

}

func main() {
	solve()
}
