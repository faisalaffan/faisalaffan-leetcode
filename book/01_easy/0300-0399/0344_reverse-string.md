# 0344 — Reverse String

## Deskripsi

**Soal:** [0344. Reverse String](https://leetcode.com/problems/reverse-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ReverseString(s []byte) `

## Solusi Go

```go
package main

// LeetCode #344: Reverse String
// https://leetcode.com/problems/reverse-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func ReverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s1 := []byte("hello")
	ReverseString(s1)
	fmt.Println(string(s1))

	s2 := []byte("Hannah")
	ReverseString(s2)
	fmt.Println(string(s2))
}
```
