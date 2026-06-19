package main

// LeetCode #3917: Count Indices With Opposite Parity
// https://leetcode.com/problems/count-indices-with-opposite-parity/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountIndicesWithOppositeParity([]int{1, 2, 3, 4}))
	fmt.Println(CountIndicesWithOppositeParity([]int{2, 4, 6}))
}

// Time: O(n)
// Space: O(1)
func CountIndicesWithOppositeParity(nums []int) []int {
	totalEven, totalOdd := 0, 0
	for _, v := range nums {
		if v%2 == 0 {
			totalEven++
		} else {
			totalOdd++
		}
	}

	ans := make([]int, len(nums))
	for i, v := range nums {
		if v%2 == 0 {
			totalEven--
			ans[i] = totalOdd
		} else {
			totalOdd--
			ans[i] = totalEven
		}
	}
	return ans
}
