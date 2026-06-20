# 0616 — Add Bold Tag In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func AddBoldTag(s string, words []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L) where n = len(s), L = total length of all words  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #616: Add Bold Tag in String
// https://leetcode.com/problems/add-bold-tag-in-string/
// Difficulty: Medium [Paid]
// Time: O(n * L) where n = len(s), L = total length of all words
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(AddBoldTag("abcxyz123", []string{"abc", "123"}))
	fmt.Println(AddBoldTag("aaabbcc", []string{"aaa", "aab", "bc"}))
}

func AddBoldTag(s string, words []string) string {
	n := len(s)
	bold := make([]bool, n)

	for _, word := range words {
		for i := 0; i <= n-len(word); i++ {
			if s[i:i+len(word)] == word {
				for j := i; j < i+len(word); j++ {
					bold[j] = true
				}
			}
		}
	}

	result := ""
	i := 0
	for i < n {
		if bold[i] {
			result += "<b>"
			for i < n && bold[i] {
				result += string(s[i])
				i++
			}
			result += "</b>"
		} else {
			result += string(s[i])
			i++
		}
	}

	return result
}
```
