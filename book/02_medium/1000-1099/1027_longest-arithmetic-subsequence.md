# 1027 — Longest Arithmetic Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestArithSeqLength(nums []int) int
```

> **💡 Hint:** DP with hash map per index tracking difference -> length

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1027: Longest Arithmetic Subsequence
// https://leetcode.com/problems/longest-arithmetic-subsequence/
// Difficulty: Medium
//
// Approach: DP with hash map per index tracking difference -> length
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(longestArithSeqLength([]int{3, 6, 9, 12}))    // 4
	fmt.Println(longestArithSeqLength([]int{9, 4, 7, 2, 10})) // 3
	fmt.Println(longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8})) // 4
}

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

  // Alokasi slice integer
	dp := make([]map[int]int, n)
	result := 2

	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			length := 2
			if prev, ok := dp[j][diff]; ok {
				length = prev + 1
			}
			dp[i][diff] = length
			if length > result {
				result = length
			}
		}
	}

	return result
}
```
