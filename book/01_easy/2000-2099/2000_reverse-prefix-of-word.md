# 2000 — Reverse Prefix Of Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ReversePrefixOfWord(word string, ch byte) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2000: Reverse Prefix of Word
// https://leetcode.com/problems/reverse-prefix-of-word/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReversePrefixOfWord("abcdefd", 'd')) // "dcbaefd"
	fmt.Println(ReversePrefixOfWord("xyxzxe", 'z'))  // "zxyxxe"
	fmt.Println(ReversePrefixOfWord("abcd", 'z'))    // "abcd"
}

// Time: O(n), Space: O(n)
func ReversePrefixOfWord(word string, ch byte) string {
	idx := -1
  // Linear scan O(n)
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			idx = i
			break
		}
	}
	if idx == -1 {
		return word
	}

	result := make([]byte, len(word))
	for i := 0; i <= idx; i++ {
		result[i] = word[idx-i]
	}
	for i := idx + 1; i < len(word); i++ {
		result[i] = word[i]
	}
	return string(result)
}
```
