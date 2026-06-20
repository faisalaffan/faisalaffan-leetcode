# 0345 — Reverse Vowels Of A String

## Deskripsi

**Soal:** [0345. Reverse Vowels Of A String](https://leetcode.com/problems/reverse-vowels-of-a-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func ReverseVowelsOfAString(s string) string`

## Solusi Go

```go
package main

// LeetCode #345: Reverse Vowels of a String
// https://leetcode.com/problems/reverse-vowels-of-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseVowelsOfAString(s string) string {
	b := []byte(s)
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}
	for i, j := 0, len(b)-1; i < j; {
		if !isVowel(b[i]) {
			i++
			continue
		}
		if !isVowel(b[j]) {
			j--
			continue
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseVowelsOfAString("hello"))
	fmt.Println(ReverseVowelsOfAString("leetcode"))
	fmt.Println(ReverseVowelsOfAString("aA"))
}
```
