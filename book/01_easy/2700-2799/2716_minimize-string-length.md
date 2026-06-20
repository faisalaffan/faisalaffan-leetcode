# 2716 — Minimize String Length

## Deskripsi

**Soal:** [2716. Minimize String Length](https://leetcode.com/problems/minimize-string-length/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2716: Minimize String Length
// https://leetcode.com/problems/minimize-string-length/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinimizeStringLength("aaabc"))
	fmt.Println(MinimizeStringLength("cbbd"))
}

func MinimizeStringLength(s string) int {
	seen := map[rune]bool{}
	for _, c := range s {
		seen[c] = true
	}
	return len(seen)
}
```
