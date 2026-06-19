package main

// LeetCode #2302: Count Subarrays With Score Less Than K
// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/
// Difficulty: Hard
//
// Approach: Sliding window. Score of subarray nums[l..r] = sum(nums[l..r]) * (r-l+1).
// Expand right pointer. While score >= k, shrink left. For each valid window,
// count = r-l+1 subarrays ending at r. Accumulate.

import "fmt"

func main() {
	// Example 1: [2,1,4,3,5], k=10 => 6
	fmt.Println(countSubarrays([]int{2, 1, 4, 3, 5}, 10))
	// Example 2: [1,1,1], k=5 => 5
	fmt.Println(countSubarrays([]int{1, 1, 1}, 5))
	// Edge: all elements too large
	fmt.Println(countSubarrays([]int{10, 10, 10}, 5))
	// Edge: single element less than k
	fmt.Println(countSubarrays([]int{3}, 10))
	// Edge: k=1
	fmt.Println(countSubarrays([]int{1, 2, 3}, 1))
}

func countSubarrays(nums []int, k int64) int64 {
	var ans int64
	var sum int64
	left := 0
	for right := 0; right < len(nums); right++ {
		sum += int64(nums[right])
		// Shrink while score >= k
		// score = sum * (right-left+1)
		for sum*int64(right-left+1) >= k {
			sum -= int64(nums[left])
			left++
			if left > right {
				break
			}
		}
		if left <= right {
			ans += int64(right - left + 1)
		}
	}
	return ans
}
