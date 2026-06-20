# 1081 — Smallest Subsequence Of Distinct Characters

## Deskripsi

**Soal:** [1081. Smallest Subsequence Of Distinct Characters](https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(26) = O(1)

**Algoritma:** Greedy (pemilihan optimal lokal), Stack (tumpukan LIFO), Monotonic Stack (tumpukan monoton)

> **Ide Kunci:** Monotonic stack (greedy). Track last occurrence and used set.

## Solusi Go

```go
package main

// LeetCode #1081: Smallest Subsequence of Distinct Characters
// https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
// Difficulty: Medium
//
// Approach: Monotonic stack (greedy). Track last occurrence and used set.
// Time: O(n)
// Space: O(26) = O(1)

import "fmt"

func main() {
	fmt.Println(smallestSubsequence("bcabc"))  // "abc"
	fmt.Println(smallestSubsequence("cbacdcbc")) // "acdb"
}

func smallestSubsequence(s string) string {
  // Membuat slice untuk menyimpan hasil
	lastOccur := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		lastOccur[s[i]-'a'] = i
	}

  // Membuat slice untuk menyimpan hasil
	used := make([]bool, 26)
  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0)

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if used[ch-'a'] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > ch && lastOccur[stack[len(stack)-1]-'a'] > i {
			used[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
		used[ch-'a'] = true
	}

	return string(stack)
}
```
