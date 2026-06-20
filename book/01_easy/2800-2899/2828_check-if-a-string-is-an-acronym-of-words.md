# 2828 — Check If A String Is An Acronym Of Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfAStringIsAnAcronymOfWords(words []string, s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2828: Check if a String Is an Acronym of Words
// https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfAStringIsAnAcronymOfWords([]string{"alice", "bob", "charlie"}, "abc"))
	fmt.Println(CheckIfAStringIsAnAcronymOfWords([]string{"an", "apple"}, "a"))
}

func CheckIfAStringIsAnAcronymOfWords(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i, w := range words {
		if w[0] != s[i] {
			return false
		}
	}
	return true
}
```
