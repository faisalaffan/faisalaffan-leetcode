# 1112 — Highest Grade For Each Student

## Deskripsi

**Soal:** [1112. Highest Grade For Each Student](https://leetcode.com/problems/highest-grade-for-each-student/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(m) where m = unique students

**Algoritma:** —

> **Ide Kunci:** Track best grade (highest, then earliest course_id) per student

## Solusi Go

```go
package main

// LeetCode #1112: Highest Grade For Each Student
// https://leetcode.com/problems/highest-grade-for-each-student/
// Difficulty: Medium
//
// Approach: Track best grade (highest, then earliest course_id) per student
// Time: O(n)
// Space: O(m) where m = unique students

import "fmt"

func main() {
	// (student_id, course_id, grade)
	enrollments := [][]int{{1, 1, 90}, {1, 2, 95}, {2, 1, 85}, {2, 2, 85}, {3, 1, 70}}
	fmt.Println(highestGradeForEachStudent(enrollments))
}

func highestGradeForEachStudent(enrollments [][]int) [][]int {
	type best struct {
		grade    int
		courseID int
	}
  // Membuat map untuk pencarian O(1): key → value
	bestMap := make(map[int]best)

	for _, e := range enrollments {
		sid, cid, grade := e[0], e[1], e[2]
		if existing, ok := bestMap[sid]; !ok || grade > existing.grade || (grade == existing.grade && cid < existing.courseID) {
			bestMap[sid] = best{grade, cid}
		}
	}

  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0, len(bestMap))
	for sid, b := range bestMap {
		result = append(result, []int{sid, b.courseID, b.grade})
	}

	return result
}
```
