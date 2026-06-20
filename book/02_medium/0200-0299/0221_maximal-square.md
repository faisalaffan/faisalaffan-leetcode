# 0221 — Maximal Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximalSquare(matrix [][]byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m*n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #221: Maximal Square
// https://leetcode.com/problems/maximal-square/
// Difficulty: Medium
// Time: O(m*n), Space: O(n)

import "fmt"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
  // Alokasi slice integer
	dp := make([]int, cols+1)
	maxLen, prev := 0, 0

	for i := 0; i < rows; i++ {
		for j := 1; j <= cols; j++ {
			temp := dp[j]
			if matrix[i][j-1] == '1' {
				dp[j] = min(dp[j], min(dp[j-1], prev)) + 1
				if dp[j] > maxLen {
					maxLen = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}

	return maxLen * maxLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix1))

	matrix2 := [][]byte{{'0', '1'}, {'1', '0'}}
	fmt.Println(maximalSquare(matrix2))

	matrix3 := [][]byte{{'0'}}
	fmt.Println(maximalSquare(matrix3))
}
```
