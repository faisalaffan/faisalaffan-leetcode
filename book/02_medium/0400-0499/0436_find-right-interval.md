# 0436 — Find Right Interval

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findRightInterval(intervals [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #436: Find Right Interval
// https://leetcode.com/problems/find-right-interval/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	// Create array of (start, index) pairs
  // Alokasi slice integer
	starts := make([][2]int, n)
	for i, iv := range intervals {
		starts[i] = [2]int{iv[0], i}
	}
  // Custom sort dengan comparator
	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})

  // Alokasi slice integer
	result := make([]int, n)
	for i, iv := range intervals {
		target := iv[1]
		// Binary search
		idx := sort.Search(n, func(j int) bool {
			return starts[j][0] >= target
		})
		if idx < n {
			result[i] = starts[idx][1]
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findRightInterval([][]int{{1, 2}}))
	// Expected: [-1]

	// Test case 2
	fmt.Println("Test 2:", findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
	// Expected: [-1, 0, 1]

	// Test case 3
	fmt.Println("Test 3:", findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}}))
	// Expected: [-1, 2, -1]
}
```
