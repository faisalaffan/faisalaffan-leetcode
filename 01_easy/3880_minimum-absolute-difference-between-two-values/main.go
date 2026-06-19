package main

// LeetCode #3880: Minimum Absolute Difference Between Two Values
// https://leetcode.com/problems/minimum-absolute-difference-between-two-values/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumAbsoluteDifferenceBetweenTwoValues([]int{1, 0, 0, 2, 0, 1}))
	fmt.Println(MinimumAbsoluteDifferenceBetweenTwoValues([]int{1, 0, 1, 0}))
}

// Time: O(n)
// Space: O(1)
func MinimumAbsoluteDifferenceBetweenTwoValues(nums []int) int {
	last1, last2 := -1, -1
	ans := -1
	for i, v := range nums {
		if v == 1 {
			last1 = i
			if last2 != -1 {
				diff := i - last2
				if ans == -1 || diff < ans {
					ans = diff
				}
			}
		} else if v == 2 {
			last2 = i
			if last1 != -1 {
				diff := i - last1
				if ans == -1 || diff < ans {
					ans = diff
				}
			}
		}
	}
	return ans
}
