package main

// LeetCode #2195: Append K Integers With Minimal Sum
// https://leetcode.com/problems/append-k-integers-with-minimal-sum/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	var sum int64 = 0
	prev := 0

	for _, num := range nums {
		if num == prev {
			continue
		}
		if num > prev+1 {
			gap := num - prev - 1
			if gap >= k {
				// arithmetic series: (prev+1) + (prev+2) + ... + (prev+k)
				first := int64(prev) + 1
				last := int64(prev) + int64(k)
				sum += (first + last) * int64(k) / 2
				k = 0
				break
			}
			first := int64(prev) + 1
			last := int64(num) - 1
			sum += (first + last) * int64(gap) / 2
			k -= gap
		}
		prev = num
		if k == 0 {
			break
		}
	}

	if k > 0 {
		first := int64(prev) + 1
		last := int64(prev) + int64(k)
		sum += (first + last) * int64(k) / 2
	}

	return sum
}

func main() {
	// Test case 1
	fmt.Println(minimalKSum([]int{1, 4, 25, 10, 25}, 2))
	// Expected: 5

	// Test case 2
	fmt.Println(minimalKSum([]int{5, 6}, 6))
	// Expected: 25
}
