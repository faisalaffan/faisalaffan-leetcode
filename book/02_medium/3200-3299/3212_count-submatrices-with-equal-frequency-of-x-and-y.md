# 3212 — Count Submatrices With Equal Frequency Of X And Y

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSubmatrices(grid [][]byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3212: Count Submatrices With Equal Frequency of X and Y
// https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func numberOfSubmatrices(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Membuat matriks/slice 2D untuk DP
	prefX := make([][]int, m+1)
  // Membuat matriks/slice 2D untuk DP
	prefY := make([][]int, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range prefX {
		prefX[i] = make([]int, n+1)
		prefY[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefX[i+1][j+1] = prefX[i][j+1] + prefX[i+1][j] - prefX[i][j]
			prefY[i+1][j+1] = prefY[i][j+1] + prefY[i+1][j] - prefY[i][j]
			if grid[i][j] == 'X' {
				prefX[i+1][j+1]++
			} else if grid[i][j] == 'Y' {
				prefY[i+1][j+1]++
			}
		}
	}

	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			x := prefX[i][j]
			y := prefY[i][j]
			if x > 0 && x == y {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}})) // Expected: 3
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'X'}, {'Y', 'Y'}}))          // Expected: 0
}
```
