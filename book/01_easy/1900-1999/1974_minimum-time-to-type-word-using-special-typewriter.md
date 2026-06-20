# 1974 — Minimum Time To Type Word Using Special Typewriter

## Deskripsi

**Soal:** [1974. Minimum Time To Type Word Using Special Typewriter](https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1974: Minimum Time to Type Word Using Special Typewriter
// https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("abc"))  // 5
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("bza"))  // 7
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("zjpc")) // 34
}

// Time: O(n), Space: O(1)
func MinimumTimeToTypeWordUsingSpecialTypewriter(word string) int {
	seconds := 0
	pos := 0 // 'a'
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(word); i++ {
		target := int(word[i] - 'a')
		diff := target - pos
		if diff < 0 {
			diff = -diff
		}
		if diff > 13 {
			diff = 26 - diff
		}
		seconds += diff + 1 // move + type
		pos = target
	}
	return seconds
}
```
