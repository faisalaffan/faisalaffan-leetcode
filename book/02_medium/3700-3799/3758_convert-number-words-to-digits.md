# 3758 — Convert Number Words To Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func convertNumberWordsToDigits(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3758: Convert Number Words to Digits
// https://leetcode.com/problems/convert-number-words-to-digits/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func convertNumberWordsToDigits(s string) string {
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var res strings.Builder
	i := 0
	for i < len(s) {
		found := false
		for d, word := range words {
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				res.WriteByte(byte(d) + '0')
				i += len(word)
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}
	return res.String()
}

func main() {
	fmt.Println(convertNumberWordsToDigits("onefourthree"))
	fmt.Println(convertNumberWordsToDigits("ninexsix"))
	fmt.Println(convertNumberWordsToDigits("zeero"))
}
```
