# 3450 — Maximum Students On A Single Bench

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaximumStudentsOnASingleBench(students [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
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
  // HashMap: O(1) lookup
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
```
