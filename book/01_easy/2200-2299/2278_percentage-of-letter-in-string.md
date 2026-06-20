# 2278 — Percentage Of Letter In String

## Deskripsi

**Soal:** [2278. Percentage Of Letter In String](https://leetcode.com/problems/percentage-of-letter-in-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2278: Percentage of Letter in String
// https://leetcode.com/problems/percentage-of-letter-in-string/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(PercentageOfLetterInString("foobar", 'o')) // 33
	fmt.Println(PercentageOfLetterInString("jjjj", 'k'))   // 0
	fmt.Println(PercentageOfLetterInString("sgawtb", 's')) // 16
}

func PercentageOfLetterInString(s string, letter byte) int {
	count := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			count++
		}
	}
	return count * 100 / len(s)
}
```
