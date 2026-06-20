# 1704 — Determine If String Halves Are Alike

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func HalvesAreAlike(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1704: Determine if String Halves Are Alike
// https://leetcode.com/problems/determine-if-string-halves-are-alike/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func HalvesAreAlike(s string) bool {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	mid := len(s) / 2
	count := 0
	for i := 0; i < mid; i++ {
		if vowels[s[i]] {
			count++
		}
		if vowels[s[i+mid]] {
			count--
		}
	}
	return count == 0
}

func main() {
	fmt.Println(HalvesAreAlike("book"))
	fmt.Println(HalvesAreAlike("textbook"))
}
```
