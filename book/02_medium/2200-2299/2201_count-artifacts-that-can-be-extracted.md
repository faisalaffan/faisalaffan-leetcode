# 2201 — Count Artifacts That Can Be Extracted

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func digArtifacts(n int, artifacts [][]int, dig [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2201: Count Artifacts That Can Be Extracted
// https://leetcode.com/problems/count-artifacts-that-can-be-extracted/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
  // Membuat matriks/slice 2D untuk DP
	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)
	}
	for _, d := range dig {
		grid[d[0]][d[1]] = true
	}

	count := 0
	for _, art := range artifacts {
		r1, c1, r2, c2 := art[0], art[1], art[2], art[3]
		extracted := true
		for r := r1; r <= r2 && extracted; r++ {
			for c := c1; c <= c2; c++ {
				if !grid[r][c] {
					extracted = false
					break
				}
			}
		}
		if extracted {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(digArtifacts(2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}}))
	// Expected: 1

	// Test case 2
	fmt.Println(digArtifacts(2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}, {1, 1}}))
	// Expected: 2
}
```
