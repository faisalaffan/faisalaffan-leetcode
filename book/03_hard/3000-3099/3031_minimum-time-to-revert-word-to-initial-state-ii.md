# 3031 — Minimum Time To Revert Word To Initial State Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumTimeToInitialState(word string, k int) int
```

> **💡 Hint:** Z-algorithm (linear-time pattern matching)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3031: Minimum Time to Revert Word to Initial State II
// https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/
// Difficulty: Hard
//
// We have a word. In each operation, we remove the first k characters
// and append k arbitrary characters to the end. We want the word to
// return to its initial state after some number of operations.
// Find the minimum number of operations needed.
//
// Approach: Z-algorithm (linear-time pattern matching)
//   After t operations, the first t*k characters have been removed.
//   The word matches its initial state if the suffix starting at t*k
//   matches the prefix of length n - t*k. i.e., the suffix starting
//   at position t*k is a prefix of the original word (of length >= n - t*k).
//   Use the Z-array to find the longest common prefix between word and
//   each suffix starting at position pos = t*k.

import "fmt"

func minimumTimeToInitialState(word string, k int) int {
	n := len(word)

	// Build Z-array
	// z[i] = length of the longest common prefix between word and word[i:]
  // Alokasi slice integer
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min2(r-i+1, z[i-l])
		}
		for i+z[i] < n && word[z[i]] == word[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}

	// Check each time t
	// At time t, we have removed t*k prefix characters.
	// The remaining string is word[t*k:]. It matches the initial
	// state if the suffix word[t*k:] is a prefix of word (i.e., the entire
	// remaining string matches the initial prefix).
	// This is equivalent to z[t*k] >= n - t*k
	for t := 1; t*k < n; t++ {
		pos := t * k
		if z[pos] >= n-pos {
			return t
		}
	}

	// If no match found, we need ceil(n/k) operations
	return (n + k - 1) / k
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example: "abaca", k=3 -> 3
	// Step 1: remove "aba", append "???" -> "aca???"
	// Step 2: remove "aca", append "???" -> "??????"
	// Step 3: after 3 operations, word = "abaca" again (need 3 ops)
	fmt.Println("Test 1:", minimumTimeToInitialState("abaca", 3))

	// Example: "abacaba", k=3 -> 2
	fmt.Println("Test 2:", minimumTimeToInitialState("abacaba", 3))

	// Example: "abacaba", k=2 -> 4
	fmt.Println("Test 3:", minimumTimeToInitialState("abacaba", 2))

	// Single character
	fmt.Println("Test 4:", minimumTimeToInitialState("a", 1))

	// All same characters
	fmt.Println("Test 5:", minimumTimeToInitialState("aaaa", 2))

	// k = n (remove all characters at once)
	fmt.Println("Test 6:", minimumTimeToInitialState("hello", 5))

	// k = 1
	fmt.Println("Test 7:", minimumTimeToInitialState("abcabc", 3))

	// Longer string
	fmt.Println("Test 8:", minimumTimeToInitialState("ababababab", 2))
}
```
