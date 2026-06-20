# 1163 — Last Substring In Lexicographical Order

## Deskripsi

**Soal:** [1163. Last Substring In Lexicographical Order](https://leetcode.com/problems/last-substring-in-lexicographical-order/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func lastSubstring(s string) string`

## Solusi Go

```go
package main

// LeetCode #1163: Last Substring in Lexicographical Order
// https://leetcode.com/problems/last-substring-in-lexicographical-order/
// Difficulty: Hard

import "fmt"

// lastSubstring returns the lexicographically largest substring.
// Uses two-pointer technique: i is best candidate start, j is current checking start.
func lastSubstring(s string) string {
	n := len(s)
	i, j, k := 0, 1, 0

	for j+k < n {
		if s[i+k] == s[j+k] {
			k++
			continue
		}
		if s[i+k] < s[j+k] {
			// s[i..i+k] is smaller, so the best candidate starts after i+k
			i = i + k + 1
			if i >= j {
				j = i + 1
			}
		} else {
			// s[j..j+k] is smaller, so move j forward
			j = j + k + 1
		}
		k = 0
	}
	return s[i:]
}

func main() {
	// Test case 1
	fmt.Println(lastSubstring("abab")) // "bab"

	// Test case 2
	fmt.Println(lastSubstring("leetcode")) // "tcode"

	// Test case 3: single char
	fmt.Println(lastSubstring("a")) // "a"

	// Test case 4
	fmt.Println(lastSubstring("cacacb")) // "cb"

	// Test case 5
	fmt.Println(lastSubstring("babcab")) // "cab"
}
```
