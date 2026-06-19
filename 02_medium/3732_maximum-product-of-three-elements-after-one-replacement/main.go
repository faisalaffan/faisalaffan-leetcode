package main

// LeetCode #3732: Maximum Product of Three Elements After One Replacement
// https://leetcode.com/problems/maximum-product-of-three-elements-after-one-replacement/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumProductOfThreeElementsAfterOneReplacement(nums []int) int64 {
	var first, second int64 = 0, 0
	for _, v := range nums {
		val := int64(v)
		if val < 0 {
			val = -val
		}
		if val > first {
			second = first
			first = val
		} else if val > second {
			second = val
		}
	}
	return 100000 * first * second
}

func main() {
	fmt.Println(maximumProductOfThreeElementsAfterOneReplacement([]int{1, 2, 3, 4}))
	fmt.Println(maximumProductOfThreeElementsAfterOneReplacement([]int{-4, -2, -1, -3}))
	fmt.Println(maximumProductOfThreeElementsAfterOneReplacement([]int{-5, 7, 0}))
}
