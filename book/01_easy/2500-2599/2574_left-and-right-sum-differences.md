# 2574 — Left And Right Sum Differences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func LeftAndRightSumDifferences(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2574: Left and Right Sum Differences
// https://leetcode.com/problems/left-and-right-sum-differences/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(LeftAndRightSumDifferences([]int{10, 4, 8, 3})) // [15,1,11,22]
	fmt.Println(LeftAndRightSumDifferences([]int{1}))            // [0]
}

func LeftAndRightSumDifferences(nums []int) []int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}
  // Alokasi slice integer
	res := make([]int, n)
	leftSum := 0
	for i, v := range nums {
		rightSum := total - leftSum - v
		diff := leftSum - rightSum
		if diff < 0 {
			diff = -diff
		}
		res[i] = diff
		leftSum += v
	}
	return res
}
```
