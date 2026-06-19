package main

// LeetCode #3300: Minimum Element After Replacement With Digit Sum
// https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumElementAfterReplacementWithDigitSum([]int{10, 12, 13, 14}))
	fmt.Println(MinimumElementAfterReplacementWithDigitSum([]int{1, 2, 3, 4}))
	fmt.Println(MinimumElementAfterReplacementWithDigitSum([]int{999, 19, 199}))
}

// digitSum returns the sum of digits of n.
func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// MinimumElementAfterReplacementWithDigitSum returns the minimum element after replacing each element with its digit sum.
// Time: O(n * log n). Space: O(1).
func MinimumElementAfterReplacementWithDigitSum(nums []int) int {
	minVal := int(^uint(0) >> 1) // MaxInt
	for _, num := range nums {
		s := digitSum(num)
		if s < minVal {
			minVal = s
		}
	}
	return minVal
}
