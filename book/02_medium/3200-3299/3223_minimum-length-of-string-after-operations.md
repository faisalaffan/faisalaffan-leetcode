# 3223 — Minimum Length Of String After Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumLength(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(26) = O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3223: Minimum Length of String After Operations
// https://leetcode.com/problems/minimum-length-of-string-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(26) = O(1)

import "fmt"

func minimumLength(s string) int {
  // Alokasi slice integer
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	ans := 0
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if f%2 == 0 {
			ans += 2
		} else {
			ans += 1
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumLength("abaacbcbb")) // Expected: 5
	fmt.Println(minimumLength("aa"))         // Expected: 2
}
```
