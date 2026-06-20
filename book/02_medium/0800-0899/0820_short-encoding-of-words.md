# 0820 — Short Encoding Of Words

## Deskripsi

**Soal:** [0820. Short Encoding Of Words](https://leetcode.com/problems/short-encoding-of-words/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * L^2)  
**Kompleksitas Ruang:** O(n * L) where L is average word length

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #820: Short Encoding of Words
// https://leetcode.com/problems/short-encoding-of-words/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ShortEncodingOfWords([]string{"time", "me", "bell"}))
	fmt.Println(ShortEncodingOfWords([]string{"t"}))
	fmt.Println(ShortEncodingOfWords([]string{"me", "time"}))
}

// Time: O(n * L^2) | Space: O(n * L) where L is average word length
func ShortEncodingOfWords(words []string) int {
  // Membuat map untuk pencarian O(1): key → value
	set := make(map[string]bool)
	for _, word := range words {
		set[word] = true
	}

	for _, word := range words {
		for i := 1; i < len(word); i++ {
			delete(set, word[i:])
		}
	}

	ans := 0
	for word := range set {
		ans += len(word) + 1
	}
	return ans
}
```
