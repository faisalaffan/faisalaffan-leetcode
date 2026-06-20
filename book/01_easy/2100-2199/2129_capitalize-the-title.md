# 2129 — Capitalize The Title

## Deskripsi

**Soal:** [2129. Capitalize The Title](https://leetcode.com/problems/capitalize-the-title/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2129: Capitalize the Title
// https://leetcode.com/problems/capitalize-the-title/
// Difficulty: Easy

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(CapitalizeTheTitle("capiTalIze tHe titLe")) // "Capitalize The Title"
	fmt.Println(CapitalizeTheTitle("First leTTER of EACH Word")) // "First Letter of Each Word"
	fmt.Println(CapitalizeTheTitle("i lOve leetcode"))           // "i Love Leetcode"
}

// Time: O(n), Space: O(n)
func CapitalizeTheTitle(title string) string {
	words := strings.Fields(title)
	for i, w := range words {
		lower := strings.ToLower(w)
		if len(lower) > 2 {
			runes := []rune(lower)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		} else {
			words[i] = lower
		}
	}
	return strings.Join(words, " ")
}
```
