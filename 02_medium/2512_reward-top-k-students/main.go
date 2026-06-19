package main

// LeetCode #2512: Reward Top K Students
// https://leetcode.com/problems/reward-top-k-students/
// Difficulty: Medium
// Time: O(n * L + n log n) | Space: O(n)
// Score each student by positive/negative keywords, sort, return top K IDs.

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(topStudents([]string{"smart", "brilliant", "studious"}, []string{"not"}, [][]string{
		{"this", "student", "is", "studious"},
		{"the", "student", "is", "smart"},
	}, []int{1, 2}, 2))
	// [2, 1]

	fmt.Println(topStudents([]string{"smart"}, []string{"boring"}, [][]string{
		{"this", "is", "smart"},
		{"this", "is", "boring"},
	}, []int{1, 2}, 1))
	// [1]
}

func topStudents(positiveFeedback []string, negativeFeedback []string, report [][]string, studentID []int, k int) []int {
	pos := make(map[string]bool)
	neg := make(map[string]bool)
	for _, w := range positiveFeedback {
		pos[w] = true
	}
	for _, w := range negativeFeedback {
		neg[w] = true
	}

	type student struct {
		id, score int
	}
	students := make([]student, len(studentID))

	for i, id := range studentID {
		score := 0
		for _, w := range report[i] {
			w = strings.ToLower(w)
			if pos[w] {
				score += 3
			} else if neg[w] {
				score -= 1
			}
		}
		students[i] = student{id, score}
	}

	sort.Slice(students, func(i, j int) bool {
		if students[i].score != students[j].score {
			return students[i].score > students[j].score
		}
		return students[i].id < students[j].id
	})

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = students[i].id
	}
	return ans
}
