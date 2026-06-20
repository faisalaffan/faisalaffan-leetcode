# 0214 — Shortest Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestPalindrome(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #214: Shortest Palindrome
// https://leetcode.com/problems/shortest-palindrome/
// Difficulty: Hard

import "fmt"

func shortestPalindrome(s string) string {
	n := len(s)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return ""
	}

	rev := make([]byte, n)
	for i := 0; i < n; i++ {
		rev[i] = s[n-1-i]
	}

	combined := s + "#" + string(rev)
  // Alokasi slice integer
	lps := make([]int, len(combined))

	for i := 1; i < len(combined); i++ {
		j := lps[i-1]
		for j > 0 && combined[i] != combined[j] {
			j = lps[j-1]
		}
		if combined[i] == combined[j] {
			j++
		}
		lps[i] = j
	}

	palLen := lps[len(lps)-1]
	suffix := rev[:n-palLen]
	return string(suffix) + s
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
}
```
