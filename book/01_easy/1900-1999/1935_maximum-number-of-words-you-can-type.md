# 1935 — Maximum Number Of Words You Can Type

## Deskripsi

**Soal:** [1935. Maximum Number Of Words You Can Type](https://leetcode.com/problems/maximum-number-of-words-you-can-type/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + m), Space: O(k) where k = len(brokenLetters)  
**Kompleksitas Ruang:** O(k) where k = len(brokenLetters)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1935: Maximum Number of Words You Can Type
// https://leetcode.com/problems/maximum-number-of-words-you-can-type/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumNumberOfWordsYouCanType("hello world", "ad"))                   // 1
	fmt.Println(MaximumNumberOfWordsYouCanType("leet code", "e"))                      // 0
	fmt.Println(MaximumNumberOfWordsYouCanType("leet code", "lt"))                     // 1
}

// Time: O(n + m), Space: O(k) where k = len(brokenLetters)
func MaximumNumberOfWordsYouCanType(text string, brokenLetters string) int {
  // Membuat map untuk pencarian O(1): key → value
	broken := make(map[byte]bool)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(brokenLetters); i++ {
		broken[brokenLetters[i]] = true
	}

	words := strings.Fields(text)
	count := 0
	for _, word := range words {
		canType := true
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(word); i++ {
			if broken[word[i]] {
				canType = false
				break
			}
		}
		if canType {
			count++
		}
	}
	return count
}
```
