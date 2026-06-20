# 3805 — Count Caesar Cipher Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountCaesarCipherPairs(words []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(N * M)  |  **Ruang:** O(N * M)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
		res := make([]byte, len(s))
  // Linear scan O(n)
		for i := 0; i < len(s); i++ {
			res[i] = byte((int(s[i]-'a')-shift+26)%26 + 'a')
		}
		return string(res)
	}

  // HashMap: O(1) lookup
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
