# 2390 — Removing Stars From A String

## Deskripsi

**Soal:** [2390. Removing Stars From A String](https://leetcode.com/problems/removing-stars-from-a-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

```go
package main

// LeetCode #2390: Removing Stars From a String
// https://leetcode.com/problems/removing-stars-from-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Stack: push non-star, pop on star.

import "fmt"

func main() {
	fmt.Println(removeStars("leet**cod*e")) // "lecoe"
	fmt.Println(removeStars("erase*****"))  // ""
}

func removeStars(s string) string {
  // Membuat slice untuk menyimpan hasil
	res := make([]byte, 0, len(s))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
```
