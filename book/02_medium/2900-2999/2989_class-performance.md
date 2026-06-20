# 2989 — Class Performance

## Deskripsi

**Soal:** [2989. Class Performance](https://leetcode.com/problems/class-performance/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func classPerformance(scores []Score) int`

## Solusi Go

```go
package main

// LeetCode #2989: Class Performance
// https://leetcode.com/problems/class-performance/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: Find the difference between the highest total score
// (sum of 3 assignments) and the lowest total score across all students.

import (
	"fmt"
)

// Score represents the Scores database table.
type Score struct {
	StudentID   int
	StudentName string
	Assignment1 int
	Assignment2 int
	Assignment3 int
}

// classPerformance computes max(total) - min(total) across all students.
// Time: O(n) | Space: O(1)
// n = number of students.
func classPerformance(scores []Score) int {
	if len(scores) == 0 {
		return 0
	}

	minTotal := 0
	maxTotal := 0

	for i, s := range scores {
		total := s.Assignment1 + s.Assignment2 + s.Assignment3
		if i == 0 {
			minTotal = total
			maxTotal = total
		} else {
			if total < minTotal {
				minTotal = total
			}
			if total > maxTotal {
				maxTotal = total
			}
		}
	}

	return maxTotal - minTotal
}

func main() {
	// Test data from the problem.
	scores := []Score{
		{StudentID: 309, StudentName: "Owen", Assignment1: 88, Assignment2: 47, Assignment3: 87},
		{StudentID: 321, StudentName: "Claire", Assignment1: 98, Assignment2: 95, Assignment3: 37},
		{StudentID: 338, StudentName: "Julian", Assignment1: 100, Assignment2: 64, Assignment3: 43},
		{StudentID: 423, StudentName: "Peyton", Assignment1: 60, Assignment2: 44, Assignment3: 47},
		{StudentID: 896, StudentName: "David", Assignment1: 32, Assignment2: 37, Assignment3: 50},
		{StudentID: 235, StudentName: "Camila", Assignment1: 31, Assignment2: 53, Assignment3: 69},
	}

	result := classPerformance(scores)

	fmt.Printf("Class Performance — difference in score: %d\n", result)
	// Expected: 111 (Claire total=230, David total=119)
}
```
