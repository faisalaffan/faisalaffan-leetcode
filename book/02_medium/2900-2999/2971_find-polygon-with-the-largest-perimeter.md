# 2971 — Find Polygon With The Largest Perimeter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func largestPerimeter(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2971: Find Polygon With the Largest Perimeter
// https://leetcode.com/problems/find-polygon-with-the-largest-perimeter/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1, 10}))
	fmt.Println(largestPerimeter([]int{1, 2, 3}))
}

func largestPerimeter(nums []int) int64 {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	s := int64(0)
	for _, x := range nums {
		s += int64(x)
	}
	for i := len(nums) - 1; i >= 2; i-- {
		if int64(nums[i]) < s-int64(nums[i]) {
			return s
		}
		s -= int64(nums[i])
	}
	return -1
}
```
