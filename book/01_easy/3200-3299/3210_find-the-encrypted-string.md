# 3210 — Find The Encrypted String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheEncryptedString(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3210: Find the Encrypted String
// https://leetcode.com/problems/find-the-encrypted-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheEncryptedString("dart", 3))
	fmt.Println(FindTheEncryptedString("aaa", 1))
	fmt.Println(FindTheEncryptedString("abcd", 5))
}

// FindTheEncryptedString returns the encrypted string by rotating each character by k positions forward.
// Time: O(n). Space: O(n).
func FindTheEncryptedString(s string, k int) string {
	n := len(s)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return s
	}
	k %= n
	return s[k:] + s[:k]
}
```
