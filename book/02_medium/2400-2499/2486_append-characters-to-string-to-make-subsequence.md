# 2486 — Append Characters To String To Make Subsequence

## Deskripsi

**Soal:** [2486. Append Characters To String To Make Subsequence](https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Two Pointer (penunjuk kiri & kanan), Two Pointer (penunjuk kiri & kanan)

## Solusi Go

```go
package main

// LeetCode #2486: Append Characters to String to Make Subsequence
// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/
// Difficulty: Medium
// Time: O(n + m) | Space: O(1)
// Two pointers: match as much of t in s as possible.

import "fmt"

func main() {
	fmt.Println(appendCharacters("coaching", "coding")) // 4
	fmt.Println(appendCharacters("abcde", "a"))         // 0
	fmt.Println(appendCharacters("z", "abcde"))         // 5
}

func appendCharacters(s string, t string) int {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return len(t) - j
}
```
