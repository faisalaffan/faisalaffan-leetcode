package main

// LeetCode #2098: Subsequence of Size K With the Largest Even Sum
// https://leetcode.com/problems/subsequence-of-size-k-with-the-largest-even-sum/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func largestEvenSum(nums []int, k int) int64 {
	// Separate evens and odds, sort descending
	evens := []int{}
	odds := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(evens)))
	sort.Sort(sort.Reverse(sort.IntSlice(odds)))

	// Prefix sums
	ePrefix := make([]int64, len(evens)+1)
	for i, v := range evens {
		ePrefix[i+1] = ePrefix[i] + int64(v)
	}
	oPrefix := make([]int64, len(odds)+1)
	for i, v := range odds {
		oPrefix[i+1] = oPrefix[i] + int64(v)
	}

	var result int64 = -1
	// Try picking i evens and (k-i) odds
	for i := 0; i <= k && i <= len(evens); i++ {
		j := k - i
		if j > len(odds) {
			continue
		}
		if i == 0 && j%2 != 0 {
			// Need at least one even number to make sum even
			if len(evens) == 0 {
				continue
			}
			// Try picking k odds with at least one even
			// Actually, problem requires subsequence of exactly size k
			// If all chosen numbers are odd and count is odd, sum is odd
			// We need at least 1 even
			_ = evens[0] // reference to avoid unused error
			continue
		}
		// Sum is even if we have even number of odd elements
		if j%2 == 0 {
			sum := ePrefix[i] + oPrefix[j]
			if sum > result {
				result = sum
			}
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", largestEvenSum([]int{4, 1, 5, 3, 1}, 3))
	// Expected: 12

	// Test case 2
	fmt.Println("Test 2:", largestEvenSum([]int{4, 6, 2}, 3))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", largestEvenSum([]int{1, 3, 5}, 3))
	// Expected: -1 (sum would be 9, odd)
}
