# 3699 — Number Of Zigzag Arrays I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func zigZagArraysI(n int, l int, r int) int
```

> **💡 Hint:** DP over possible values with three-state tracking.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3699: Number of ZigZag Arrays I
// https://leetcode.com/problems/number-of-zigzag-arrays-i/
// Difficulty: Hard
//
// Count arrays of length n with values in [l, r] such that:
// 1. No adjacent elements equal
// 2. No three consecutive elements are strictly increasing or decreasing
//
// Approach: DP over possible values with three-state tracking.

import "fmt"

func main() {
	// Example 1
	fmt.Println(zigZagArraysI(3, 4, 5))
	// Example 2
	fmt.Println(zigZagArraysI(4, 1, 3))
	// Edge: n = 3, small range
	fmt.Println(zigZagArraysI(3, 1, 2))
}

const MOD = 1000000007

func zigZagArraysI(n int, l int, r int) int {
	if n < 3 || l > r {
		return 0
	}
	m := r - l + 1
	if m < 2 {
		return 0
	}

	// dp[last][state]
	// state 0: last > second-last (increasing at end)
	// state 1: last < second-last (decreasing at end)
	// state 2: equal doesn't happen (no adjacent equal)
	// Actually since no adjacent equal, we only track up/down

	// For position i, we track counts for each possible value
  // Alokasi slice integer
	dp := make([][2]int, m)
	for v := 0; v < m; v++ {
		dp[v][0] = 1 // increasing (single element)
		dp[v][1] = 1 // decreasing (single element)
	}

	for pos := 2; pos <= n; pos++ {
  // Alokasi slice integer
		ndp := make([][2]int, m)
		for cur := 0; cur < m; cur++ {
			// For arrays where cur is at an odd position (1-indexed from end):
			// array ends with cur, and cur should be a peak or valley
			// We need to consider previous values that are different from cur

			// Previous was less than cur: cur is at a peak
			for prev := 0; prev < cur; prev++ {
				ndp[cur][0] = (ndp[cur][0] + dp[prev][1]) % MOD
			}
			// Previous was greater than cur: cur is at a valley
			for prev := cur + 1; prev < m; prev++ {
				ndp[cur][1] = (ndp[cur][1] + dp[prev][0]) % MOD
			}
		}
		dp = ndp
	}

	result := 0
	for v := 0; v < m; v++ {
		result = (result + dp[v][0] + dp[v][1]) % MOD
	}
	return result
}
```
