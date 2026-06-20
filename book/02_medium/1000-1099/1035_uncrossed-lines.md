# 1035 — Uncrossed Lines

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxUncrossedLines(nums1 []int, nums2 []int) int
```

> **💡 Hint:** DP - Longest Common Subsequence (LCS)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1035: Uncrossed Lines
// https://leetcode.com/problems/uncrossed-lines/
// Difficulty: Medium
//
// Approach: DP - Longest Common Subsequence (LCS)
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))    // 2
	fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2})) // 3
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[m][n]
}
```
