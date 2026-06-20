# 1592 — Rearrange Spaces Between Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func reorderSpaces(text string) string

import (
	"fmt"
	"strings"
)

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1592: Rearrange Spaces Between Words
// https://leetcode.com/problems/rearrange-spaces-between-words/
// Difficulty: Easy
//
// LeetCode submission: func reorderSpaces(text string) string

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RearrangeSpacesBetweenWords("  this   is  a sentence ")) // "this   is   a   sentence"
	fmt.Println(RearrangeSpacesBetweenWords(" practice   makes   perfect")) // "practice   makes   perfect "
}

// Time: O(n), Space: O(n)
func RearrangeSpacesBetweenWords(text string) string {
	words := strings.Fields(text)
	spaces := strings.Count(text, " ")
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", spaces)
	}
	between := spaces / (len(words) - 1)
	extra := spaces % (len(words) - 1)
	res := strings.Join(words, strings.Repeat(" ", between))
	res += strings.Repeat(" ", extra)
	return res
}
```
