# 0290 — Word Pattern

## Deskripsi

**Soal:** [0290. Word Pattern](https://leetcode.com/problems/word-pattern/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func WordPattern(pattern string, s string) bool`

## Solusi Go

```go
package main

// LeetCode #290: Word Pattern
// https://leetcode.com/problems/word-pattern/
// Difficulty: Easy

import "strings"
import "fmt"

// Time: O(n) | Space: O(n)
func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
  // Membuat map untuk pencarian O(1): key → value
	p2w := make(map[byte]string)
  // Membuat map untuk pencarian O(1): key → value
	w2p := make(map[string]byte)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]
		if mappedW, ok := p2w[p]; ok && mappedW != w {
			return false
		}
		if mappedP, ok := w2p[w]; ok && mappedP != p {
			return false
		}
		p2w[p] = w
		w2p[w] = p
	}
	return true
}

func main() {
	fmt.Println(WordPattern("abba", "dog cat cat dog"))
	fmt.Println(WordPattern("abba", "dog cat cat fish"))
	fmt.Println(WordPattern("aaaa", "dog cat cat dog"))
}
```
