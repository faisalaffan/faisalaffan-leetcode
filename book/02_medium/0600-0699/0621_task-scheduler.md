# 0621 — Task Scheduler

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func LeastInterval(tasks []byte, n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1) (26 letters)


## 💻 Solusi Go

```go
package main

// LeetCode #621: Task Scheduler
// https://leetcode.com/problems/task-scheduler/
// Difficulty: Medium
// Time: O(n)
// Space: O(1) (26 letters)

import "fmt"

func main() {
	fmt.Println(LeastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(LeastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	fmt.Println(LeastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
}

func LeastInterval(tasks []byte, n int) int {
  // Alokasi slice
	counts := make([]int, 26)
	maxFreq := 0
	for _, t := range tasks {
		counts[t-'A']++
		if counts[t-'A'] > maxFreq {
			maxFreq = counts[t-'A']
		}
	}

	maxCount := 0
	for _, c := range counts {
		if c == maxFreq {
			maxCount++
		}
	}

	partLen := maxFreq - 1
	emptySlots := partLen * (n - (maxCount - 1))
	availableTasks := len(tasks) - maxFreq*maxCount
	idles := 0
	if emptySlots > availableTasks {
		idles = emptySlots - availableTasks
	}

	return len(tasks) + idles
}
```
