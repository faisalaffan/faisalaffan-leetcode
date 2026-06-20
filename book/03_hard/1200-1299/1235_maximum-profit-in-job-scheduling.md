# 1235 — Maximum Profit In Job Scheduling

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func jobScheduling(startTime []int, endTime []int, profit []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, DP, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1235: Maximum Profit in Job Scheduling
// https://leetcode.com/problems/maximum-profit-in-job-scheduling/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

type Job struct {
	start, end, profit int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([]Job, n)
	for i := 0; i < n; i++ {
		jobs[i] = Job{startTime[i], endTime[i], profit[i]}
	}

	// Sort by end time
  // Custom sort
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})

	// dp[i] = max profit considering first i jobs (i = number of jobs processed)
  // Alokasi slice
	dp := make([]int, n+1)
  // Alokasi slice
	endTimes := make([]int, n)
	for i := 0; i < n; i++ {
		endTimes[i] = jobs[i].end
	}

	for i := 1; i <= n; i++ {
		job := jobs[i-1]

		// Option 1: skip current job
		dp[i] = dp[i-1]

		// Option 2: take current job, find last non-overlapping job
		// Binary search for the last job with end <= job.start
		idx := sort.Search(n, func(j int) bool { return endTimes[j] > job.start })
		// idx is first ending after job.start, so idx previous jobs end <= job.start
		dp[i] = max(dp[i], dp[idx]+job.profit)
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println(jobScheduling(
		[]int{1, 2, 3, 3},
		[]int{3, 4, 5, 6},
		[]int{50, 10, 40, 70},
	)) // 120

	// Test case 2
	fmt.Println(jobScheduling(
		[]int{1, 2, 3, 4, 6},
		[]int{3, 5, 10, 6, 9},
		[]int{20, 20, 100, 70, 60},
	)) // 150

	// Test case 3: single job
	fmt.Println(jobScheduling(
		[]int{1},
		[]int{2},
		[]int{100},
	)) // 100
}
```
