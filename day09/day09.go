package day09

import (
	"fmt"
	"strings"

	"github.com/sidkhuntia/advent-of-code-2023/utils"
)

func predictNextNumForPart01(nums []int) int {
	ans := nums[len(nums)-1]
	lengthOfNums := len(nums)
	for steps := 0; steps < lengthOfNums; steps++ {
		flag := true
		tempNums := make([]int, 0)
		for i := 1; i < len(nums); i++ {
			diff := nums[i] - nums[i-1]

			if diff != 0 {
				flag = false
			}

			tempNums = append(tempNums, diff)

			if i == len(nums)-1 {
				ans += diff
			}
		}

		nums = tempNums

		if flag {
			break
		}
	}

	return ans
}
func predictNextNumForPart02(nums []int) int {
	ans := nums[0]
	lengthOfNums := len(nums)

	for steps := 0; steps < lengthOfNums; steps++ {
		flag := true
		tempNums := make([]int, 0)
		for i := 1; i < len(nums); i++ {
			diff := nums[i] - nums[i-1]

			if diff != 0 {
				flag = false
			}

			tempNums = append(tempNums, diff)

			if i == 1 {
				if steps%2 == 1 {
					ans += diff
				} else {
					ans -= diff
				}

			}
		}

		nums = tempNums

		if flag {
			break
		}
	}
	return ans
}

func Part01(input string) {
	lines := utils.SplitLines(input)
	ans := 0
	for _, line := range lines {
		historyNumStrArr := strings.Fields(line)
		historyNumArr := make([]int, 0)

		for _, numStr := range historyNumStrArr {
			num := utils.ParseInt(numStr)

			historyNumArr = append(historyNumArr, num)
		}

		newNum := predictNextNumForPart01(historyNumArr)
		ans += newNum
		// fmt.Printf("newNum: %v\n", newNum)
	}

	fmt.Printf("Answer: %v\n", ans)
}

func Part02(input string) {
	lines := utils.SplitLines(input)
	ans := 0
	for _, line := range lines {
		historyNumStrArr := strings.Fields(line)
		historyNumArr := make([]int, 0)

		for _, numStr := range historyNumStrArr {
			num := utils.ParseInt(numStr)

			historyNumArr = append(historyNumArr, num)
		}

		newNum := predictNextNumForPart02(historyNumArr)
		ans += newNum
		// fmt.Printf("newNum: %v\n", newNum)
	}

	fmt.Printf("Answer: %v\n", ans)
}
