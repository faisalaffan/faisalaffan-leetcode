# 1936 — Add Minimum Number Of Rungs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func AddRungs(rungs []int, dist int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1936: Add Minimum Number of Rungs
// https://leetcode.com/problems/add-minimum-number-of-rungs/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(AddRungs([]int{1, 3, 5, 10}, 2))
	fmt.Println(AddRungs([]int{3, 6, 8, 10}, 3))
	fmt.Println(AddRungs([]int{3, 4, 6, 7}, 2))
}

// Time: O(n), Space: O(1)
func AddRungs(rungs []int, dist int) int {
	count := 0
	prev := 0
	for _, r := range rungs {
		gap := r - prev
		if gap > dist {
			count += (gap - 1) / dist
		}
		prev = r
	}
	return count
}
```
