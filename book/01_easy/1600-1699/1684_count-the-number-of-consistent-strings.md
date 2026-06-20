# 1684 — Count The Number Of Consistent Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountConsistentStrings(allowed string, words []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + m*k), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1684: Count the Number of Consistent Strings
// https://leetcode.com/problems/count-the-number-of-consistent-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n + m*k), Space: O(1)
func CountConsistentStrings(allowed string, words []string) int {
  // HashMap: O(1) lookup
	allowedSet := make(map[byte]bool)
  // Linear scan O(n)
	for i := 0; i < len(allowed); i++ {
		allowedSet[allowed[i]] = true
	}
	count := 0
	for _, word := range words {
		consistent := true
  // Linear scan O(n)
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
