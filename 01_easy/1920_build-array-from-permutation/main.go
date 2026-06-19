package main

// LeetCode #1920: Build Array from Permutation
// https://leetcode.com/problems/build-array-from-permutation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(BuildArrayFromPermutation([]int{0, 2, 1, 5, 3, 4}))    // [0,1,2,4,5,3]
	fmt.Println(BuildArrayFromPermutation([]int{5, 0, 1, 2, 3, 4}))    // [4,5,0,1,2,3]
}

// Time: O(n), Space: O(n)
func BuildArrayFromPermutation(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = nums[v]
	}
	return ans
}
