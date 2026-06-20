# 0791 — Custom Sort String

## Deskripsi

**Soal:** [0791. Custom Sort String](https://leetcode.com/problems/custom-sort-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #791: Custom Sort String
// https://leetcode.com/problems/custom-sort-string/
// Difficulty: Medium
// Time: O(n + m)
// Space: O(1)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println(customSortString("bcafg", "abcd"))
}

func customSortString(order string, s string) string {
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	var result strings.Builder
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(order); i++ {
		c := order[i]
		for freq[c-'a'] > 0 {
			result.WriteByte(c)
			freq[c-'a']--
		}
	}

	for i := 0; i < 26; i++ {
		for freq[i] > 0 {
			result.WriteByte(byte(i + 'a'))
			freq[i]--
		}
	}

	return result.String()
}
```
