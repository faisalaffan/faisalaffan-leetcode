# 2742 — Painting The Walls

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func paintWalls(cost []int, time []int) int
```

> **💡 Hint:** 0/1 Knapsack DP.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2742: Painting the Walls
// https://leetcode.com/problems/painting-the-walls/
// Difficulty: Hard
//
// Approach: 0/1 Knapsack DP.
// Each paid painter paints 1 wall (cost[i]) while the free painter
// paints time[i] walls simultaneously during that minute.
// So hiring a paid painter for wall i "covers" 1 + time[i] walls.
// dp[j] = minimum cost to cover at least j walls.
// Return dp[n].

import (
	"fmt"
	"math"
)

func main() {
	// Example 1: cost=[1,2,3,2], time=[1,2,3,2] -> 3
	fmt.Println(paintWalls([]int{1, 2, 3, 2}, []int{1, 2, 3, 2}))
	// Example 2: cost=[2,3,4,2], time=[1,1,1,1] -> 4
	fmt.Println(paintWalls([]int{2, 3, 4, 2}, []int{1, 1, 1, 1}))
}

func paintWalls(cost []int, time []int) int {
	n := len(cost)
  // Alokasi slice integer
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < n; i++ {
		coverage := 1 + time[i]
		c := cost[i]
		// Iterate backwards for 0/1 knapsack
		for j := n; j >= 0; j-- {
			if dp[j] == math.MaxInt32 {
				continue
			}
			next := j + coverage
			if next > n {
				next = n
			}
			if dp[j]+c < dp[next] {
				dp[next] = dp[j] + c
			}
		}
	}

	return dp[n]
}
```
