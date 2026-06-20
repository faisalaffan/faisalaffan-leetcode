# 0880 — Decoded String At Index

## Deskripsi

**Soal:** [0880. Decoded String At Index](https://leetcode.com/problems/decoded-string-at-index/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #880: Decoded String at Index
// https://leetcode.com/problems/decoded-string-at-index/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(DecodedStringAtIndex("leet2code3", 10))
	fmt.Println(DecodedStringAtIndex("ha22", 5))
	fmt.Println(DecodedStringAtIndex("a2345678999999999999999", 1))
}

// Time: O(n) | Space: O(1)
func DecodedStringAtIndex(s string, k int) string {
	var size int64 = 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			size++
		} else {
			size *= int64(s[i] - '0')
		}
	}

	target := int64(k)
	for i := len(s) - 1; i >= 0; i-- {
		target %= size
		if target == 0 && s[i] >= 'a' && s[i] <= 'z' {
			return string(s[i])
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			size--
		} else {
			size /= int64(s[i] - '0')
		}
	}

	return ""
}
```
