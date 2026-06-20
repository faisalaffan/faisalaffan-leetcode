# 1312 — Minimum Insertion Steps To Make A String Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minInsertions(s string) int
```

> **💡 Hint:** DP on intervals.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1312: Minimum Insertion Steps to Make a String Palindrome
// https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
// Difficulty: Hard
//
// Approach: DP on intervals.
// dp[i][j] = minimum insertions needed to make s[i..j] a palindrome.
//   If s[i] == s[j]: dp[i][j] = dp[i+1][j-1] (0 for length < 2)
//   Else: dp[i][j] = 1 + min(dp[i+1][j], dp[i][j-1])
// Answer = dp[0][n-1].

import "fmt"

func minInsertions(s string) int {
	n := len(s)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				if i+1 <= j-1 {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = 1 + min(dp[i+1][j], dp[i][j-1])
			}
		}
	}

  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minInsertions("zzazz"))     // 0
	fmt.Println(minInsertions("mbadm"))     // 2
	fmt.Println(minInsertions("leetcode"))  // 5
	fmt.Println(minInsertions("g"))         // 0
	fmt.Println(minInsertions("no"))        // 1
	fmt.Println(minInsertions(""))          // 0
}
```
