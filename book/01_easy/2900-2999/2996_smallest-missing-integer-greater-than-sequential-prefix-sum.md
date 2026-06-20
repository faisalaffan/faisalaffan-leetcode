# 2996 — Smallest Missing Integer Greater Than Sequential Prefix Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestMissingIntegerGreaterThanSequentialPrefixSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2996: Smallest Missing Integer Greater Than Sequential Prefix Sum
// https://leetcode.com/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: missingInteger
	fmt.Println(SmallestMissingIntegerGreaterThanSequentialPrefixSum([]int{1, 2, 3, 2, 5}))        // 6
	fmt.Println(SmallestMissingIntegerGreaterThanSequentialPrefixSum([]int{3, 4, 5, 1, 12, 14, 13})) // 15
	fmt.Println(SmallestMissingIntegerGreaterThanSequentialPrefixSum([]int{4, 5, 6, 7, 8, 9, 10, 11})) // 60
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: missingInteger
func SmallestMissingIntegerGreaterThanSequentialPrefixSum(nums []int) int {
	// Find longest sequential prefix (each element = prev + 1)
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			sum += nums[i]
		} else {
			break
		}
	}

	// Find smallest missing >= sum
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	for _, v := range nums {
		seen[v] = true
	}
	for seen[sum] {
		sum++
	}
	return sum
}
```
