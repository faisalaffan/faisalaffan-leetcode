# 2323 — Find Minimum Time To Finish All Jobs Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumTime(jobs []int, workers []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2323: Find Minimum Time to Finish All Jobs II
// https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumTime(jobs []int, workers []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(jobs)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(workers)
	maxDays := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(jobs); i++ {
		days := (jobs[i] + workers[i] - 1) / workers[i] // ceil division
		if days > maxDays {
			maxDays = days
		}
	}
	return maxDays
}

func main() {
	// Test case 1
	fmt.Println(minimumTime([]int{5, 2, 4}, []int{1, 7, 5}))
	// Expected: 2

	// Test case 2
	fmt.Println(minimumTime([]int{3, 18, 30}, []int{3, 15, 5}))
	// Expected: 6
}
```
