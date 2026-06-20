# 3701 — Compute Alternating Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func ComputeAlternatingSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3701: Compute Alternating Sum
// https://leetcode.com/problems/compute-alternating-sum/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ComputeAlternatingSum([]int{1, 2, 3, 4, 5}))
	fmt.Println(ComputeAlternatingSum([]int{10, 5, 3}))
}

// Time: O(n)
// Space: O(1)
func ComputeAlternatingSum(nums []int) int {
	ans := 0
	for i, x := range nums {
		if i%2 == 0 {
			ans += x
		} else {
			ans -= x
		}
	}
	return ans
}
```
