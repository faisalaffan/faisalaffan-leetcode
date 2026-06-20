# 3623 — Count Number Of Trapezoids I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountNumberOfTrapezoidsI(points [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3623: Count Number of Trapezoids I
// https://leetcode.com/problems/count-number-of-trapezoids-i/
// Difficulty: Medium
// Complexity: O(n^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	points := [][]int{{0, 0}, {1, 0}, {2, 0}, {0, 1}, {1, 1}}
	fmt.Println("Test 1:", CountNumberOfTrapezoidsI(points))
	// Test case 2
	points2 := [][]int{{0, 0}, {1, 0}, {0, 1}}
	fmt.Println("Test 2:", CountNumberOfTrapezoidsI(points2))
	// Test case 3
	points3 := [][]int{{0, 0}, {1, 0}, {2, 0}, {0, 1}, {1, 1}, {2, 1}}
	fmt.Println("Test 3:", CountNumberOfTrapezoidsI(points3))
}

func CountNumberOfTrapezoidsI(points [][]int) int {
	// Count trapezoids: quadrilaterals with at least one pair of parallel sides
	n := len(points)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx1 := points[j][0] - points[i][0]
			dy1 := points[j][1] - points[i][1]
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					dx2 := points[l][0] - points[k][0]
					dy2 := points[l][1] - points[k][1]
					// Check if two sides are parallel (cross product = 0)
					if dx1*dy2 == dx2*dy1 {
						count++
					}
				}
			}
		}
	}
	return count
}
```
