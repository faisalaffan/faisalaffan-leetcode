# 0758 — Bold Words In String

## Deskripsi

**Soal:** [0758. Bold Words In String](https://leetcode.com/problems/bold-words-in-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * L)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #758: Bold Words in String
// https://leetcode.com/problems/bold-words-in-string/
// Difficulty: Medium [Paid]
// Time: O(n * L)
// Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(boldWords([]string{"ab", "bc"}, "aabcd"))
	fmt.Println(boldWords([]string{"abc", "123"}, "abcxyz123"))
}

func boldWords(words []string, s string) string {
	n := len(s)
  // Membuat slice untuk menyimpan hasil
	bold := make([]bool, n)

	for _, word := range words {
		start := 0
		for {
			idx := strings.Index(s[start:], word)
			if idx == -1 {
				break
			}
			pos := start + idx
			for i := pos; i < pos+len(word); i++ {
				bold[i] = true
			}
			start = pos + 1
		}
	}

	var result strings.Builder
	i := 0
	for i < n {
		if bold[i] {
			result.WriteString("<b>")
			for i < n && bold[i] {
				result.WriteByte(s[i])
				i++
			}
			result.WriteString("</b>")
		} else {
			result.WriteByte(s[i])
			i++
		}
	}

	return result.String()
}
```
