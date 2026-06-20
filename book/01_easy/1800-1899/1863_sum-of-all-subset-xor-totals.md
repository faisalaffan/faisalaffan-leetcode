# 1863 — Sum Of All Subset Xor Totals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SubsetXORSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Stack

**Kompleksitas Waktu:** O(2^n), Space: O(n) (recursion stack)  
**Kompleksitas Ruang:** O(n) (recursion stack)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1863: Sum of All Subset XOR Totals
// https://leetcode.com/problems/sum-of-all-subset-xor-totals/
// Difficulty: Easy

import "fmt"

// Time: O(2^n), Space: O(n) (recursion stack)
func SubsetXORSum(nums []int) int {
	return dfs(nums, 0, 0)
}

func dfs(nums []int, idx int, currentXor int) int {
	if idx == len(nums) {
		return currentXor
	}
	// Include nums[idx] or skip it
	return dfs(nums, idx+1, currentXor^nums[idx]) + dfs(nums, idx+1, currentXor)
}

func main() {
	fmt.Println(SubsetXORSum([]int{1, 3}))
	fmt.Println(SubsetXORSum([]int{5, 1, 6}))
	fmt.Println(SubsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}
```
