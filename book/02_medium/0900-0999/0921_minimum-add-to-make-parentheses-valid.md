# 0921 — Minimum Add To Make Parentheses Valid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func minAddToMakeValid(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #921: Minimum Add to Make Parentheses Valid
// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(1)
func minAddToMakeValid(s string) int {
	open, add := 0, 0
	for _, ch := range s {
		if ch == '(' {
			open++
		} else {
			if open > 0 {
				open--
			} else {
				add++
			}
		}
	}
	return add + open
}

func main() {
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
	fmt.Println(minAddToMakeValid("()"))
	fmt.Println(minAddToMakeValid("()))(("))
}
```
