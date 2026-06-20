# 0917 — Reverse Only Letters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func reverseOnlyLetters(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #917: Reverse Only Letters
// https://leetcode.com/problems/reverse-only-letters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))      // "dc-ba"
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj")) // "j-Ih-gfE-dCba"
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!")) // "Qedo1ct-eeLg=ntse-T!"
}

// reverseOnlyLetters reverses only the letters in the string, keeping non-letters in place.
// Time: O(n). Space: O(n).
func reverseOnlyLetters(s string) string {
	b := []byte(s)
	l, r := 0, len(b)-1
  // Two-pointer: gerakkan kiri atau kanan
	for l < r {
		if !isLetter(b[l]) {
			l++
			continue
		}
		if !isLetter(b[r]) {
			r--
			continue
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
```
