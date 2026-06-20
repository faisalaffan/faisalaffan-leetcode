# 0494 — Target Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func TargetSum(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * sum)  
**Kompleksitas Ruang:** O(sum)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #494: Target Sum
// https://leetcode.com/problems/target-sum/
// Difficulty: Medium
// Time: O(n * sum)
// Space: O(sum)

import "fmt"

func main() {
	fmt.Println(TargetSum([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(TargetSum([]int{1}, 1))
}

func TargetSum(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < target || (sum-target)%2 != 0 {
		return 0
	}
	negSum := (sum - target) / 2

  // Alokasi slice integer
	dp := make([]int, negSum+1)
	dp[0] = 1
	for _, num := range nums {
		for s := negSum; s >= num; s-- {
			dp[s] += dp[s-num]
		}
	}
	return dp[negSum]
}
```
