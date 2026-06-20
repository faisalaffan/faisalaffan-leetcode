# 1662 — Check If Two String Arrays Are Equivalent

## Deskripsi

**Soal:** [1662. Check If Two String Arrays Are Equivalent](https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n) where n is total characters  
**Kompleksitas Ruang:** O(n) where n is total characters

**Algoritma:** —

**Fungsi Solusi:** `func ArrayStringsAreEqual(word1 []string, word2 []string) bool`

## Solusi Go

```go
package main

// LeetCode #1662: Check If Two String Arrays Are Equivalent
// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
// Difficulty: Easy

import "fmt"
import "strings"

// Time: O(n), Space: O(n) where n is total characters
func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func main() {
	fmt.Println(ArrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(ArrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	fmt.Println(ArrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}
```
