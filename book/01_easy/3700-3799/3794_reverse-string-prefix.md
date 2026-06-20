# 3794 — Reverse String Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ReverseStringPrefix(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3794: Reverse String Prefix
// https://leetcode.com/problems/reverse-string-prefix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReverseStringPrefix("abcd", 2))
	fmt.Println(ReverseStringPrefix("xyz", 3))
	fmt.Println(ReverseStringPrefix("hey", 1))
}

// Time: O(n)
// Space: O(n)
func ReverseStringPrefix(s string, k int) string {
	b := []byte(s)
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
```
