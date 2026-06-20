# 1738 — Find Kth Largest Xor Coordinate Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func kthLargestValue(matrix [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(m * n * log(m*n)), Space: O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1738: Find Kth Largest XOR Coordinate Value
// https://leetcode.com/problems/find-kth-largest-xor-coordinate-value/
// Difficulty: Medium
// Time: O(m * n * log(m*n)), Space: O(m * n)

import (
	"fmt"
	"sort"
)

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
  // Membuat matriks/slice 2D untuk DP
	prefix := make([][]int, m)
	for i := 0; i < m; i++ {
		prefix[i] = make([]int, n)
	}

  // Alokasi slice integer
	values := make([]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := matrix[i][j]
			if i > 0 {
				val ^= prefix[i-1][j]
			}
			if j > 0 {
				val ^= prefix[i][j-1]
			}
			if i > 0 && j > 0 {
				val ^= prefix[i-1][j-1]
			}
			prefix[i][j] = val
			values = append(values, val)
		}
	}

  // Custom sort dengan comparator
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	return values[k-1]
}

func main() {
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 1)) // Expected: 7
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 2)) // Expected: 5
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 3)) // Expected: 4
}
```
