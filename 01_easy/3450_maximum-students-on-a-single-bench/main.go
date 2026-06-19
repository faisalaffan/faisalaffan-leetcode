package main

// LeetCode #3450: Maximum Students on a Single Bench
// https://leetcode.com/problems/maximum-students-on-a-single-bench/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(MaximumStudentsOnASingleBench([][]int{{1, 2}, {2, 3}, {1, 3}, {1, 2}}))
	fmt.Println(MaximumStudentsOnASingleBench([][]int{{1, 1}, {2, 1}, {3, 1}}))
}

// MaximumStudentsOnASingleBench returns the maximum number of different students on any single bench.
// Each entry is [student_id, bench_id].
// Time: O(n). Space: O(n).
func MaximumStudentsOnASingleBench(students [][]int) int {
	benchStudents := make(map[int]map[int]bool)
	for _, s := range students {
		studentID, benchID := s[0], s[1]
		if benchStudents[benchID] == nil {
			benchStudents[benchID] = make(map[int]bool)
		}
		benchStudents[benchID][studentID] = true
	}
	maxCount := 0
	for _, students := range benchStudents {
		if len(students) > maxCount {
			maxCount = len(students)
		}
	}
	return maxCount
}
