# 2365 — Task Scheduler Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func taskSchedulerII(tasks []int, space int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2365: Task Scheduler II
// https://leetcode.com/problems/task-scheduler-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Track last completion day for each task type. If within cooldown, advance day.

import "fmt"

func main() {
	fmt.Println(taskSchedulerII([]int{1, 2, 1, 2, 3, 1}, 3)) // 9
	fmt.Println(taskSchedulerII([]int{5, 8, 8, 5}, 2))       // 6
}

func taskSchedulerII(tasks []int, space int) int64 {
  // HashMap: O(1) lookup
	last := make(map[int]int64)
	var day int64 = 0
	for _, t := range tasks {
		day++
		if prev, ok := last[t]; ok && day-prev <= int64(space) {
			day = prev + int64(space) + 1
		}
		last[t] = day
	}
	return day
}
```
