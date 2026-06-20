# 0474 — Ones And Zeroes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func OnesAndZeroes(strs []string, m int, n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m * n * len(strs))  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #474: Ones and Zeroes
// https://leetcode.com/problems/ones-and-zeroes/
// Difficulty: Medium
// Time: O(m * n * len(strs))
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(OnesAndZeroes([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(OnesAndZeroes([]string{"10", "0", "1"}, 1, 1))
}

func OnesAndZeroes(strs []string, m int, n int) int {
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros, ones := countBits(s)
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				if dp[i-zeros][j-ones]+1 > dp[i][j] {
					dp[i][j] = dp[i-zeros][j-ones] + 1
				}
			}
		}
	}

	return dp[m][n]
}

func countBits(s string) (int, int) {
	zeros, ones := 0, 0
	for _, c := range s {
		if c == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}
```
