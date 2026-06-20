# 1918 — Kth Smallest Subarray Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func KthSmallestSubarraySum(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search, Prefix Sum

**Kompleksitas Waktu:** O(n log sum), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1918: Kth Smallest Subarray Sum
// https://leetcode.com/problems/kth-smallest-subarray-sum/
// Difficulty: Medium [Paid]

import (
	"fmt"
)

func main() {
	fmt.Println(KthSmallestSubarraySum([]int{2, 1, 3}, 4))
	fmt.Println(KthSmallestSubarraySum([]int{3, 3, 3}, 4))
}

// Time: O(n log sum), Space: O(1)
func KthSmallestSubarraySum(nums []int, k int) int {
	n := len(nums)
  // Alokasi slice integer
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	totalSum := prefix[n]

	// Binary search on sum value
	left, right := 0, totalSum
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if countSubarraysLE(nums, prefix, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func countSubarraysLE(nums []int, prefix []int, target int) int {
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		// Find first j where prefix[j+1]-prefix[i] > target
		lo, hi := i, len(nums)-1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			sum := prefix[mid+1] - prefix[i]
			if sum <= target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		count += lo - i
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
