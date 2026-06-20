# 1478 — Allocate Mailboxes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minDistance(houses []int, k int) int
```

> **💡 Hint:** DP + Median Cost

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1478: Allocate Mailboxes
// https://leetcode.com/problems/allocate-mailboxes/
// Difficulty: Hard
//
// Approach: DP + Median Cost
// Sort houses first. dp[i][j] = min distance to place j mailboxes
// among first i houses (0-indexed).
// cost[l][r] = min total distance to serve houses[l..r] with 1 mailbox
// placed at the median (optimal for minimizing sum of absolute distances).
// Transition: dp[i][j] = min over p < i of dp[p][j-1] + cost[p+1][i].

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minDistance([]int{1, 4, 8, 10, 20}, 3))
	// Expected: 5

	fmt.Println(minDistance([]int{2, 3, 5, 12, 18}, 2))
	// Expected: 9

	fmt.Println(minDistance([]int{7, 4, 6, 1}, 1))
	// Expected: 8
}

func minDistance(houses []int, k int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(houses)
	n := len(houses)

	if k >= n {
		return 0
	}

	// precompute cost[i][j] = min dist for 1 mailbox serving houses[i..j]
  // Membuat matriks/slice 2D untuk DP
	cost := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range cost {
		cost[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// Median index
			mid := i + (j-i)/2
			median := houses[mid]
			total := 0
			for t := i; t <= j; t++ {
				total += absInt(houses[t] - median)
			}
			cost[i][j] = total
		}
	}

	// dp[i][j] = min distance for first i+1 houses with j+1 mailboxes
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, k)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// base: 1 mailbox
	for i := 0; i < n; i++ {
		dp[i][0] = cost[0][i]
	}

	// fill dp
	for j := 1; j < k; j++ {
		for i := j; i < n; i++ {
			for p := j - 1; p < i; p++ {
				dp[i][j] = minInt(dp[i][j], dp[p][j-1]+cost[p+1][i])
			}
		}
	}

	return dp[n-1][k-1]
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
