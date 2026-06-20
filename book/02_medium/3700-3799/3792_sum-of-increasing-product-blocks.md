# 3792 — Sum Of Increasing Product Blocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumOfIncreasingProductBlocks(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3792: Sum of Increasing Product Blocks
// https://leetcode.com/problems/sum-of-increasing-product-blocks/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

const mod3792 = 1000000007

func sumOfIncreasingProductBlocks(n int) int {
	ans := 0
	k := 1
	for i := 1; i <= n; i++ {
		prod := 1
		for j := k; j < k+i; j++ {
			prod = (prod * j) % mod3792
		}
		ans = (ans + prod) % mod3792
		k += i
	}
	return ans
}

func main() {
	fmt.Println(sumOfIncreasingProductBlocks(3))
	fmt.Println(sumOfIncreasingProductBlocks(7))
	fmt.Println(sumOfIncreasingProductBlocks(1))
}
```
