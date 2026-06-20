# 1119 — Remove Vowels From A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func removeVowels(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1119: Remove Vowels from a String
// https://leetcode.com/problems/remove-vowels-from-a-string/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeVowels("leetcodeisacommunityforcoders")) // "ltcdscmmntyfrcdrs"
	fmt.Println(removeVowels("aeiou"))                         // ""
}

// LeetCode submission: removeVowels
func removeVowels(s string) string {
	ans := make([]byte, 0, len(s))
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
```
