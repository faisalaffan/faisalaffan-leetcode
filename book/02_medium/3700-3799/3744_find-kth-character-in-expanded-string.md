# 3744 — Find Kth Character In Expanded String

## Deskripsi

**Soal:** [3744. Find Kth Character In Expanded String](https://leetcode.com/problems/find-kth-character-in-expanded-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func findKthCharacterInExpandedString(s string, k int) byte`

## Solusi Go

```go
package main

// LeetCode #3744: Find Kth Character in Expanded String
// https://leetcode.com/problems/find-kth-character-in-expanded-string/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func findKthCharacterInExpandedString(s string, k int) byte {
	words := strings.Fields(s)

	for _, word := range words {
		l := len(word)
		m := l * (l + 1) / 2

		if k == m {
			return ' '
		} else if k > m {
			k -= (m + 1)
			continue
		} else {
			cur := 0
			for i, ch := range word {
				cur += (i + 1)
				if k < cur {
					return byte(ch)
				}
			}
			return ' '
		}
	}
	return ' '
}

func main() {
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 0))
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 15))
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 20))
}
```
