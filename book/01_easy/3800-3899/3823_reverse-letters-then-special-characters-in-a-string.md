# 3823 — Reverse Letters Then Special Characters In A String

## Deskripsi

**Soal:** [3823. Reverse Letters Then Special Characters In A String](https://leetcode.com/problems/reverse-letters-then-special-characters-in-a-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3823: Reverse Letters Then Special Characters in a String
// https://leetcode.com/problems/reverse-letters-then-special-characters-in-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReverseLettersThenSpecialCharactersInAString(")ebc#da@f("))
	fmt.Println(ReverseLettersThenSpecialCharactersInAString("z"))
	fmt.Println(ReverseLettersThenSpecialCharactersInAString("!@#$%^&*()"))
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func isSpecial(ch byte) bool {
	return !isLetter(ch)
}

func reverseRange(s []byte, cond func(byte) bool) {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !cond(s[i]) {
			i++
		}
		for i < j && !cond(s[j]) {
			j--
		}
		if i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
}

// Time: O(n)
// Space: O(n)
func ReverseLettersThenSpecialCharactersInAString(s string) string {
	b := []byte(s)
	reverseRange(b, isLetter)
	reverseRange(b, isSpecial)
	return string(b)
}
```
