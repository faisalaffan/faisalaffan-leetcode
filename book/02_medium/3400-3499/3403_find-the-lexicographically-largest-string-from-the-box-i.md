# 3403 — Find The Lexicographically Largest String From The Box I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func answerString(word string, numFriends int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2) Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3403: Find the Lexicographically Largest String From the Box I
// https://leetcode.com/problems/find-the-lexicographically-largest-string-from-the-box-i/
// Difficulty: Medium
// Time: O(n^2) Space: O(n)

import "fmt"

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	n := len(word)
	maxLen := n - numFriends + 1
	ans := word[:maxLen]
	for i := 0; i < n; i++ {
		end := i + maxLen
		if end > n {
			end = n
		}
		sub := word[i:end]
		if sub > ans {
			ans = sub
		}
	}
	return ans
}

func main() {
	fmt.Println(answerString("dbca", 2)) // "dbc"
	fmt.Println(answerString("gggg", 2)) // "ggg"
	fmt.Println(answerString("abc", 3))  // "c"
}
```
