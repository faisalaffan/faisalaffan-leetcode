# 2085 — Count Common Words With One Occurrence

## Deskripsi

**Soal:** [2085. Count Common Words With One Occurrence](https://leetcode.com/problems/count-common-words-with-one-occurrence/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + m), Space: O(n + m)  
**Kompleksitas Ruang:** O(n + m)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2085: Count Common Words With One Occurrence
// https://leetcode.com/problems/count-common-words-with-one-occurrence/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountCommonWordsWithOneOccurrence(
		[]string{"leetcode", "is", "amazing", "as", "is"},
		[]string{"amazing", "leetcode", "is"},
	)) // 2
	fmt.Println(CountCommonWordsWithOneOccurrence(
		[]string{"a", "ab"},
		[]string{"a", "a", "a", "ab"},
	)) // 1
}

// Time: O(n + m), Space: O(n + m)
func CountCommonWordsWithOneOccurrence(words1 []string, words2 []string) int {
  // Membuat map untuk pencarian O(1): key → value
	freq1 := make(map[string]int)
  // Membuat map untuk pencarian O(1): key → value
	freq2 := make(map[string]int)

	for _, w := range words1 {
		freq1[w]++
	}
	for _, w := range words2 {
		freq2[w]++
	}

	count := 0
	for w, c := range freq1 {
		if c == 1 && freq2[w] == 1 {
			count++
		}
	}
	return count
}
```
