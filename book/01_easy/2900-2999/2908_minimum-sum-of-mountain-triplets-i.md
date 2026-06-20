# 2908 — Minimum Sum Of Mountain Triplets I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumSumOfMountainTripletsI(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2908: Minimum Sum of Mountain Triplets I
// https://leetcode.com/problems/minimum-sum-of-mountain-triplets-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumSum
	fmt.Println(MinimumSumOfMountainTripletsI([]int{8, 6, 1, 5, 3})) // 9
	fmt.Println(MinimumSumOfMountainTripletsI([]int{5, 4, 8, 7, 10, 2})) // 13
	fmt.Println(MinimumSumOfMountainTripletsI([]int{6, 5, 4, 3, 4, 5})) // -1
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: minimumSum
func MinimumSumOfMountainTripletsI(nums []int) int {
	n := len(nums)
	if n < 3 {
		return -1
	}

	// prefix[i] = min value in nums[0..i]
  // Alokasi slice integer
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < prefix[i-1] {
			prefix[i] = nums[i]
		} else {
			prefix[i] = prefix[i-1]
		}
	}

	// suffix[i] = min value in nums[i..n-1]
  // Alokasi slice integer
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffix[i+1] {
			suffix[i] = nums[i]
		} else {
			suffix[i] = suffix[i+1]
		}
	}

	minSum := -1
	for j := 1; j < n-1; j++ {
		if prefix[j-1] < nums[j] && suffix[j+1] < nums[j] {
			sum := prefix[j-1] + nums[j] + suffix[j+1]
			if minSum == -1 || sum < minSum {
				minSum = sum
			}
		}
	}
	return minSum
}
```
