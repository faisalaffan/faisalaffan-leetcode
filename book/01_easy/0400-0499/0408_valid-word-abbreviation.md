# 0408 — Valid Word Abbreviation

## Deskripsi

**Soal:** [0408. Valid Word Abbreviation](https://leetcode.com/problems/valid-word-abbreviation/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ValidWordAbbreviation(word, abbr string) bool`

## Solusi Go

```go
package main

// LeetCode #408: Valid Word Abbreviation
// https://leetcode.com/problems/valid-word-abbreviation/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n), Space: O(1)
func ValidWordAbbreviation(word, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if abbr[j] >= 'a' && abbr[j] <= 'z' {
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
			continue
		}
		if abbr[j] == '0' {
			return false
		}
		num := 0
		for j < len(abbr) && abbr[j] >= '0' && abbr[j] <= '9' {
			num = num*10 + int(abbr[j]-'0')
			j++
		}
		i += num
	}
	return i == len(word) && j == len(abbr)
}

func main() {
	fmt.Println(ValidWordAbbreviation("internationalization", "i12iz4n"))
	fmt.Println(ValidWordAbbreviation("apple", "a2e"))
	fmt.Println(ValidWordAbbreviation("hi", "1"))
}
```
