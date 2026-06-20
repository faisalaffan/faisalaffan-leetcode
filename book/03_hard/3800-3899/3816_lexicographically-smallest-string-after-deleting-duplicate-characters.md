# 3816 — Lexicographically Smallest String After Deleting Duplicate Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func lexSmallestAfterDeletion(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack, Monotonic Stack

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3816: Lexicographically Smallest String After Deleting
// Duplicate Characters
// https://leetcode.com/problems/lexicographically-smallest-string-after-deleting-duplicate-characters/
// Difficulty: Hard
//
// Delete duplicate characters from s, keeping each at most once,
// to obtain lexicographically smallest possible string.
//
// Approach: Monotonic stack. Track last occurrence of each char.
// Maintain stack with increasing characters. Pop if a larger char
// appears later.

import "fmt"

func main() {
	// Example 1
	fmt.Println(lexSmallestAfterDeletion("bcabc"))
	// Example 2
	fmt.Println(lexSmallestAfterDeletion("cbacdcbc"))
	// Edge: already unique
	fmt.Println(lexSmallestAfterDeletion("abc"))
	// Edge: reversed
	fmt.Println(lexSmallestAfterDeletion("cba"))
}

func lexSmallestAfterDeletion(s string) string {
  // Alokasi slice
	lastPos := make([]int, 26)
  // Range loop
	for i := range lastPos {
		lastPos[i] = -1
	}
	for i, ch := range s {
		lastPos[ch-'a'] = i
	}

	used := make([]bool, 26)
	stack := make([]byte, 0, len(s))

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if used[c] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > c && lastPos[stack[len(stack)-1]] > i {
			used[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
		used[c] = true
	}

	res := make([]byte, len(stack))
	for i, v := range stack {
		res[i] = v + 'a'
	}
	return string(res)
}
```
