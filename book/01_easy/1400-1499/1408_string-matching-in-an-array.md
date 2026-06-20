# 1408 — String Matching In An Array

## Deskripsi

**Soal:** [1408. String Matching In An Array](https://leetcode.com/problems/string-matching-in-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2 * L) where L is average word length, Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func stringMatching(words []string) []string`

## Solusi Go

```go
package main

// LeetCode #1408: String Matching in an Array
// https://leetcode.com/problems/string-matching-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func stringMatching(words []string) []string

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(StringMatchingInAnArray([]string{"mass", "as", "hero", "superhero"})) // [as hero]
	fmt.Println(StringMatchingInAnArray([]string{"leetcode", "et", "code"}))          // [et code]
}

// Time: O(n^2 * L) where L is average word length, Space: O(n)
func StringMatchingInAnArray(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
  // Membuat slice untuk menyimpan hasil
	res := make([]string, 0)
	for i, w := range words {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], w) {
				res = append(res, w)
				break
			}
		}
	}
	return res
}
```
