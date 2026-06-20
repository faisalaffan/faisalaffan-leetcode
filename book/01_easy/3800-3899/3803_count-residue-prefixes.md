# 3803 — Count Residue Prefixes

## Deskripsi

**Soal:** [3803. Count Residue Prefixes](https://leetcode.com/problems/count-residue-prefixes/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) — at most 26 distinct chars

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3803: Count Residue Prefixes
// https://leetcode.com/problems/count-residue-prefixes/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountResiduePrefixes("abc"))
	fmt.Println(CountResiduePrefixes("dd"))
	fmt.Println(CountResiduePrefixes("bob"))
}

// Time: O(n)
// Space: O(1) — at most 26 distinct chars
func CountResiduePrefixes(s string) int {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[byte]bool)
	count := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		seen[s[i]] = true
		if len(seen) == (i+1)%3 {
			count++
		}
	}
	return count
}
```
