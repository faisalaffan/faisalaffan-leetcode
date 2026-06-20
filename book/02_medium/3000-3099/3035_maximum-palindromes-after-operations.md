# 3035 — Maximum Palindromes After Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxPalindromesAfterOperations(words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L + A log A)  
**Kompleksitas Ruang:** O(A)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3035: Maximum Palindromes After Operations
// https://leetcode.com/problems/maximum-palindromes-after-operations/
// Difficulty: Medium
// Time: O(n * L + A log A) | Space: O(A)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxPalindromesAfterOperations([]string{"abbb", "ba", "aa"}))
	fmt.Println(maxPalindromesAfterOperations([]string{"abc", "ab"}))
	fmt.Println(maxPalindromesAfterOperations([]string{"cd", "ef", "a"}))
}

func maxPalindromesAfterOperations(words []string) int {
	freq := [26]int{}
  // Alokasi slice integer
	lens := make([]int, len(words))
	for i, w := range words {
		lens[i] = len(w)
		for _, ch := range w {
			freq[ch-'a']++
		}
	}
	pairs := 0
	for _, c := range freq {
		pairs += c / 2
	}
  // Custom sort dengan comparator
	sort.Slice(lens, func(i, j int) bool {
		return lens[i] < lens[j]
	})
	ans := 0
	for _, l := range lens {
		need := l / 2
		if pairs >= need {
			pairs -= need
			ans++
		}
	}
	return ans
}
```
