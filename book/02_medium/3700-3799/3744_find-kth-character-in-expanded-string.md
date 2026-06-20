# 3744 — Find Kth Character In Expanded String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func findKthCharacterInExpandedString(s string, k int) byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3744: Find Kth Character in Expanded String
// https://leetcode.com/problems/find-kth-character-in-expanded-string/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func findKthCharacterInExpandedString(s string, k int) byte {
	words := strings.Fields(s)

	for _, word := range words {
		l := len(word)
		m := l * (l + 1) / 2

		if k == m {
			return ' '
		} else if k > m {
			k -= (m + 1)
			continue
		} else {
			cur := 0
			for i, ch := range word {
				cur += (i + 1)
				if k < cur {
					return byte(ch)
				}
			}
			return ' '
		}
	}
	return ' '
}

func main() {
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 0))
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 15))
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 20))
}
```
