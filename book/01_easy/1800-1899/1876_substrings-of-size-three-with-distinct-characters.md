# 1876 — Substrings Of Size Three With Distinct Characters

## Deskripsi

**Soal:** [1876. Substrings Of Size Three With Distinct Characters](https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CountGoodSubstrings(s string) int`

## Solusi Go

```go
package main

// LeetCode #1876: Substrings of Size Three with Distinct Characters
// https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountGoodSubstrings(s string) int {
	count := 0
	for i := 0; i+2 < len(s); i++ {
		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountGoodSubstrings("xyzzaz"))
	fmt.Println(CountGoodSubstrings("aababcabc"))
}
```
