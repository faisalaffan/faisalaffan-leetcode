# 3262 — Find Overlapping Shifts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func countOverlappingShifts(shifts []Shift) [][2]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3262: Find Overlapping Shifts
// https://leetcode.com/problems/find-overlapping-shifts/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	shifts1 := []Shift{
		{1, 8, 12}, {1, 11, 15}, {1, 14, 18},
		{2, 9, 17}, {2, 16, 20},
		{3, 10, 12}, {3, 13, 15}, {3, 16, 18},
		{4, 8, 10}, {4, 9, 11},
	}
	fmt.Println(countOverlappingShifts(shifts1)) // [[1 2] [2 1] [4 1]]

	// Test case 2
	shifts2 := []Shift{{1, 1, 3}, {1, 2, 4}}
	fmt.Println(countOverlappingShifts(shifts2)) // [[1 1]]

	// Test case 3
	shifts3 := []Shift{{1, 1, 2}, {1, 3, 4}}
	fmt.Println(countOverlappingShifts(shifts3)) // []
}

type Shift struct {
	EmployeeID int
	StartTime  int
	EndTime    int
}

func countOverlappingShifts(shifts []Shift) [][2]int {
	// Group shifts by employee
  // Membuat map (HashMap) — pencarian O(1)
	empShifts := make(map[int][]Shift)
	for _, s := range shifts {
		empShifts[s.EmployeeID] = append(empShifts[s.EmployeeID], s)
	}

	type result struct {
		employeeID int
		count      int
	}
	var results []result

	for empID, s := range empShifts {
		// Sort shifts by start time
  // Custom sort dengan comparator
		sort.Slice(s, func(i, j int) bool {
			return s[i].StartTime < s[j].StartTime
		})

		count := 0
		// Sweep line: track maximum end time seen so far
		maxEnd := s[0].EndTime
		for i := 1; i < len(s); i++ {
			if s[i].StartTime < maxEnd {
				count++
			}
			if s[i].EndTime > maxEnd {
				maxEnd = s[i].EndTime
			}
		}

		if count > 0 {
			results = append(results, result{empID, count})
		}
	}

	// Sort by employee ID
  // Custom sort dengan comparator
	sort.Slice(results, func(i, j int) bool {
		return results[i].employeeID < results[j].employeeID
	})

  // Alokasi slice integer
	out := make([][2]int, len(results))
	for i, r := range results {
		out[i] = [2]int{r.employeeID, r.count}
	}
	return out
}
```
