# 2218 — Maximum Value Of K Coins From Piles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxValueOfCoins(piles [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2218: Maximum Value of K Coins from Piles
// https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/
// Difficulty: Hard
//
// DP: dp[k] = max value for k coins processed so far.
// For each pile, try taking 0..min(len(pile), k) coins from its top.
// Use prefix sums per pile for O(1) sum of top N coins.

import (
	"fmt"
)

func main() {
	// piles = [[1,100,3],[7,8,9]], k = 2 => 101
	piles := [][]int{{1, 100, 3}, {7, 8, 9}}
	fmt.Println(maxValueOfCoins(piles, 2))

	// Single pile
	fmt.Println(maxValueOfCoins([][]int{{5, 10, 15}}, 3))

	// Multiple piles
	fmt.Println(maxValueOfCoins([][]int{{100}, {200}, {300}}, 2))

	// Empty case
	fmt.Println(maxValueOfCoins([][]int{{}, {}}, 1))

	// Large k
	fmt.Println(maxValueOfCoins([][]int{{1, 2, 3, 4, 5}}, 3))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValueOfCoins(piles [][]int, k int) int {
	// Prefix sums for each pile.
  // Membuat matriks/slice 2D untuk DP
	pref := make([][]int, len(piles))
	for i, pile := range piles {
		pref[i] = make([]int, len(pile)+1)
		for j, v := range pile {
			pref[i][j+1] = pref[i][j] + v
		}
	}

	// dp[x] = max value using x coins from processed piles.
	// Use -1 as unreachable sentinel.
  // Alokasi slice integer
	dp := make([]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i] = -1
	}
	dp[0] = 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(piles); i++ {
  // Alokasi slice integer
		ndp := make([]int, k+1)
		copy(ndp, dp)
		maxTake := len(piles[i])
		if maxTake > k {
			maxTake = k
		}
		for take := 1; take <= maxTake; take++ {
			sum := pref[i][take]
			for used := take; used <= k; used++ {
				if dp[used-take] != -1 {
					ndp[used] = max(ndp[used], dp[used-take]+sum)
				}
			}
		}
		dp = ndp
	}

	if dp[k] == -1 {
		return 0
	}
	return dp[k]
}
```
