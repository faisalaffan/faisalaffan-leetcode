# 0120 — Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumTotal(triangle [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #120: Triangle
// https://leetcode.com/problems/triangle/
// Difficulty: Medium

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
  // Alokasi slice integer
	dp := make([]int, n)
	copy(dp, triangle[n-1])

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if dp[j] < dp[j+1] {
				dp[j] = triangle[i][j] + dp[j]
			} else {
				dp[j] = triangle[i][j] + dp[j+1]
			}
		}
	}

	return dp[0]
}

func main() {
	// Test case 1
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})) // 11

	// Test case 2
	fmt.Println(minimumTotal([][]int{{-10}})) // -10

	// Test case 3
	fmt.Println(minimumTotal([][]int{{1}, {2, 3}})) // 3
}

// Time: O(n^2) | Space: O(n)
```
