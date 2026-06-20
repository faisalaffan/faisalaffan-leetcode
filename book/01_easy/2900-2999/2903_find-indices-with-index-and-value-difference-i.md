# 2903 — Find Indices With Index And Value Difference I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindIndicesWithIndexAndValueDifferenceI(nums []int, indexDifference int, valueDifference int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2903: Find Indices With Index and Value Difference I
// https://leetcode.com/problems/find-indices-with-index-and-value-difference-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findIndices
	fmt.Println(FindIndicesWithIndexAndValueDifferenceI([]int{5, 1, 4, 1}, 2, 4)) // [0,3]
	fmt.Println(FindIndicesWithIndexAndValueDifferenceI([]int{2, 1}, 0, 0))       // [0,0]
	fmt.Println(FindIndicesWithIndexAndValueDifferenceI([]int{1, 2, 3}, 2, 4))    // [-1,-1]
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: findIndices
func FindIndicesWithIndexAndValueDifferenceI(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if abs(i-j) >= indexDifference && abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
