# 3188 — Find Top Scoring Students Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findTopScoringStudentsIi(enrollments [][]int, threshold int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3188: Find Top Scoring Students II
// https://leetcode.com/problems/find-top-scoring-students-ii/
// Difficulty: Hard [Paid]
//
// Given student enrollment and score data, find students whose score in each
// of their enrolled courses meets or exceeds a given threshold.
//
// Input:
//   enrollments = [[student_id, course_id, score], ...]
//   threshold   = minimum score required in each enrolled course
//
// Output:
//   list of student_ids (sorted) who meet the threshold in ALL their courses.

import (
	"fmt"
	"sort"
)

func findTopScoringStudentsIi(enrollments [][]int, threshold int) []int {
	// Map: student_id -> map of course_id -> max score
	type courseScore struct {
		best  int // best score for this course
		valid bool
	}
  // Membuat map (HashMap) — pencarian O(1)
	studentCourses := make(map[int]map[int]*courseScore)

	for _, e := range enrollments {
		sid, cid, score := e[0], e[1], e[2]
		if studentCourses[sid] == nil {
			studentCourses[sid] = make(map[int]*courseScore)
		}
		if studentCourses[sid][cid] == nil {
			studentCourses[sid][cid] = &courseScore{}
		}
		cs := studentCourses[sid][cid]
		if score > cs.best {
			cs.best = score
		}
		if score >= threshold {
			cs.valid = true
		}
	}

  // Alokasi slice integer
	ans := make([]int, 0)
	for sid, courses := range studentCourses {
		allValid := true
		for _, cs := range courses {
			if !cs.valid || cs.best < threshold {
				allValid = false
				break
			}
		}
		if allValid {
			ans = append(ans, sid)
		}
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(ans)
	return ans
}

func main() {
	enrollments := [][]int{
		{1, 101, 95},
		{1, 102, 100},
		{2, 101, 100},
		{2, 102, 90},
		{3, 101, 100},
	}
	fmt.Println(findTopScoringStudentsIi(enrollments, 100))
	// Only student 2 has >= 100 in all courses? Student 2 has 100 in 101 and 90 in 102 -> no.
	// Student 3 has 100 in 101 -> need to check other courses... student 3 only has one course.
	// Actually, we need to check threshold against score directly.
	// Student 1: 95,100 -> 95 < 100 -> fail
	// Student 2: 100,90 -> 90 < 100 -> fail
	// Student 3: 100 -> >=100 -> pass
	// Output: [3]
}
```
