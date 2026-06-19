package main

// LeetCode #406: Queue Reconstruction by Height
// https://leetcode.com/problems/queue-reconstruction-by-height/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// Sort by height descending, then by k ascending
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		}
		return people[i][1] < people[j][1]
	})

	result := make([][]int, 0, len(people))
	for _, p := range people {
		// Insert at index k
		k := p[1]
		result = append(result, nil)
		copy(result[k+1:], result[k:])
		result[k] = p
	}
	return result
}

func main() {
	// Test case 1
	p1 := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println("Test 1:", reconstructQueue(p1))
	// Expected: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]

	// Test case 2
	p2 := [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}
	fmt.Println("Test 2:", reconstructQueue(p2))
	// Expected: [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
}
