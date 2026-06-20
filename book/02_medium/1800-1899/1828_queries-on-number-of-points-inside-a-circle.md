# 1828 — Queries On Number Of Points Inside A Circle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPoints(points [][]int, queries [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * q), Space: O(q)  
**Kompleksitas Ruang:** O(q)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1828: Queries on Number of Points Inside a Circle
// https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/
// Difficulty: Medium
// Time: O(n * q), Space: O(q)

import "fmt"

func countPoints(points [][]int, queries [][]int) []int {
  // Alokasi slice integer
	result := make([]int, len(queries))

	for i, q := range queries {
		cx, cy, r := q[0], q[1], q[2]
		r2 := r * r
		count := 0
		for _, p := range points {
			dx := p[0] - cx
			dy := p[1] - cy
			if dx*dx+dy*dy <= r2 {
				count++
			}
		}
		result[i] = count
	}
	return result
}

func main() {
	fmt.Println(countPoints([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}))
	// Expected: [3, 2, 2]

	fmt.Println(countPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}))
	// Expected: [2, 3, 2, 3]
}
```
