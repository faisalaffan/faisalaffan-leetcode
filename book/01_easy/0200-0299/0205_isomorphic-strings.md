# 0205 — Isomorphic Strings

## Deskripsi

**Soal:** [0205. Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (fixed ASCII chars)

**Algoritma:** —

**Fungsi Solusi:** `func IsIsomorphic(s string, t string) bool`

## Solusi Go

```go
package main

// LeetCode #205: Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (fixed ASCII chars)
func IsIsomorphic(s string, t string) bool {
  // Membuat slice untuk menyimpan hasil
	m1 := make([]int, 256)
  // Membuat slice untuk menyimpan hasil
	m2 := make([]int, 256)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

func main() {
	fmt.Println(IsIsomorphic("egg", "add"))
	fmt.Println(IsIsomorphic("foo", "bar"))
	fmt.Println(IsIsomorphic("paper", "title"))
}
```
