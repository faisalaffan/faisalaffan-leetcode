# 3374 — First Letter Capitalization Ii

## Deskripsi

**Soal:** [3374. First Letter Capitalization Ii](https://leetcode.com/problems/first-letter-capitalization-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3374: First Letter Capitalization II
// https://leetcode.com/problems/first-letter-capitalization-ii/
// Difficulty: Hard
//
// Capitalize first letter of each word after punctuation separators.

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(FirstLetterCapitalizationIi("hello world"))
	fmt.Println(FirstLetterCapitalizationIi("Leetcode is fun"))
}

func FirstLetterCapitalizationIi(title string) string {
	r := []rune(title)
	n := len(r)
	capitalize := true
	for i := 0; i < n; i++ {
		if unicode.IsLetter(r[i]) {
			if capitalize {
				r[i] = unicode.ToUpper(r[i])
				capitalize = false
			} else {
				r[i] = unicode.ToLower(r[i])
			}
		} else {
			capitalize = true
		}
	}
	return string(r)
}
```
