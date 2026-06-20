# 1957 — Delete Characters To Make Fancy String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func DeleteCharactersToMakeFancyString(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1957: Delete Characters to Make Fancy String
// https://leetcode.com/problems/delete-characters-to-make-fancy-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DeleteCharactersToMakeFancyString("leeetcode"))     // "leetcode"
	fmt.Println(DeleteCharactersToMakeFancyString("aaabaaaa"))      // "aabaa"
	fmt.Println(DeleteCharactersToMakeFancyString("aab"))           // "aab"
}

// Time: O(n), Space: O(n)
func DeleteCharactersToMakeFancyString(s string) string {
	result := make([]byte, 0, len(s))
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		n := len(result)
		if n >= 2 && result[n-1] == s[i] && result[n-2] == s[i] {
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}
```
