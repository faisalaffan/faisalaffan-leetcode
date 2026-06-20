# 1763 — Longest Nice Substring

## Deskripsi

**Soal:** [1763. Longest Nice Substring](https://leetcode.com/problems/longest-nice-substring/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func LongestNiceSubstring(s string) string`

## Solusi Go

```go
package main

// LeetCode #1763: Longest Nice Substring
// https://leetcode.com/problems/longest-nice-substring/
// Difficulty: Easy

import "fmt"
import "unicode"

// Time: O(n^2), Space: O(n)
func LongestNiceSubstring(s string) string {
	result := ""
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		lower := 0
		upper := 0
		for j := i; j < len(s); j++ {
			ch := rune(s[j])
			if unicode.IsUpper(ch) {
				upper |= 1 << (unicode.ToLower(ch) - 'a')
			} else {
				lower |= 1 << (ch - 'a')
			}
			if lower == upper && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(LongestNiceSubstring("YazaAay"))
	fmt.Println(LongestNiceSubstring("Bb"))
	fmt.Println(LongestNiceSubstring("c"))
}
```
