package main

// LeetCode #3388: Count Beautiful Splits in an Array
// https://leetcode.com/problems/count-beautiful-splits-in-an-array/
// Difficulty: Medium
// Time: O(n^2) Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(beautifulSplits([]int{1, 1, 2, 1})) // 2
	fmt.Println(beautifulSplits([]int{1, 2, 3, 4})) // 0
}

func beautifulSplits(nums []int) int {
	n := len(nums)

	// lcp[i][j] = longest common prefix of nums[i:] and nums[j:]
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums[i] == nums[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			// nums1 = [0,i), nums2 = [i,j), nums3 = [j,n)
			ok := false
			// Check if nums1 is prefix of nums2
			if lcp[0][i] >= i {
				ok = true
			}
			// Check if nums2 is prefix of nums3
			if lcp[i][j] >= j-i {
				ok = true
			}
			if ok {
				ans++
			}
		}
	}
	return ans
}
