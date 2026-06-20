# 1963 — Minimum Number Of Swaps To Make The String Balanced

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinSwapsBalanced(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1963: Minimum Number of Swaps to Make the String Balanced
// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSwapsBalanced("][]["))
	fmt.Println(MinSwapsBalanced("]]][[["))
	fmt.Println(MinSwapsBalanced("[]"))
}

// Time: O(n), Space: O(1)
func MinSwapsBalanced(s string) int {
	unmatched := 0
	maxUnmatched := 0
	for _, c := range s {
		if c == '[' {
			unmatched--
		} else {
			unmatched++
		}
		if unmatched > maxUnmatched {
			maxUnmatched = unmatched
		}
	}
	return (maxUnmatched + 1) / 2
}
```
