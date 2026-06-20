# 3429 — Paint House Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minCost3429(n int, cost [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3429: Paint House IV
// https://leetcode.com/problems/paint-house-iv/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import (
	"fmt"
	"math"
)

func minCost3429(n int, cost [][]int) int64 {
	half := n / 2
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]int64, half+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][]int64, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = make([]int64, 3)
			for k := 0; k < 3; k++ {
				dp[0][j][k] = 0
			}
		}
	}

	for i := 0; i < half; i++ {
		left := i
		right := n - 1 - i
		for preJ := 0; preJ < 3; preJ++ {
			for preK := 0; preK < 3; preK++ {
				best := int64(math.MaxInt64)
				for j := 0; j < 3; j++ {
					if j == preJ {
						continue
					}
					for k := 0; k < 3; k++ {
						if k == preK || k == j {
							continue
						}
						val := dp[i][j][k] + int64(cost[left][j]) + int64(cost[right][k])
						if val < best {
							best = val
						}
					}
				}
				dp[i+1][preJ][preK] = best
			}
		}
	}

	ans := int64(math.MaxInt64)
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			if dp[half][j][k] < ans {
				ans = dp[half][j][k]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minCost3429(4, [][]int{{3, 5, 7}, {6, 2, 9}, {4, 8, 1}, {7, 3, 5}})) // 9
	fmt.Println(minCost3429(2, [][]int{{1, 2, 3}, {4, 5, 6}})) // 7
}
```
