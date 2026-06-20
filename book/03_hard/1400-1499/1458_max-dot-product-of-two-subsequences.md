# 1458 — Max Dot Product Of Two Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxDotProduct(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1458: Max Dot Product of Two Subsequences
// https://leetcode.com/problems/max-dot-product-of-two-subsequences/
// Difficulty: Hard

import "fmt"

func maxDotProduct(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n2)
	}

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			product := nums1[i] * nums2[j]
			dp[i][j] = product
			if i > 0 && dp[i-1][j] > dp[i][j] {
				dp[i][j] = dp[i-1][j]
			}
			if j > 0 && dp[i][j-1] > dp[i][j] {
				dp[i][j] = dp[i][j-1]
			}
			if i > 0 && j > 0 {
				candidate := dp[i-1][j-1]
				if candidate > 0 {
					candidate += product
				} else {
					candidate = product
				}
				if candidate > dp[i][j] {
					dp[i][j] = candidate
				}
			}
		}
	}
	return dp[n1-1][n2-1]
}

func main() {
	// Example: [2,1,-2,5], [3,0,-6] -> 18
	fmt.Println(maxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6}))
}
```
