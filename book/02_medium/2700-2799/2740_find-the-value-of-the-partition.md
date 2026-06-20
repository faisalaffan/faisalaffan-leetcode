# 2740 — Find The Value Of The Partition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheValueOfThePartition(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2740: Find the Value of the Partition
// https://leetcode.com/problems/find-the-value-of-the-partition/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func FindTheValueOfThePartition(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	minDiff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if diff := nums[i] - nums[i-1]; diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func main() {
	fmt.Println(FindTheValueOfThePartition([]int{1, 3, 2, 4}))
	fmt.Println(FindTheValueOfThePartition([]int{100, 1, 10}))
}
```
