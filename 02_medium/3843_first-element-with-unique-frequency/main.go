package main

// LeetCode #3843: First Element with Unique Frequency
// https://leetcode.com/problems/first-element-with-unique-frequency/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Count frequencies, then count frequency-of-frequency, find first with freq=1.

import "fmt"

func FirstElementWithUniqueFrequency(nums []int) int {
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}

	freqCnt := make(map[int]int)
	for _, c := range cnt {
		freqCnt[c]++
	}

	for _, v := range nums {
		if freqCnt[cnt[v]] == 1 {
			return v
		}
	}

	return -1
}

func main() {
	// Example 1
	fmt.Println(FirstElementWithUniqueFrequency([]int{20, 10, 30, 30})) // Expected: 30

	// Example 2
	fmt.Println(FirstElementWithUniqueFrequency([]int{20, 20, 10, 30, 30, 30})) // Expected: 20

	// Example 3
	fmt.Println(FirstElementWithUniqueFrequency([]int{10, 10, 20, 20})) // Expected: -1
}
