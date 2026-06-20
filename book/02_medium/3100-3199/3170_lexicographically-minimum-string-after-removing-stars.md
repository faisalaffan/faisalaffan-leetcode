# 3170 — Lexicographically Minimum String After Removing Stars

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func clearStars(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(n * 26)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3170: Lexicographically Minimum String After Removing Stars
// https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/
// Difficulty: Medium
// Time: O(n * 26) | Space: O(n)

import "fmt"

func clearStars(s string) string {
	n := len(s)
	bytes := []byte(s)
  // Matriks 2D
	queues := make([][]int, 26)
  // Range loop
	for i := range queues {
		queues[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		if s[i] == '*' {
			for j := 0; j < 26; j++ {
				if len(queues[j]) > 0 {
					idx := queues[j][len(queues[j])-1]
					queues[j] = queues[j][:len(queues[j])-1]
					bytes[idx] = '*'
					break
				}
			}
			bytes[i] = '*'
		} else {
			queues[s[i]-'a'] = append(queues[s[i]-'a'], i)
		}
	}

	ans := make([]byte, 0, n)
	for _, ch := range bytes {
		if ch != '*' {
			ans = append(ans, ch)
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(clearStars("aaba*"))       // Expected: "aab"
	fmt.Println(clearStars("abc"))          // Expected: "abc"
	fmt.Println(clearStars("a*b*c*"))       // Expected: ""
}
```
