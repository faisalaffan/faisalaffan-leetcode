package main

// LeetCode #1712: Ways to Split Array Into Three Subarrays
// https://leetcode.com/problems/ways-to-split-array-into-three-subarrays/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

const mod = 1_000_000_007

func waysToSplit(nums []int) int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	total := prefix[n-1]
	result := 0
	left := 0

	// For each possible left split point i (first part ends at i)
	for i := 0; i < n-2; i++ {
		// Left sum = prefix[i]
		left += nums[i]

		// Find min j where mid sum >= left sum
		// mid sum = prefix[j] - prefix[i], need prefix[j] >= 2*prefix[i]
		minMid := lowerBound(prefix, 2*left, i+1)

		// Find max j where mid sum <= right sum
		// right sum = total - prefix[j], need prefix[j] <= (total + left) / 2
		maxMid := upperBound(prefix, (total+left)/2, i+1, n-2)

		if minMid <= maxMid {
			result = (result + (maxMid - minMid + 1)) % mod
		}
	}
	return result
}

func lowerBound(arr []int, target int, start int) int {
	lo, hi := start, len(arr)-2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] >= target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func upperBound(arr []int, target int, start int, end int) int {
	lo, hi := start, end
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func main() {
	fmt.Println(waysToSplit([]int{1, 1, 1}))             // Expected: 1
	fmt.Println(waysToSplit([]int{1, 2, 2, 2, 5, 0}))   // Expected: 3
	fmt.Println(waysToSplit([]int{3, 2, 1}))             // Expected: 0
}
