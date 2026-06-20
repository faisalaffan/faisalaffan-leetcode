# 1119 — Remove Vowels From A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func removeVowels(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1119: Remove Vowels from a String
// https://leetcode.com/problems/remove-vowels-from-a-string/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeVowels("leetcodeisacommunityforcoders")) // "ltcdscmmntyfrcdrs"
	fmt.Println(removeVowels("aeiou"))                         // ""
}

// LeetCode submission: removeVowels
func removeVowels(s string) string {
	ans := make([]byte, 0, len(s))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
```
