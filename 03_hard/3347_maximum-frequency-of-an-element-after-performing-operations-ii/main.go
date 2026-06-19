package main

// LeetCode #3347: Maximum Frequency of an Element After Performing Operations II
// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-ii/
// Difficulty: Hard
//
// Given an array nums, you can perform at most numOperations operations. In
// each operation, you pick an element and change it to any integer within
// [element-k, element+k]. Maximize the frequency of any value after operations.
//
// Approach: Use difference array and sweep line. For each element, it can
// become any value in [nums[i]-k, nums[i]+k]. Track coverage counts using
// events, plus count of elements already at each value.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maxFrequency([]int{1, 4, 5}, 1, 2))
	// Example 2
	fmt.Println(maxFrequency([]int{5, 3, 5}, 2, 1))
	// Example 3: all same
	fmt.Println(maxFrequency([]int{1, 1, 1, 1}, 0, 0))
	// Edge: single element
	fmt.Println(maxFrequency([]int{5}, 10, 3))
	// No operations
	fmt.Println(maxFrequency([]int{1, 2, 3}, 0, 0))
}

func maxFrequency(nums []int, k int, numOperations int) int {
	// Frequency of each original value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Events for sweep line: +1 when entering range, -1 when leaving
	type event struct {
		pos    int
		delta  int
	}
	events := make([]event, 0)
	for _, v := range nums {
		events = append(events, event{v - k, 1})  // start of reachable range
		events = append(events, event{v + k + 1, -1}) // end of reachable range
	}

	// Sort events by position
	sort.Slice(events, func(i, j int) bool {
		if events[i].pos != events[j].pos {
			return events[i].pos < events[j].pos
		}
		return events[i].delta > events[j].delta
	})

	ans := 0
	coverage := 0
	i := 0
	n := len(events)

	// Also collect unique values for existing element processing
	unique := make([]int, 0, len(freq))
	for val := range freq {
		unique = append(unique, val)
	}
	sort.Ints(unique)

	valIdx := 0
	// Process each position
	for i < n {
		pos := events[i].pos
		// Process all events at this position
		for i < n && events[i].pos == pos {
			coverage += events[i].delta
			i++
		}

		// Elements already at this position
		existing := freq[pos]

		// We can convert up to numOperations elements to this value
		// coverage includes both existing and convertible elements
		canConvert := coverage - existing
		if canConvert > numOperations {
			canConvert = numOperations
		}
		total := existing + canConvert
		if total > ans {
			ans = total
		}
	}

	return ans
}
