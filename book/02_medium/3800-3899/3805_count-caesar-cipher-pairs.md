# 3805 — Count Caesar Cipher Pairs

## Deskripsi

**Soal:** [3805. Count Caesar Cipher Pairs](https://leetcode.com/problems/count-caesar-cipher-pairs/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N * M)  
**Kompleksitas Ruang:** O(N * M)

**Algoritma:** HashMap (tabel pencarian O(1))

**Fungsi Solusi:** `func CountCaesarCipherPairs(words []string) int`

> **Ide Kunci:** Normalize each string by shifting so that its first character

## Solusi Go

```go
package main

// LeetCode #3805: Count Caesar Cipher Pairs
// https://leetcode.com/problems/count-caesar-cipher-pairs/
// Difficulty: Medium
// Time: O(N * M) | Space: O(N * M)
// Approach: Normalize each string by shifting so that its first character
// becomes 'a'. Strings in the same Caesar-shift equivalence class will have
// the same normalized form. Count pairs using hash map.

import "fmt"

func CountCaesarCipherPairs(words []string) int {
	normalize := func(s string) string {
		shift := int(s[0] - 'a')
  // Membuat slice untuk menyimpan hasil
		res := make([]byte, len(s))
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(s); i++ {
			res[i] = byte((int(s[i]-'a')-shift+26)%26 + 'a')
		}
		return string(res)
	}

  // Membuat map untuk pencarian O(1): key → value
	count := make(map[string]int)
	ans := 0

	for _, w := range words {
		norm := normalize(w)
		ans += count[norm]
		count[norm]++
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(CountCaesarCipherPairs([]string{"fusion", "layout"})) // Expected: 1

	// Example 2
	fmt.Println(CountCaesarCipherPairs([]string{"ab", "aa", "za", "aa"})) // Expected: 2

	// Example 3
	fmt.Println(CountCaesarCipherPairs([]string{"abc", "bcd", "cde", "xyz"})) // Expected: 3
}
```
