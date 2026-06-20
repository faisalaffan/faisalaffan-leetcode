# 2490 — Circular Sentence

## Deskripsi

**Soal:** [2490. Circular Sentence](https://leetcode.com/problems/circular-sentence/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2490: Circular Sentence
// https://leetcode.com/problems/circular-sentence/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CircularSentence("leetcode exercises sound delightful")) // true
	fmt.Println(CircularSentence("eetcode"))                            // true
	fmt.Println(CircularSentence("Leetcode is cool"))                   // false
}

func CircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}
```
