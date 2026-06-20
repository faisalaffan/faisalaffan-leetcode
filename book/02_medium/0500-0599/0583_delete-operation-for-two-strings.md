# 0583 — Delete Operation For Two Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinDistance(word1 string, word2 string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #583: Delete Operation for Two Strings
// https://leetcode.com/problems/delete-operation-for-two-strings/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinDistance("sea", "eat"))
	fmt.Println(MinDistance("leetcode", "etco"))
}

func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
  // Alokasi slice integer
	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		prev := 0
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = prev + 1
			} else {
				if dp[j] > dp[j-1] {
					dp[j] = dp[j]
				} else {
					dp[j] = dp[j-1]
				}
			}
			prev = temp
		}
	}

	lcs := dp[n]
	return (m - lcs) + (n - lcs)
}
```
