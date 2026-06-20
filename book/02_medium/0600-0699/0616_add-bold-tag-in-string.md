# 0616 — Add Bold Tag In String

## Deskripsi

**Soal:** [0616. Add Bold Tag In String](https://leetcode.com/problems/add-bold-tag-in-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * L) where n = len(s), L = total length of all words  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #616: Add Bold Tag in String
// https://leetcode.com/problems/add-bold-tag-in-string/
// Difficulty: Medium [Paid]
// Time: O(n * L) where n = len(s), L = total length of all words
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(AddBoldTag("abcxyz123", []string{"abc", "123"}))
	fmt.Println(AddBoldTag("aaabbcc", []string{"aaa", "aab", "bc"}))
}

func AddBoldTag(s string, words []string) string {
	n := len(s)
  // Membuat slice untuk menyimpan hasil
	bold := make([]bool, n)

	for _, word := range words {
		for i := 0; i <= n-len(word); i++ {
			if s[i:i+len(word)] == word {
				for j := i; j < i+len(word); j++ {
					bold[j] = true
				}
			}
		}
	}

	result := ""
	i := 0
	for i < n {
		if bold[i] {
			result += "<b>"
			for i < n && bold[i] {
				result += string(s[i])
				i++
			}
			result += "</b>"
		} else {
			result += string(s[i])
			i++
		}
	}

	return result
}
```
