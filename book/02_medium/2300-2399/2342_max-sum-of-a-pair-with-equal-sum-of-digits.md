# 2342 — Max Sum Of A Pair With Equal Sum Of Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * d)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2342: Max Sum of a Pair With Equal Sum of Digits
// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/
// Difficulty: Medium
// Time: O(n * d) | Space: O(n)

import "fmt"

func maximumSum(nums []int) int {
  // Alokasi slice integer
	digitSums := make([]int, len(nums))
	for i, v := range nums {
		sum := 0
		for v > 0 {
			sum += v % 10
			v /= 10
		}
		digitSums[i] = sum
	}

  // Membuat map (HashMap) — pencarian O(1)
	maxVal := make(map[int][2]int) // top 2 for each digit sum
	for i, ds := range digitSums {
		a, b := maxVal[ds][0], maxVal[ds][1]
		if nums[i] > a {
			maxVal[ds] = [2]int{nums[i], a}
		} else if nums[i] > b {
			maxVal[ds] = [2]int{a, nums[i]}
		}
	}

	maxSum := -1
	for _, vals := range maxVal {
		if vals[1] > 0 {
			s := vals[0] + vals[1]
			if s > maxSum {
				maxSum = s
			}
		}
	}
	return maxSum
}

func main() {
	// Test case 1
	fmt.Println(maximumSum([]int{18, 43, 36, 13, 7}))
	// Expected: 54

	// Test case 2
	fmt.Println(maximumSum([]int{10, 12, 19, 14}))
	// Expected: -1

	// Test case 3
	fmt.Println(maximumSum([]int{229, 398, 269, 317, 420, 464, 491, 218, 439, 153, 482, 169, 411, 93, 147, 50, 347, 210, 251, 366, 401}))
	// Expected: 973
}
```
