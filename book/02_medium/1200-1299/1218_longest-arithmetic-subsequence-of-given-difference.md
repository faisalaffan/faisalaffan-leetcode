# 1218 — Longest Arithmetic Subsequence Of Given Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestSubsequence(arr []int, difference int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1218: Longest Arithmetic Subsequence of Given Difference
// https://leetcode.com/problems/longest-arithmetic-subsequence-of-given-difference/
// Difficulty: Medium

// dp[x] = length of longest arithmetic subsequence ending with value x.
// dp[x] = dp[x-difference] + 1

// Time: O(n)
// Space: O(n)

func longestSubsequence(arr []int, difference int) int {
  // Membuat map (HashMap) — pencarian O(1)
	dp := make(map[int]int)
	maxLen := 0

	for _, v := range arr {
		prev := v - difference
		if count, exists := dp[prev]; exists {
			dp[v] = count + 1
		} else {
			dp[v] = 1
		}
		if dp[v] > maxLen {
			maxLen = dp[v]
		}
	}

	return maxLen
}

func main() {
	fmt.Printf("%d (expected: 4)\n", longestSubsequence([]int{1, 2, 3, 4}, 1))
	fmt.Printf("%d (expected: 1)\n", longestSubsequence([]int{1, 3, 5, 7}, 1))
	fmt.Printf("%d (expected: 4)\n", longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}
```
