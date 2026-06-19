package main

// LeetCode #1988: Find Cutoff Score for Each School
// https://leetcode.com/problems/find-cutoff-score-for-each-school/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: For each school, find the minimum exam score such that
// the school's capacity >= student_count (school can accept all students
// from that score group). If no such score exists, output -1.
// SQL equivalent: LEFT JOIN Exam ON Schools.capacity >= Exam.student_count
// then MIN(score) per school.

import (
	"fmt"
	"sort"
)

// School represents the Schools database table.
type School struct {
	SchoolID int
	Capacity int
}

// Exam represents the Exam database table.
type Exam struct {
	Score        int
	StudentCount int
}

// SchoolResult holds the output.
type SchoolResult struct {
	SchoolID int
	Score    int
}

// findCutoffScore simulates the SQL query.
// Time: O(s * e) | Space: O(s)
// s = number of schools, e = number of exam rows.
// SQL equivalent:
//   SELECT school_id, MIN(IFNULL(score, -1)) AS score
//   FROM Schools LEFT JOIN Exam ON Schools.capacity >= Exam.student_count
//   GROUP BY school_id
func findCutoffScore(schools []School, exams []Exam) []SchoolResult {
	var results []SchoolResult

	for _, school := range schools {
		bestScore := -1
		for _, exam := range exams {
			if school.Capacity >= exam.StudentCount {
				if bestScore == -1 || exam.Score < bestScore {
					bestScore = exam.Score
				}
			}
		}
		results = append(results, SchoolResult{SchoolID: school.SchoolID, Score: bestScore})
	}

	// Order by school_id.
	sort.Slice(results, func(i, j int) bool {
		return results[i].SchoolID < results[j].SchoolID
	})

	return results
}

func main() {
	// Test data from the problem.
	schools := []School{
		{SchoolID: 5, Capacity: 48},
		{SchoolID: 9, Capacity: 9},
		{SchoolID: 10, Capacity: 99},
		{SchoolID: 11, Capacity: 151},
	}

	exams := []Exam{
		{Score: 975, StudentCount: 10},
		{Score: 966, StudentCount: 60},
		{Score: 844, StudentCount: 76},
		{Score: 749, StudentCount: 76},
		{Score: 744, StudentCount: 100},
	}

	results := findCutoffScore(schools, exams)

	fmt.Println("Cutoff Score for Each School (school_id | score):")
	for _, r := range results {
		fmt.Printf("%d | %d\n", r.SchoolID, r.Score)
	}
	// Expected output:
	// 5 | 975  (capacity 48, only 975(10) qualifies; 966 has 60 > 48)
	// 9 | -1   (capacity 9, minimum student_count is 10 > 9)
	// 10 | 749 (capacity 99, 749(76) qualifies, 744(100) > 99)
	// 11 | 744 (capacity 151, all scores qualify; minimum is 744)
}
