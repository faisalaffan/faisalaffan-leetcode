# 1925 — Count Square Sum Triples

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSquareSumTriples(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1925: Count Square Sum Triples
// https://leetcode.com/problems/count-square-sum-triples/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSquareSumTriples(5))  // 2
	fmt.Println(CountSquareSumTriples(10)) // 4
}

// Time: O(n^2), Space: O(1)
func CountSquareSumTriples(n int) int {
	count := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			c2 := a*a + b*b
			c := 1
			for c*c < c2 {
				c++
			}
			if c*c == c2 && c <= n {
				count++
			}
		}
	}
	return count
}
```
