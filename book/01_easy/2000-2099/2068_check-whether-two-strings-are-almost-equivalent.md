# 2068 — Check Whether Two Strings Are Almost Equivalent

## Deskripsi

**Soal:** [2068. Check Whether Two Strings Are Almost Equivalent](https://leetcode.com/problems/check-whether-two-strings-are-almost-equivalent/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2068: Check Whether Two Strings are Almost Equivalent
// https://leetcode.com/problems/check-whether-two-strings-are-almost-equivalent/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("aaaa", "bccb"))   // false
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("abcdeef", "abaaacc")) // true
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("cccddabba", "babababab")) // true
}

// Time: O(n), Space: O(1)
func CheckWhetherTwoStringsAreAlmostEquivalent(word1 string, word2 string) bool {
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(word1); i++ {
		freq[word1[i]-'a']++
	}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(word2); i++ {
		freq[word2[i]-'a']--
	}
	for _, v := range freq {
		if v < 0 {
			v = -v
		}
		if v > 3 {
			return false
		}
	}
	return true
}
```
