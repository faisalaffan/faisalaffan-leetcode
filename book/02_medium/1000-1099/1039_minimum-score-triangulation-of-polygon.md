# 1039 — Minimum Score Triangulation Of Polygon

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minScoreTriangulation(values []int) int
```

> **💡 Hint:** Interval DP. dp[i][j] = min score triangulating polygon from i to j.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** O(n^3)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1039: Minimum Score Triangulation of Polygon
// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/
// Difficulty: Medium
//
// Approach: Interval DP. dp[i][j] = min score triangulating polygon from i to j.
// Time: O(n^3)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(minScoreTriangulation([]int{1, 2, 3}))    // 6
	fmt.Println(minScoreTriangulation([]int{3, 7, 4, 5})) // 144
}

func minScoreTriangulation(values []int) int {
	n := len(values)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length < n; length++ {
		for i := 0; i+length < n; i++ {
			j := i + length
			dp[i][j] = 1<<31 - 1
			for k := i + 1; k < j; k++ {
				score := dp[i][k] + dp[k][j] + values[i]*values[j]*values[k]
				if score < dp[i][j] {
					dp[i][j] = score
				}
			}
		}
	}

	return dp[0][n-1]
}
```
