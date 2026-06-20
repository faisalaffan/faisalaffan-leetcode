# 1974 — Minimum Time To Type Word Using Special Typewriter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumTimeToTypeWordUsingSpecialTypewriter(word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1974: Minimum Time to Type Word Using Special Typewriter
// https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("abc"))  // 5
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("bza"))  // 7
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("zjpc")) // 34
}

// Time: O(n), Space: O(1)
func MinimumTimeToTypeWordUsingSpecialTypewriter(word string) int {
	seconds := 0
	pos := 0 // 'a'
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
		target := int(word[i] - 'a')
		diff := target - pos
		if diff < 0 {
			diff = -diff
		}
		if diff > 13 {
			diff = 26 - diff
		}
		seconds += diff + 1 // move + type
		pos = target
	}
	return seconds
}
```
