# 1285 — Find The Start And End Number Of Continuous Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findContinuousRanges(nums []int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1285: Find the Start and End Number of Continuous Ranges
// https://leetcode.com/problems/find-the-start-and-end-number-of-continuous-ranges/
// Difficulty: Medium [Paid]

// Given a sorted list of unique integers, find continuous ranges.

// Time: O(n)
// Space: O(n)

func findContinuousRanges(nums []int) [][]int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return [][]int{}
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)
	start := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			result = append(result, []int{start, nums[i-1]})
			start = nums[i]
		}
	}
	result = append(result, []int{start, nums[len(nums)-1]})

  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func main() {
	fmt.Printf("%v (expected: [[1 3] [6 7] [9 9]])\n",
		findContinuousRanges([]int{1, 2, 3, 6, 7, 9}))

	fmt.Printf("%v (expected: [[1 1]])\n",
		findContinuousRanges([]int{1}))

	fmt.Printf("%v (expected: [])\n",
		findContinuousRanges([]int{}))
}
```
