package main

// LeetCode #3164: Find the Number of Good Pairs II
// https://leetcode.com/problems/find-the-number-of-good-pairs-ii/
// Difficulty: Medium
// Time: O(n * sqrt(max) + m) | Space: O(max)

import "fmt"

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	freq := make(map[int]int)
	for _, v := range nums1 {
		if v%k != 0 {
			continue
		}
		v /= k
		for d := 1; d*d <= v; d++ {
			if v%d == 0 {
				freq[d]++
				if d*d != v {
					freq[v/d]++
				}
			}
		}
	}

	var ans int64
	for _, v := range nums2 {
		ans += int64(freq[v])
	}
	return ans
}

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1)) // Expected: 5
	fmt.Println(numberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3)) // Expected: 2
}
