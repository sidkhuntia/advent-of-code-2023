package day05

import (
	"fmt"
	"math"
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
	level int64
}

func dequeue(queue []Range) (Range, []Range) {
	if len(queue) == 0 {
		return Range{}, queue
	}
	if len(queue) == 1 {
		return queue[0], []Range{}
	}
	return queue[0], queue[1:]
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

func getDest(seed int64, conversionMap []MapItem) int64 {
	for _, mapItem := range conversionMap {
		if seed >= mapItem.src && seed < mapItem.src+mapItem.rangeLen {
			seed = mapItem.dest + (seed - mapItem.src)
			return seed
		}
	}

	return seed
}

func Part01(input string) {
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
			seed = getDest(seed, conversionMap)
		}
		// fmt.Printf("Seed: %d\n", seed)
		// fmt.Println()
		ans = utils.MinInt64(ans, seed)
		// fmt.Println("----------------------------")
	}

	fmt.Printf("Answer: %d\n", ans)

}

func Part02(input string) {
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
	seedsStr := strings.Fields(seeds)

	var seedRangeQueue []Range

	for i := 0; i < len(seedsStr)-1; i += 2 {
		startRange, _ := strconv.ParseInt(seedsStr[i], 10, 64)
		rangeLen, _ := strconv.ParseInt(seedsStr[i+1], 10, 64)
		endRange := startRange + rangeLen

		seedRangeQueue = append(seedRangeQueue, Range{
			start: startRange,
			end:   endRange,
			level: int64(0),
		})

	}

	// fmt.Println(seedRangeQueue)

	for len(seedRangeQueue) > 0 {
		seedRange, newQueue := dequeue(seedRangeQueue)
		seedRangeQueue = newQueue

		// fmt.Printf("Seed Range: %v\n", seedRange)

		seedStart := seedRange.start
		seedEnd := seedRange.end
		level := seedRange.level

		if level == 7 {
			ans = utils.MinInt64(ans, seedStart)
			continue
		}
		flag := 1
		for _, mapItem := range conversionMaps[level] {
			mapStart := mapItem.src
			mapEnd := mapItem.src + mapItem.rangeLen
			mapDest := mapItem.dest

			// mapRange := Range{
			// 	start: mapStart,
			// 	end:   mapEnd,

			// }

			// fmt.Printf("Map Range: %v\n", mapRange)

			offset := mapDest - mapStart

			if mapStart >= seedEnd || mapEnd <= seedStart {
				continue
			}

			if seedStart < mapStart {
				seedRangeQueue = append(seedRangeQueue, Range{
					start: seedStart,
					end:   mapStart,
					level: level,
				})
				seedStart = mapStart
				// fmt.Printf("Partial Overlap 1: %v\n", seedRangeQueue)
			}
			if seedEnd > mapEnd {
				seedRangeQueue = append(seedRangeQueue, Range{
					start: mapEnd,
					end:   seedEnd,
					level: level,
				})
				seedEnd = mapEnd
				// fmt.Printf("Partial Overlap 2: %v\n", seedRangeQueue)
			}
			seedRangeQueue = append(seedRangeQueue, Range{
				start: seedStart + offset,
				end:   seedEnd + offset,
				level: level + 1,
			})
			// fmt.Printf("Complete Overlap: %v\n", seedRangeQueue)
			flag = 0
			break

		}

		if flag == 1 {
			seedRangeQueue = append(seedRangeQueue, Range{
				start: seedStart,
				end:   seedEnd,
				level: level + 1,
			})
		}

	}

	fmt.Printf("Answer: %d\n", ans)

}
