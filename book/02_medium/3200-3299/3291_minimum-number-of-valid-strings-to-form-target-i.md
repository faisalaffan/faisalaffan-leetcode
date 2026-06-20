# 3291 — Minimum Number Of Valid Strings To Form Target I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func minValidStrings(words []string, target string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Trie, Prefix Sum

**Waktu:** O(n * L) Space: O(total_chars + n) where L = average prefix length  |  **Ruang:** O(total_chars + n) where L = average prefix length

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3291: Minimum Number of Valid Strings to Form Target I
// https://leetcode.com/problems/minimum-number-of-valid-strings-to-form-target-i/
// Difficulty: Medium
// Time: O(n * L) Space: O(total_chars + n) where L = average prefix length

import (
	"fmt"
)

func main() {
	fmt.Println(minValidStrings([]string{"abc", "aaaaa", "bcdef"}, "aabcdabc")) // 3
	fmt.Println(minValidStrings([]string{"ab", "bc", "cd"}, "abc"))            // 2
	fmt.Println(minValidStrings([]string{"a", "b", "c"}, "xyz"))               // -1
}

type trieNode struct {
	children [26]*trieNode
}

func minValidStrings(words []string, target string) int {
	root := &trieNode{}
	for _, w := range words {
		node := root
		for _, ch := range w {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &trieNode{}
			}
			node = node.children[idx]
		}
	}

	n := len(target)
	inf := int(1e9)
  // Alokasi slice
	dp := make([]int, n+1)
  // Range loop
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		if dp[i] == inf {
			continue
		}
		node := root
		for j := i; j < n; j++ {
			idx := target[j] - 'a'
			if node.children[idx] == nil {
				break
			}
			node = node.children[idx]
			if dp[i]+1 < dp[j+1] {
				dp[j+1] = dp[i] + 1
			}
		}
	}

	if dp[n] == inf {
		return -1
	}
	return dp[n]
}
```
