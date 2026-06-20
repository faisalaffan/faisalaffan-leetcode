# 0718 — Maximum Length Of Repeated Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findLength(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #718: Maximum Length of Repeated Subarray
// https://leetcode.com/problems/maximum-length-of-repeated-subarray/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	fmt.Println(findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))
}

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLen := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}

	return maxLen
}
```
