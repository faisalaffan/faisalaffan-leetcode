# 1684 — Count The Number Of Consistent Strings

## Deskripsi

**Soal:** [1684. Count The Number Of Consistent Strings](https://leetcode.com/problems/count-the-number-of-consistent-strings/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + m*k), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CountConsistentStrings(allowed string, words []string) int`

## Solusi Go

```go
package main

// LeetCode #1684: Count the Number of Consistent Strings
// https://leetcode.com/problems/count-the-number-of-consistent-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n + m*k), Space: O(1)
func CountConsistentStrings(allowed string, words []string) int {
  // Membuat map untuk pencarian O(1): key → value
	allowedSet := make(map[byte]bool)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(allowed); i++ {
		allowedSet[allowed[i]] = true
	}
	count := 0
	for _, word := range words {
		consistent := true
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(word); i++ {
			if !allowedSet[word[i]] {
				consistent = false
				break
			}
		}
		if consistent {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(CountConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	fmt.Println(CountConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}
```
