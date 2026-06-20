# 2589 — Minimum Time To Complete All Tasks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findMinimumTime(tasks [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2589: Minimum Time to Complete All Tasks
// https://leetcode.com/problems/minimum-time-to-complete-all-tasks/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

// findMinimumTime finds min total time to complete all tasks.
// tasks[i] = [start, end, duration]. Sort by end, schedule greedily
// as late as possible in the window.
//
// Complexity: O(n * maxEnd) time, O(maxEnd) space
func findMinimumTime(tasks [][]int) int {
	// Sort by end time
  // Custom sort dengan comparator
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})

	maxEnd := 0
	for _, t := range tasks {
		if t[1] > maxEnd {
			maxEnd = t[1]
		}
	}

	used := make([]bool, maxEnd+1)

	for _, task := range tasks {
		start, end, duration := task[0], task[1], task[2]

		// Count already used slots in [start, end]
		alreadyUsed := 0
		for t := start; t <= end; t++ {
			if used[t] {
				alreadyUsed++
			}
		}

		// Schedule remaining from the right (latest time slots)
		remaining := duration - alreadyUsed
		for t := end; remaining > 0; t-- {
			if !used[t] {
				used[t] = true
				remaining--
			}
		}
	}

	total := 0
	for _, u := range used {
		if u {
			total++
		}
	}
	return total
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: ->", findMinimumTime([][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}})) // 2

	// Additional test cases
	fmt.Println("Test 2: ->", findMinimumTime([][]int{{1, 3, 2}, {2, 4, 1}})) // 3
	fmt.Println("Test 3: ->", findMinimumTime([][]int{{1, 2, 1}}))            // 1
	fmt.Println("Test 4: ->", findMinimumTime([][]int{{1, 5, 3}, {2, 4, 2}})) // 4
}
```
