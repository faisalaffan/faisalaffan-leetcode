# 0058 — Length Of Last Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func LengthOfLastWord(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #58: Length of Last Word
// https://leetcode.com/problems/length-of-last-word/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func LengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}
	return length
}

func main() {
	fmt.Println(LengthOfLastWord("Hello World"))
	fmt.Println(LengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(LengthOfLastWord("luffy is still joyboy"))
}
```
