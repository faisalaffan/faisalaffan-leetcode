# 3607 — Power Grid Maintenance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func PowerGridMaintenance(intervals [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Merge Sort** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3607: Power Grid Maintenance
// https://leetcode.com/problems/power-grid-maintenance/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	intervals := [][]int{{1, 3}, {2, 5}, {4, 6}}
	fmt.Println("Test 1:", PowerGridMaintenance(intervals))
	// Test case 2
	intervals2 := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Test 2:", PowerGridMaintenance(intervals2))
	// Test case 3
	intervals3 := [][]int{{1, 10}}
	fmt.Println("Test 3:", PowerGridMaintenance(intervals3))
}

func PowerGridMaintenance(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
  // Custom sort dengan comparator
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// Merge intervals and count time needed
	time := 0
	currentEnd := 0
	for _, iv := range intervals {
		if iv[0] > currentEnd {
			time += iv[1] - iv[0]
		} else if iv[1] > currentEnd {
			time += iv[1] - currentEnd
		}
		if iv[1] > currentEnd {
			currentEnd = iv[1]
		}
	}
	return time
}
```
