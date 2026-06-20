# 2370 — Longest Ideal Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestIdealString(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * 26)  
**Kompleksitas Ruang:** O(26)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2370: Longest Ideal Subsequence
// https://leetcode.com/problems/longest-ideal-subsequence/
// Difficulty: Medium
// Time: O(n * 26) | Space: O(26)
// DP over alphabet: dp[c] = longest ideal subsequence ending with char c.

import "fmt"

func main() {
	fmt.Println(longestIdealString("acfgbd", 2)) // 4
	fmt.Println(longestIdealString("abcd", 3))   // 4
}

func longestIdealString(s string, k int) int {
  // Alokasi slice integer
	dp := make([]int, 26)
	var ans int
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		best := 0
		for p := 0; p < 26; p++ {
			if abs(c-p) <= k && dp[p] > best {
				best = dp[p]
			}
		}
		dp[c] = best + 1
		if dp[c] > ans {
			ans = dp[c]
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
