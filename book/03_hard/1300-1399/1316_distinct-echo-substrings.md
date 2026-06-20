# 1316 — Distinct Echo Substrings

## Deskripsi

**Soal:** [1316. Distinct Echo Substrings](https://leetcode.com/problems/distinct-echo-substrings/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func distinctEchoSubstrings(text string) int`

> **Ide Kunci:** Rolling hash + set.

## Solusi Go

```go
package main

// LeetCode #1316: Distinct Echo Substrings
// https://leetcode.com/problems/distinct-echo-substrings/
// Difficulty: Hard
//
// Approach: Rolling hash + set.
// An echo substring is one that can be split into two equal halves.
// For each even length L and each start i, check if text[i:i+L/2] == text[i+L/2:i+L].
// Use a set of string hashes to deduplicate. For correctness we use a simple
// O(n^2) substring equality check on the string itself since n <= 100 for
// this stub. For larger n, rolling hash (Rabin-Karp) should be used.

import "fmt"

func distinctEchoSubstrings(text string) int {
	n := len(text)
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[string]bool)

	for length := 2; length <= n; length += 2 {
		half := length / 2
		for i := 0; i+length <= n; i++ {
			if text[i:i+half] == text[i+half:i+length] {
				seen[text[i:i+length]] = true
			}
		}
	}
	return len(seen)
}

func main() {
	fmt.Println(distinctEchoSubstrings("abcabcabc"))        // 3
	fmt.Println(distinctEchoSubstrings("leetcodeleetcode")) // 2
	fmt.Println(distinctEchoSubstrings("aaa"))              // 1
	fmt.Println(distinctEchoSubstrings(""))                 // 0
	fmt.Println(distinctEchoSubstrings("ab"))               // 0
}
```
