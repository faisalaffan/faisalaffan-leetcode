# 2600 — K Items With The Maximum Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func KItemsWithTheMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2600: K Items With the Maximum Sum
// https://leetcode.com/problems/k-items-with-the-maximum-sum/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(KItemsWithTheMaximumSum(3, 2, 0, 2)) // 2
	fmt.Println(KItemsWithTheMaximumSum(3, 2, 0, 4)) // 3
}

func KItemsWithTheMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k <= numOnes {
		return k
	}
	k -= numOnes
	if k <= numZeros {
		return numOnes
	}
	k -= numZeros
	return numOnes - k
}
```
