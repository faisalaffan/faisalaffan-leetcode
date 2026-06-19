package main

// LeetCode #3182: Find Top Scoring Students
// https://leetcode.com/problems/find-top-scoring-students/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findTopScoringStudents(scores [][]int, threshold int) []int {
	type student struct {
		id    int
		total int
	}

	var list []student
	for _, s := range scores {
		list = append(list, student{s[0], s[1]})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].total != list[j].total {
			return list[i].total > list[j].total
		}
		return list[i].id < list[j].id
	})

	ans := make([]int, 0)
	for _, s := range list {
		if s.total >= threshold {
			ans = append(ans, s.id)
		}
	}
	return ans
}

func main() {
	fmt.Println(findTopScoringStudents([][]int{{1, 95}, {2, 85}, {3, 90}}, 90)) // Expected: [1 3]
	fmt.Println(findTopScoringStudents([][]int{{1, 70}, {2, 65}}, 80))          // Expected: []
}
