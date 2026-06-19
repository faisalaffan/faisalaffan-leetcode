package main

// LeetCode #2578: Split With Minimum Sum
// https://leetcode.com/problems/split-with-minimum-sum/
// Difficulty: Easy
// Time O(log n) | Space O(log n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SplitWithMinimumSum(4325)) // 59
	fmt.Println(SplitWithMinimumSum(687))  // 75
}

func SplitWithMinimumSum(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Ints(digits)

	num1, num2 := 0, 0
	for i := 0; i < len(digits); i++ {
		if i%2 == 0 {
			num1 = num1*10 + digits[i]
		} else {
			num2 = num2*10 + digits[i]
		}
	}
	return num1 + num2
}
