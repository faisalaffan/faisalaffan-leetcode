# 3403 — Find The Lexicographically Largest String From The Box I

## Deskripsi

**Soal:** [3403. Find The Lexicographically Largest String From The Box I](https://leetcode.com/problems/find-the-lexicographically-largest-string-from-the-box-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func answerString(word string, numFriends int) string`

## Solusi Go

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
