# 1679 — Max Number Of K Sum Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxOperations(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1679: Max Number of K-Sum Pairs
// https://leetcode.com/problems/max-number-of-k-sum-pairs/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maxOperations(nums []int, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[int]int)
	ops := 0

	for _, num := range nums {
		complement := k - num
		if counts[complement] > 0 {
			ops++
			counts[complement]--
		} else {
			counts[num]++
		}
	}
	return ops
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxOperations([]int{1, 2, 3, 4}, 5)) // Expected: 2

	// Test case 2
	fmt.Println("Test 2:", maxOperations([]int{3, 1, 3, 4, 3}, 6)) // Expected: 1

	// Test case 3
	fmt.Println("Test 3:", maxOperations([]int{1, 2, 3, 4, 5, 6}, 7)) // Expected: 3
}
```
