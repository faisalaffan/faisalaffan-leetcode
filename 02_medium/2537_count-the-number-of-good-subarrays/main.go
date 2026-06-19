package main

// LeetCode #2537: Count the Number of Good Subarrays
// https://leetcode.com/problems/count-the-number-of-good-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countGood(nums []int, k int) int64 {
	n := len(nums)
	freq := make(map[int]int)
	var pairs int64
	var ans int64
	left := 0

	for right := 0; right < n; right++ {
		val := nums[right]
		pairs += int64(freq[val])
		freq[val]++

		for pairs >= int64(k) {
			ans += int64(n - right)
			leftVal := nums[left]
			freq[leftVal]--
			pairs -= int64(freq[leftVal])
			left++
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countGood([]int{1, 1, 1, 1, 1}, 10))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", countGood([]int{1, 2, 3}, 1))
	// Expected: 0
}
