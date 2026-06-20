# 2904 — Shortest And Lexicographically Smallest Beautiful String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func shortestBeautifulSubstring(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2904: Shortest and Lexicographically Smallest Beautiful String
// https://leetcode.com/problems/shortest-and-lexicographically-smallest-beautiful-string/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(shortestBeautifulSubstring("100011001", 3))
	fmt.Println(shortestBeautifulSubstring("1011", 2))
	fmt.Println(shortestBeautifulSubstring("000", 1))
}

func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)
	ans := ""
	for i := 0; i < n; i++ {
		for j := i + k; j <= n; j++ {
			t := s[i:j]
			cnt := 0
			for _, c := range t {
				if c == '1' {
					cnt++
				}
			}
			if cnt == k && (ans == "" || j-i < len(ans) || (j-i == len(ans) && t < ans)) {
				ans = t
			}
		}
	}
	return ans
}
```
