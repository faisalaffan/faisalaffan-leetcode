# 3555 — Smallest Subarray To Sort In Every Sliding Window

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestSubarrayToSortInEverySlidingWindow(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3555: Smallest Subarray to Sort in Every Sliding Window
// https://leetcode.com/problems/smallest-subarray-to-sort-in-every-sliding-window/
// Difficulty: Medium [Paid]
// Complexity: O(n*log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", SmallestSubarrayToSortInEverySlidingWindow([]int{3, 2, 1}, 2))
	// Test case 2
	fmt.Println("Test 2:", SmallestSubarrayToSortInEverySlidingWindow([]int{1, 2, 3, 4}, 3))
	// Test case 3
	fmt.Println("Test 3:", SmallestSubarrayToSortInEverySlidingWindow([]int{4, 3, 2, 1}, 2))
}

func SmallestSubarrayToSortInEverySlidingWindow(nums []int, k int) int {
	n := len(nums)
	if k > n {
		k = n
	}
	// For each sliding window of size k, find the minimum length subarray
	// that when sorted makes the entire window sorted
	minLen := n
	for i := 0; i <= n-k; i++ {
  // Alokasi slice integer
		window := make([]int, k)
		copy(window, nums[i:i+k])
  // Alokasi slice integer
		sorted := make([]int, k)
		copy(sorted, window)
  // Urutkan secara ascending — O(n log n)
		sort.Ints(sorted)

		left, right := 0, k-1
		for left < k && window[left] == sorted[left] {
			left++
		}
		for right >= 0 && window[right] == sorted[right] {
			right--
		}
		if left <= right {
			length := right - left + 1
			if length < minLen {
				minLen = length
			}
		} else {
			return 0
		}
	}
	return minLen
}
```
