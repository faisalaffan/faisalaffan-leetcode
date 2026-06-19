package main

// LeetCode #2025: Maximum Number of Ways to Partition an Array
// https://leetcode.com/problems/maximum-number-of-ways-to-partition-an-array/
// Difficulty: Hard
// Approach: Prefix sum + hash maps.
// Compute prefix sums. Count ways without changes.
// For each position i, compute ways if nums[i] is changed to k.
// Use left_freq (prefix sums before i) and right_freq (prefix sums from i to n-2).
// For j < i: condition is prefix[j] == new_total / 2 (unchanged)
// For j >= i: condition is prefix[j] + diff == new_total / 2, i.e. prefix[j] == target - diff

import "fmt"

func maxNumberOfWaysToPartition(nums []int, k int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Compute prefix sums
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}
	total := prefix[n-1]

	ans := 0

	// Without changes: count j where prefix[j]*2 == total (j < n-1)
	if total%2 == 0 {
		target := total / 2
		for j := 0; j < n-1; j++ {
			if prefix[j] == target {
				ans++
			}
		}
	}

	// With changes: try changing each nums[i] to k
	// leftFreq: prefix values for partition positions j < i
	// rightFreq: prefix values for partition positions j >= i
	leftFreq := make(map[int]int)
	rightFreq := make(map[int]int)
	for j := 0; j < n-1; j++ {
		rightFreq[prefix[j]]++
	}

	for i := 0; i < n; i++ {
		diff := k - nums[i]
		newTotal := total + diff

		if newTotal%2 == 0 {
			target := newTotal / 2
			cnt := 0

			// j < i: unchanged prefix
			cnt += leftFreq[target]

			// j >= i: prefix[j] + diff == target => prefix[j] == target - diff
			cnt += rightFreq[target-diff]

			if cnt > ans {
				ans = cnt
			}
		}

		// Move prefix[i] from rightFreq to leftFreq for next iteration
		if i < n-1 {
			rightFreq[prefix[i]]--
			if rightFreq[prefix[i]] == 0 {
				delete(rightFreq, prefix[i])
			}
			leftFreq[prefix[i]]++
		}
	}

	return ans
}

func main() {
	// Example: nums=[2,-1,2], k=3 -> 1
	fmt.Println(maxNumberOfWaysToPartition([]int{2, -1, 2}, 3))

	// Additional tests
	fmt.Println(maxNumberOfWaysToPartition([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxNumberOfWaysToPartition([]int{1, 1, 1}, 2))
	fmt.Println(maxNumberOfWaysToPartition([]int{0, 0, 0}, 1))
}
