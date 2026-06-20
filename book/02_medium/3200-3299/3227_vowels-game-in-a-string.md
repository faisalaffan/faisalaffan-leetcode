# 3227 — Vowels Game In A String

## Deskripsi

**Soal:** [3227. Vowels Game In A String](https://leetcode.com/problems/vowels-game-in-a-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func doesAliceWin(s string) bool`

## Solusi Go

```go
package main

// LeetCode #3227: Vowels Game in a String
// https://leetcode.com/problems/vowels-game-in-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func doesAliceWin(s string) bool {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if vowels[s[i]] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(doesAliceWin("leetcoder")) // Expected: true
	fmt.Println(doesAliceWin("bbcd"))       // Expected: false
}
```
