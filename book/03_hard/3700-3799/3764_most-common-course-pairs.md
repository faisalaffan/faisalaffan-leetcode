# 3764 — Most Common Course Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func findCommonPairs(completions []CourseCompletion, minGrade string) []CoursePair
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3764: Most Common Course Pairs
// https://leetcode.com/problems/most-common-course-pairs/
// Difficulty: Hard (SQL / Database)
//
// Find the most common consecutive course pairs among high achievers.
//
// Go port: simulate SQL logic with data structures.

import (
	"fmt"
	"sort"
)

type CourseCompletion struct {
	UserID   int
	CourseID int
	Term     int
	Grade    string
}

type CoursePair struct {
	CourseA int
	CourseB int
	Count   int
}

func main() {
	// Example
	completions := []CourseCompletion{
		{1, 101, 1, "A"},
		{1, 102, 1, "A"},
		{2, 101, 1, "A"},
		{2, 102, 1, "B"},
		{2, 103, 2, "A"},
	}
	fmt.Println(findCommonPairs(completions, "A"))
	// Edge: empty
	fmt.Println(findCommonPairs([]CourseCompletion{}, "B"))
}

func findCommonPairs(completions []CourseCompletion, minGrade string) []CoursePair {
	// Filter to high achievers (grade >= minGrade)
	// Group by user, order by term, find consecutive pairs
  // Membuat map (HashMap) — pencarian O(1)
	userCourses := make(map[int][]int)

	for _, c := range completions {
		if c.Grade >= minGrade {
			userCourses[c.UserID] = append(userCourses[c.UserID], c.CourseID)
		}
	}

	// Sort each user's courses by term (assumes input order is term order)
	// Count pairs
  // Membuat map (HashMap) — pencarian O(1)
	pairCount := make(map[[2]int]int)

	for _, courses := range userCourses {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(courses) // sort by course ID to establish canonical order
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(courses)-1; i++ {
			a, b := courses[i], courses[i+1]
			if a > b {
				a, b = b, a
			}
			key := [2]int{a, b}
			pairCount[key]++
		}
	}

	var result []CoursePair
	maxCount := 0
	for pair, cnt := range pairCount {
		if cnt > maxCount {
			maxCount = cnt
			result = []CoursePair{{pair[0], pair[1], cnt}}
		} else if cnt == maxCount {
			result = append(result, CoursePair{pair[0], pair[1], cnt})
		}
	}

  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		if result[i].CourseA != result[j].CourseA {
			return result[i].CourseA < result[j].CourseA
		}
		return result[i].CourseB < result[j].CourseB
	})

	return result
}
```
