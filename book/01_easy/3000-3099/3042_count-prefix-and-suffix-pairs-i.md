# 3042 — Count Prefix And Suffix Pairs I

## Deskripsi

**Soal:** [3042. Count Prefix And Suffix Pairs I](https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2 * m) where m is max word length  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3042: Count Prefix and Suffix Pairs I
// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: countPrefixSuffixPairs
	fmt.Println(CountPrefixAndSuffixPairsI([]string{"a", "aba", "ababa", "aa"})) // 4
	fmt.Println(CountPrefixAndSuffixPairsI([]string{"pa", "papa", "ma", "mama"})) // 2
}

// Time: O(n^2 * m) where m is max word length | Space: O(1)
// LeetCode submission name: countPrefixSuffixPairs
func CountPrefixAndSuffixPairsI(words []string) int {
	n := len(words)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

func isPrefixAndSuffix(a, b string) bool {
	if len(a) > len(b) {
		return false
	}
	// Check prefix
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	// Check suffix
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(a); i++ {
		if a[i] != b[len(b)-len(a)+i] {
			return false
		}
	}
	return true
}
```
