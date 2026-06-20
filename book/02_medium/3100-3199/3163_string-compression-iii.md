# 3163 — String Compression Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func compressedString(word string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3163: String Compression III
// https://leetcode.com/problems/string-compression-iii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func compressedString(word string) string {
	ans := make([]byte, 0, len(word)*2)
	n := len(word)
	i := 0

	for i < n {
		ch := word[i]
		j := i
		for j < n && j-i < 9 && word[j] == ch {
			j++
		}
		ans = append(ans, byte('0'+j-i), ch)
		i = j
	}
	return string(ans)
}

func main() {
	fmt.Println(compressedString("abcde"))               // Expected: "1a1b1c1d1e"
	fmt.Println(compressedString("aaaaaaaaaaaaaabb"))     // Expected: "9a5a2b"
}
```
