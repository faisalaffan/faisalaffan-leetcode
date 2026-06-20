# 0880 — Decoded String At Index

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func DecodedStringAtIndex(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #880: Decoded String at Index
// https://leetcode.com/problems/decoded-string-at-index/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(DecodedStringAtIndex("leet2code3", 10))
	fmt.Println(DecodedStringAtIndex("ha22", 5))
	fmt.Println(DecodedStringAtIndex("a2345678999999999999999", 1))
}

// Time: O(n) | Space: O(1)
func DecodedStringAtIndex(s string, k int) string {
	var size int64 = 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			size++
		} else {
			size *= int64(s[i] - '0')
		}
	}

	target := int64(k)
	for i := len(s) - 1; i >= 0; i-- {
		target %= size
		if target == 0 && s[i] >= 'a' && s[i] <= 'z' {
			return string(s[i])
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			size--
		} else {
			size /= int64(s[i] - '0')
		}
	}

	return ""
}
```
