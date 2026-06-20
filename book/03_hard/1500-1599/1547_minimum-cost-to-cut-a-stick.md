# 1547 — Minimum Cost To Cut A Stick

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minCost(n int, cuts []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1547: Minimum Cost to Cut a Stick
// https://leetcode.com/problems/minimum-cost-to-cut-a-stick/
// Difficulty: Hard
//
// DP interval approach:
// 1. Add 0 and n to the cuts array, sort.
// 2. DP[i][j] = minimum cost to cut stick from cuts[i] to cuts[j].
// 3. For each interval [i,j], try every cut point k in (i,j).
// 4. DP[i][j] = min(DP[i][k] + DP[k][j] + (cuts[j]-cuts[i]))
// 5. Return DP[0][len(cuts)-1].

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example: n=7, cuts=[1,3,4,5] -> 16
	fmt.Println(minCost(7, []int{1, 3, 4, 5}))

	// Additional tests
	fmt.Println(minCost(9, []int{5, 6, 1, 4, 2}))
	fmt.Println(minCost(10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(minCost(3, []int{1}))
	fmt.Println(minCost(4, []int{2}))
}

func minCost(n int, cuts []int) int {
	// Add boundaries and sort
  // Alokasi slice integer
	extended := make([]int, 0, len(cuts)+2)
	extended = append(extended, 0)
	extended = append(extended, cuts...)
	extended = append(extended, n)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(extended)

	m := len(extended)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, m)
	}

	// Interval DP: process by increasing length
	for length := 2; length < m; length++ {
		for i := 0; i+length < m; i++ {
			j := i + length
			dp[i][j] = math.MaxInt32
			for k := i + 1; k < j; k++ {
				cost := dp[i][k] + dp[k][j] + (extended[j] - extended[i])
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}

	return dp[0][m-1]
}
```
