package day07

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sidkhuntia/advent-of-code-2023/utils"
)

type rankFn func(string) int

var ranks = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func rankOfHand(hand string) int {
	frequency := make(map[string]int)

	for _, card := range hand {
		frequency[string(card)]++
	}
	// fmt.Println(frequency)

	switch len(frequency) {
	case 1:
		return 7
	case 2:
		rank := 5
		for _, v := range frequency {
			if v == 4 {
				rank = 6
			}
		}
		return rank
	case 3:
		for _, v := range frequency {
			if v == 3 {
				return 4
			}
		}
		return 3
	case 4:
		return 2
	case 5:
		return 1
	}

	return -1
}
func rankOfHandForPart02(hand string) int {
	frequency := make(map[string]int)

	for _, card := range hand {
		frequency[string(card)]++
	}
	// fmt.Println(frequency)

	maxCardOtherThanJoker := "*"
	maxFrequency := 0

	for card, freq := range frequency {
		if card == "J" {
			continue
		}
		if freq > maxFrequency {
			maxFrequency = freq
			maxCardOtherThanJoker = card
		}
	}

	frequency[maxCardOtherThanJoker] += frequency["J"]
	delete(frequency, "J")

	switch len(frequency) {
	case 1:
		return 7
	case 2:
		rank := 5
		for _, v := range frequency {
			if v == 4 {
				rank = 6
			}
		}
		return rank
	case 3:
		for _, v := range frequency {
			if v == 3 {
				return 4
			}
		}
		return 3
	case 4:
		return 2
	case 5:
		return 1
	}

	return -1
}

func compareSameRankHands(hand1, hand2 string) bool {
	for i := 0; i < len(hand1); i++ {
		idx1 := ranks[string(hand1[i])]
		idx2 := ranks[string(hand2[i])]

		if idx1 == idx2 {
			continue
		}

		return idx1 < idx2
	}
	return true
}

func compareHands(hand1 string, hand2 string, rankFunc rankFn) bool {
	rankOfHand1 := rankFunc(hand1)
	rankOfHand2 := rankFunc(hand2)

	// fmt.Printf("Hand1: %s, Rank: %d\n", hand1, rankOfHand1)
	// fmt.Printf("Hand2: %s, Rank: %d\n", hand2, rankOfHand2)
	// fmt.Println()

	if rankOfHand1 == rankOfHand2 {
		return compareSameRankHands(hand1, hand2)
	}
	return rankOfHand1 < rankOfHand2
}

func Part01(input string) {
	lines := utils.SplitLines(input)
	cardToScore := make(map[string]int)
	var cards []string
	for _, line := range lines {
		fields := strings.Fields(line)

		card := fields[0]
		score := utils.ParseInt(fields[1])
		cards = append(cards, card)
		cardToScore[card] = score

	}

	sort.Slice(cards, func(i, j int) bool {
		return compareHands(cards[i], cards[j], rankOfHand)
	})
	ans := 0
	for cardIdx, card := range cards {
		ans += cardToScore[card] * (cardIdx + 1)
	}

	fmt.Printf("Answer: %d\n", ans)
}

func Part02(input string) {
	lines := utils.SplitLines(input)
	cardToScore := make(map[string]int)
	var cards []string
	for _, line := range lines {
		fields := strings.Fields(line)

		card := fields[0]
		score := utils.ParseInt(fields[1])
		cards = append(cards, card)
		cardToScore[card] = score

	}

	sort.Slice(cards, func(i, j int) bool {
		return compareHands(cards[i], cards[j], rankOfHandForPart02)
	})
	ans := 0
	for cardIdx, card := range cards {
		ans += cardToScore[card] * (cardIdx + 1)
	}

	fmt.Printf("Answer: %d\n", ans)
}
