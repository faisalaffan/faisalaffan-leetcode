# 2901 — Longest Unequal Adjacent Groups Subsequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestUnequalAdjacentGroupsSubsequenceIi(words []string, groups []int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2901: Longest Unequal Adjacent Groups Subsequence II
// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func LongestUnequalAdjacentGroupsSubsequenceIi(words []string, groups []int) []string {
	n := len(words)
  // Alokasi slice integer
	dp := make([]int, n)
  // Alokasi slice integer
	prev := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range prev {
		prev[i] = -1
	}

	hamming := func(a, b string) int {
		if len(a) != len(b) {
			return -1
		}
		diff := 0
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff
	}

	bestLen := 0
	bestIdx := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if groups[j] != groups[i] && hamming(words[j], words[i]) == 1 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					prev[i] = j
				}
			}
		}
		if dp[i] > bestLen {
			bestLen = dp[i]
			bestIdx = i
		}
	}

	result := make([]string, bestLen)
	for i := bestLen - 1; i >= 0; i-- {
		result[i] = words[bestIdx]
		bestIdx = prev[bestIdx]
	}

	return result
}

func main() {
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceIi([]string{"bab", "dab", "cab"}, []int{1, 2, 2}))
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceIi([]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4}))
}
```
