# 3823 — Reverse Letters Then Special Characters In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func isLetter(ch byte) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3823: Reverse Letters Then Special Characters in a String
// https://leetcode.com/problems/reverse-letters-then-special-characters-in-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReverseLettersThenSpecialCharactersInAString(")ebc#da@f("))
	fmt.Println(ReverseLettersThenSpecialCharactersInAString("z"))
	fmt.Println(ReverseLettersThenSpecialCharactersInAString("!@#$%^&*()"))
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func isSpecial(ch byte) bool {
	return !isLetter(ch)
}

func reverseRange(s []byte, cond func(byte) bool) {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !cond(s[i]) {
			i++
		}
		for i < j && !cond(s[j]) {
			j--
		}
		if i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
}

// Time: O(n)
// Space: O(n)
func ReverseLettersThenSpecialCharactersInAString(s string) string {
	b := []byte(s)
	reverseRange(b, isLetter)
	reverseRange(b, isSpecial)
	return string(b)
}
```
