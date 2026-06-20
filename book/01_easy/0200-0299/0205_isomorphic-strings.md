# 0205 — Isomorphic Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func IsIsomorphic(s string, t string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (fixed ASCII chars)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #205: Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (fixed ASCII chars)
func IsIsomorphic(s string, t string) bool {
  // Alokasi slice integer
	m1 := make([]int, 256)
  // Alokasi slice integer
	m2 := make([]int, 256)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

func main() {
	fmt.Println(IsIsomorphic("egg", "add"))
	fmt.Println(IsIsomorphic("foo", "bar"))
	fmt.Println(IsIsomorphic("paper", "title"))
}
```
