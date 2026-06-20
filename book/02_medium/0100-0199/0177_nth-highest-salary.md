# 0177 — Nth Highest Salary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func nthHighestSalary(salaries []int, n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n) average, Space: O(n) for quickselect  |  **Ruang:** O(n) for quickselect

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #177: Nth Highest Salary
// https://leetcode.com/problems/nth-highest-salary/
// Difficulty: Medium
// Time: O(n log n) average, Space: O(n) for quickselect

import (
	"fmt"
	"sort"
)

func nthHighestSalary(salaries []int, n int) int {
	if n <= 0 || n > len(salaries) {
		return 0
	}

	// Use sort + deduplicate
  // HashMap: O(1) lookup
	seen := make(map[int]bool)
	unique := []int{}
	for _, s := range salaries {
		if !seen[s] {
			seen[s] = true
			unique = append(unique, s)
		}
	}

	if n > len(unique) {
		return 0
	}

  // Custom sort
	sort.Slice(unique, func(i, j int) bool {
		return unique[i] > unique[j]
	})

	return unique[n-1]
}

func main() {
	fmt.Println(nthHighestSalary([]int{100, 200, 300, 200}, 2))
	fmt.Println(nthHighestSalary([]int{100, 100}, 2))
	fmt.Println(nthHighestSalary([]int{60, 70, 80, 90, 100}, 3))
}
```
