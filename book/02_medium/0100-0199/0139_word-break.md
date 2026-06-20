# 0139 — Word Break

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func wordBreak(s string, wordDict []string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #139: Word Break
// https://leetcode.com/problems/word-break/
// Difficulty: Medium

import "fmt"

func wordBreak(s string, wordDict []string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {
	// Test case 1
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"})) // true

	// Test case 2
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"})) // true

	// Test case 3
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}

// Time: O(n^2) | Space: O(n)
```
