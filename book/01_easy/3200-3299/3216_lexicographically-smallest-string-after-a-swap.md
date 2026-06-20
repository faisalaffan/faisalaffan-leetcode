# 3216 — Lexicographically Smallest String After A Swap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func LexicographicallySmallestStringAfterASwap(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3216: Lexicographically Smallest String After a Swap
// https://leetcode.com/problems/lexicographically-smallest-string-after-a-swap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LexicographicallySmallestStringAfterASwap("45320"))
	fmt.Println(LexicographicallySmallestStringAfterASwap("001"))
}

// LexicographicallySmallestStringAfterASwap makes the smallest string by swapping one pair of adjacent same-parity digits where left > right.
// Time: O(n). Space: O(n).
func LexicographicallySmallestStringAfterASwap(s string) string {
	b := []byte(s)
  // Linear scan O(n)
	for i := 0; i < len(b)-1; i++ {
		if b[i] > b[i+1] && (b[i]%2 == b[i+1]%2) {
			b[i], b[i+1] = b[i+1], b[i]
			break
		}
	}
	return string(b)
}
```
