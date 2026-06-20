# 2042 — Check If Numbers Are Ascending In A Sentence

## Deskripsi

**Soal:** [2042. Check If Numbers Are Ascending In A Sentence](https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2042: Check if Numbers Are Ascending in a Sentence
// https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/
// Difficulty: Easy

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(CheckIfNumbersAreAscendingInASentence("1 box has 3 blue 4 red 6 green and 12 yellow marbles")) // true
	fmt.Println(CheckIfNumbersAreAscendingInASentence("hello world 5 x 5"))                                     // false
	fmt.Println(CheckIfNumbersAreAscendingInASentence("sunset is at 7 11 pm overnight"))                        // false
}

// Time: O(n), Space: O(1)
func CheckIfNumbersAreAscendingInASentence(s string) bool {
	prev := 0
	num := 0
	hasNum := false

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			num = num*10 + int(s[i]-'0')
			hasNum = true
		} else {
			if hasNum {
				if num <= prev {
					return false
				}
				prev = num
				num = 0
				hasNum = false
			}
		}
	}
	if hasNum && num <= prev {
		return false
	}
	return true
}
```
