# 1415 — The K Th Lexicographical String Of All Happy Strings Of Length N

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func getHappyString(n int, k int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) where n = length of string  |  **Ruang:** O(n) for recursion


## 💻 Solusi Go

```go
package main

// LeetCode #1415: The k-th Lexicographical String of All Happy Strings of Length n
// https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(getHappyString(1, 3)) // "c"

	// Test case 2
	fmt.Println(getHappyString(1, 4)) // ""

	// Test case 3
	fmt.Println(getHappyString(3, 9)) // "cab"

	// Test case 4
	fmt.Println(getHappyString(2, 7)) // ""
}

// Time: O(n) where n = length of string
// Space: O(n) for recursion
func getHappyString(n int, k int) string {
	// Total happy strings = 3 * 2^(n-1)
	total := 1
	for i := 1; i < n; i++ {
		total *= 2
	}
	total *= 3

	if k > total {
		return ""
	}

	result := make([]byte, n)
	choices := []byte{'a', 'b', 'c'}

	// Each position's choices depend on previous
	prev := byte(0)
	for i := 0; i < n; i++ {
		for _, c := range choices {
			if c == prev {
				continue
			}
			// Count remaining strings if we pick c here
			remaining := 1
			for j := i + 1; j < n; j++ {
				remaining *= 2
			}
			if k > remaining {
				k -= remaining
			} else {
				result[i] = c
				prev = c
				break
			}
		}
	}

	return string(result)
}
```
