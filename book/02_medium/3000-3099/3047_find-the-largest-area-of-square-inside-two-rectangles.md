# 3047 — Find The Largest Area Of Square Inside Two Rectangles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3047: Find the Largest Area of Square Inside Two Rectangles
// https://leetcode.com/problems/find-the-largest-area-of-square-inside-two-rectangles/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(largestSquareArea([][]int{{1, 1}, {2, 2}, {3, 3}}, [][]int{{2, 2}, {3, 3}, {4, 4}}))
	fmt.Println(largestSquareArea([][]int{{1, 1}, {2, 2}, {1, 2}}, [][]int{{3, 3}, {4, 4}, {3, 4}}))
}

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	ans := int64(0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x1 := max(bottomLeft[i][0], bottomLeft[j][0])
			y1 := max(bottomLeft[i][1], bottomLeft[j][1])
			x2 := min(topRight[i][0], topRight[j][0])
			y2 := min(topRight[i][1], topRight[j][1])
			if x1 < x2 && y1 < y2 {
				side := min(x2-x1, y2-y1)
				area := int64(side) * int64(side)
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
