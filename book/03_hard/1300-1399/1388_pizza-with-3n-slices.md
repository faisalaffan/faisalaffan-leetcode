# 1388 — Pizza With 3N Slices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSizeSlices(slices []int) int
```

> **💡 Hint:** DP for circular "House Robber" variant.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1388: Pizza With 3n Slices
// https://leetcode.com/problems/pizza-with-3n-slices/
// Difficulty: Hard
//
// Approach: DP for circular "House Robber" variant.
// From a circular array of 3n slices, pick n non-adjacent slices to maximize
// the sum. Equivalent to running linear House-Robber DP twice:
//   Case 1: exclude last element (allows picking first freely)
//   Case 2: exclude first element (allows picking last freely)
// dp[i][j] = max sum from first i elements picking j non-adjacent elements.

import "fmt"

func maxSizeSlices(slices []int) int {
	n := len(slices)
	k := n / 3

	maxPick := func(arr []int) int {
		m := len(arr)
  // Membuat matriks/slice 2D untuk DP
		dp := make([][]int, m+1)
		for i := 0; i <= m; i++ {
			dp[i] = make([]int, k+1)
		}
		for i := 1; i <= m; i++ {
			for j := 1; j <= k; j++ {
				// Skip arr[i-1]
				dp[i][j] = dp[i-1][j]
				// Take arr[i-1]: must skip adjacent, use dp[i-2][j-1]
				if i >= 2 {
					v := dp[i-2][j-1] + arr[i-1]
					if v > dp[i][j] {
						dp[i][j] = v
					}
				} else if j == 1 {
					if arr[i-1] > dp[i][j] {
						dp[i][j] = arr[i-1]
					}
				}
			}
		}
		return dp[m][k]
	}

	// Case 1: exclude last
	best := maxPick(slices[:n-1])
	// Case 2: exclude first
	if v := maxPick(slices[1:]); v > best {
		best = v
	}
	return best
}

func main() {
	fmt.Println(maxSizeSlices([]int{1, 2, 3, 4, 5, 6}))       // 10
	fmt.Println(maxSizeSlices([]int{8, 9, 8, 6, 1, 1}))       // 16
	fmt.Println(maxSizeSlices([]int{2, 4, 3, 5, 6, 7, 8, 9, 9})) // 23
}
```
