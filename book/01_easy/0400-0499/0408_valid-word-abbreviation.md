# 0408 — Valid Word Abbreviation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ValidWordAbbreviation(word, abbr string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #408: Valid Word Abbreviation
// https://leetcode.com/problems/valid-word-abbreviation/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n), Space: O(1)
func ValidWordAbbreviation(word, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if abbr[j] >= 'a' && abbr[j] <= 'z' {
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
			continue
		}
		if abbr[j] == '0' {
			return false
		}
		num := 0
		for j < len(abbr) && abbr[j] >= '0' && abbr[j] <= '9' {
			num = num*10 + int(abbr[j]-'0')
			j++
		}
		i += num
	}
	return i == len(word) && j == len(abbr)
}

func main() {
	fmt.Println(ValidWordAbbreviation("internationalization", "i12iz4n"))
	fmt.Println(ValidWordAbbreviation("apple", "a2e"))
	fmt.Println(ValidWordAbbreviation("hi", "1"))
}
```
