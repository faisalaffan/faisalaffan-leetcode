# 0131 — Palindrome Partitioning

## Deskripsi

**Soal:** [0131. Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 2^n)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** —

**Fungsi Solusi:** `func partition(s string) [][]string`

## Solusi Go

```go
package main

// LeetCode #131: Palindrome Partitioning
// https://leetcode.com/problems/palindrome-partitioning/
// Difficulty: Medium

import "fmt"

func partition(s string) [][]string {
	result := [][]string{}
	n := len(s)

	// Precompute palindrome table
  // Membuat slice 2D untuk DP/tabel
	pal := make([][]bool, n)
  // Iterasi seluruh elemen
	for i := range pal {
		pal[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || pal[i+1][j-1]) {
				pal[i][j] = true
			}
		}
	}

	var backtrack func(start int, path []string)
	backtrack = func(start int, path []string) {
		if start == n {
  // Membuat slice untuk menyimpan hasil
			part := make([]string, len(path))
			copy(part, path)
			result = append(result, part)
			return
		}
		for end := start; end < n; end++ {
			if pal[start][end] {
				path = append(path, s[start:end+1])
				backtrack(end+1, path)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0, []string{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(partition("aab")) // [["a","a","b"],["aa","b"]]

	// Test case 2
	fmt.Println(partition("a")) // [["a"]]

	// Test case 3
	fmt.Println(partition("ab")) // [["a","b"]]
}

// Time: O(n * 2^n) | Space: O(n^2)
```
