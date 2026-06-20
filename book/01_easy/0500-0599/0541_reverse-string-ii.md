# 0541 — Reverse String Ii

## Deskripsi

**Soal:** [0541. Reverse String Ii](https://leetcode.com/problems/reverse-string-ii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func ReverseStringIi(s string, k int) string`

## Solusi Go

```go
package main

// LeetCode #541: Reverse String II
// https://leetcode.com/problems/reverse-string-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseStringIi(s string, k int) string {
	b := []byte(s)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(b); i += 2 * k {
		lo, hi := i, i+k-1
		if hi >= len(b) {
			hi = len(b) - 1
		}
		for lo < hi {
			b[lo], b[hi] = b[hi], b[lo]
			lo++
			hi--
		}
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseStringIi("abcdefg", 2))
	fmt.Println(ReverseStringIi("abcd", 2))
}
```
