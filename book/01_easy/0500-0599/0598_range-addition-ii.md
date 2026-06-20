# 0598 — Range Addition Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func RangeAdditionIi(m, n int, ops [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(k), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #598: Range Addition II
// https://leetcode.com/problems/range-addition-ii/
// Difficulty: Easy

import "fmt"

// Time: O(k), Space: O(1)
func RangeAdditionIi(m, n int, ops [][]int) int {
	minA, minB := m, n
	for _, op := range ops {
		if op[0] < minA {
			minA = op[0]
		}
		if op[1] < minB {
			minB = op[1]
		}
	}
	return minA * minB
}

func main() {
	fmt.Println(RangeAdditionIi(3, 3, [][]int{{2, 2}, {3, 3}}))
	fmt.Println(RangeAdditionIi(3, 3, [][]int{}))
}
```
