# 0266 — Palindrome Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CanPermutePalindrome(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (fixed 256 chars)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #266: Palindrome Permutation
// https://leetcode.com/problems/palindrome-permutation/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(1) (fixed 256 chars)
func CanPermutePalindrome(s string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	oddCount := 0
	for _, v := range count {
		if v%2 == 1 {
			oddCount++
		}
	}
	return oddCount <= 1
}

func main() {
	fmt.Println(CanPermutePalindrome("code"))
	fmt.Println(CanPermutePalindrome("aab"))
	fmt.Println(CanPermutePalindrome("carerac"))
}
```
