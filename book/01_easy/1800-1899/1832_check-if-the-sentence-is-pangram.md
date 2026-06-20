# 1832 — Check If The Sentence Is Pangram

## Deskripsi

**Soal:** [1832. Check If The Sentence Is Pangram](https://leetcode.com/problems/check-if-the-sentence-is-pangram/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CheckIfPangram(sentence string) bool`

## Solusi Go

```go
package main

// LeetCode #1832: Check if the Sentence Is Pangram
// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckIfPangram(sentence string) bool {
	seen := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(sentence); i++ {
		seen |= 1 << (sentence[i] - 'a')
	}
	return seen == (1<<26)-1
}

func main() {
	fmt.Println(CheckIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(CheckIfPangram("leetcode"))
}
```
