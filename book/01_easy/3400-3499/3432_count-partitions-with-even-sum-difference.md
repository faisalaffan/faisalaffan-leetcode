# 3432 — Count Partitions With Even Sum Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountPartitionsWithEvenSumDifference(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3432: Count Partitions with Even Sum Difference
// https://leetcode.com/problems/count-partitions-with-even-sum-difference/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{10, 10, 10, 10, 10}))
}

// CountPartitionsWithEvenSumDifference counts partitions where the difference between left and right sums is even.
// Time: O(n). Space: O(1).
func CountPartitionsWithEvenSumDifference(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	leftSum := 0
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		rightSum := totalSum - leftSum
		if (leftSum-rightSum)%2 == 0 {
			count++
		}
	}
	return count
}
```
