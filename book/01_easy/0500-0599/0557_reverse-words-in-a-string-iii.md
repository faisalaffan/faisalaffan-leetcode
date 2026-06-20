# 0557 — Reverse Words In A String Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ReverseWordsInAStringIii(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #557: Reverse Words in a String III
// https://leetcode.com/problems/reverse-words-in-a-string-iii/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseWordsInAStringIii(s string) string {
	b := []byte(s)
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			for lo, hi := start, i-1; lo < hi; lo, hi = lo+1, hi-1 {
				b[lo], b[hi] = b[hi], b[lo]
			}
			start = i + 1
		}
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseWordsInAStringIii("Let's take LeetCode contest"))
	fmt.Println(ReverseWordsInAStringIii("Mr Ding"))
}
```
