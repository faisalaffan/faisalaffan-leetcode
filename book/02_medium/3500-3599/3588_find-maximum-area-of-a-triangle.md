# 3588 — Find Maximum Area Of A Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindMaximumAreaOfATriangle(points [][]int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3588: Find Maximum Area of a Triangle
// https://leetcode.com/problems/find-maximum-area-of-a-triangle/
// Difficulty: Medium
// Complexity: O(n^3) time, O(1) space

import (
	"fmt"
	"math"
)

func main() {
	// Test case 1
	points := [][]int{{0, 0}, {1, 0}, {0, 1}}
	fmt.Println("Test 1:", FindMaximumAreaOfATriangle(points))
	// Test case 2
	points2 := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	fmt.Println("Test 2:", FindMaximumAreaOfATriangle(points2))
	// Test case 3
	points3 := [][]int{{0, 0}}
	fmt.Println("Test 3:", FindMaximumAreaOfATriangle(points3))
}

func FindMaximumAreaOfATriangle(points [][]int) float64 {
	n := len(points)
	maxArea := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				area := math.Abs(float64(
					points[i][0]*(points[j][1]-points[k][1]) +
						points[j][0]*(points[k][1]-points[i][1]) +
						points[k][0]*(points[i][1]-points[j][1]),
				)) / 2.0
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
```
