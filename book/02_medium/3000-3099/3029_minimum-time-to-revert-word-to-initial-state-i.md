# 3029 — Minimum Time To Revert Word To Initial State I

## Deskripsi

**Soal:** [3029. Minimum Time To Revert Word To Initial State I](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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
