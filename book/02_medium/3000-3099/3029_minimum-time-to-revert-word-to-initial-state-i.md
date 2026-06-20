# 3029 — Minimum Time To Revert Word To Initial State I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumTimeToInitialState(word string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3029: Minimum Time to Revert Word to Initial State I
// https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(minimumTimeToInitialState("abacaba", 3))
	fmt.Println(minimumTimeToInitialState("abacaba", 2))
	fmt.Println(minimumTimeToInitialState("abcbabcd", 2))
}

func minimumTimeToInitialState(word string, k int) int {
	n := len(word)
  // Alokasi slice integer
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && word[i] != word[j] {
			j = pi[j-1]
		}
		if word[i] == word[j] {
			j++
		}
		pi[i] = j
	}
	j := n
	for j > 0 && j > n-k {
		j = pi[j-1]
	}
	for t := 1; ; t++ {
		if t*k >= n {
			return t
		}
		if n-t*k <= j && (n-t*k)%k == 0 {
			return t
		}
	}
}
```
