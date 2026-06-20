# 1235 — Maximum Profit In Job Scheduling

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func jobScheduling(startTime []int, endTime []int, profit []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Custom sort dengan comparator
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})

	// dp[i] = max profit considering first i jobs (i = number of jobs processed)
  // Alokasi slice integer
	dp := make([]int, n+1)
  // Alokasi slice integer
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
