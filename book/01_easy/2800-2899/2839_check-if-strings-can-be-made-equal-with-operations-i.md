# 2839 — Check If Strings Can Be Made Equal With Operations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfStringsCanBeMadeEqualWithOperationsI(s1 string, s2 string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2839: Check if Strings Can be Made Equal With Operations I
// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-i/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsI("abcd", "cdab"))
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsI("abcd", "dacb"))
}

func CheckIfStringsCanBeMadeEqualWithOperationsI(s1 string, s2 string) bool {
	// Can swap characters at even indices (0<->2) and odd indices (1<->3)
	// Check that multiset of chars at same parity positions match
	return (s1[0] == s2[0] || s1[0] == s2[2]) &&
		(s1[2] == s2[2] || s1[2] == s2[0]) &&
		(s1[1] == s2[1] || s1[1] == s2[3]) &&
		(s1[3] == s2[3] || s1[3] == s2[1])
}
```
