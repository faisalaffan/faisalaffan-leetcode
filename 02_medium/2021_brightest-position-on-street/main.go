package main

// LeetCode #2021: Brightest Position on Street
// https://leetcode.com/problems/brightest-position-on-street/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func brightestPosition(lights [][]int) int {
	events := make([][2]int, 0, len(lights)*2)

	for _, l := range lights {
		pos, rng := l[0], l[1]
		events = append(events, [2]int{pos - rng, 1})
		events = append(events, [2]int{pos + rng + 1, -1})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	maxBrightness := 0
	currBrightness := 0
	bestPos := events[0][0]

	for _, e := range events {
		currBrightness += e[1]
		if currBrightness > maxBrightness {
			maxBrightness = currBrightness
			bestPos = e[0]
		}
	}

	return bestPos
}

func main() {
	// Test case 1
	lights1 := [][]int{{-3, 2}, {1, 2}, {3, 3}}
	fmt.Println("Test 1:", brightestPosition(lights1))
	// Expected: -1

	// Test case 2
	lights2 := [][]int{{1, 0}, {0, 1}}
	fmt.Println("Test 2:", brightestPosition(lights2))
	// Expected: 1

	// Test case 3
	lights3 := [][]int{{1, 2}}
	fmt.Println("Test 3:", brightestPosition(lights3))
	// Expected: -1
}
