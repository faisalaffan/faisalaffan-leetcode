# 3136 — Valid Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ValidWord(word string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3136: Valid Word
// https://leetcode.com/problems/valid-word/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isValid
	fmt.Println(ValidWord("Hello123")) // true
	fmt.Println(ValidWord("hi"))       // false
	fmt.Println(ValidWord("AAab"))     // false (no vowel)
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: isValid
func ValidWord(word string) bool {
	if len(word) < 3 {
		return false
	}
	hasVowel := false
	hasConsonant := false
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= '0' && c <= '9' {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
		}
		if c < 'a' || c > 'z' {
			return false
		}
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			hasVowel = true
		} else {
			hasConsonant = true
		}
	}
	return hasVowel && hasConsonant
}
```
