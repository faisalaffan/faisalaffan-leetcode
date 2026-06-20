# 0477 — Total Hamming Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func TotalHammingDistance(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Bitmask

**Kompleksitas Waktu:** O(n * 32) = O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Bitmask** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #477: Total Hamming Distance
// https://leetcode.com/problems/total-hamming-distance/
// Difficulty: Medium
// Time: O(n * 32) = O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(TotalHammingDistance([]int{4, 14, 2}))
	fmt.Println(TotalHammingDistance([]int{4, 14, 4}))
}

func TotalHammingDistance(nums []int) int {
	total := 0
	n := len(nums)

	for bit := 0; bit < 32; bit++ {
		countOnes := 0
		for _, num := range nums {
			if num&(1<<bit) != 0 {
				countOnes++
			}
		}
		countZeros := n - countOnes
		total += countOnes * countZeros
	}

	return total
}
```
