# 1498 — Number Of Subsequences That Satisfy The Given Sum Condition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumSubseq(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(N log N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1498: Number of Subsequences That Satisfy the Given Sum Condition
// https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumSubseq([]int{3, 5, 6, 7}, 9))
	fmt.Println(NumSubseq([]int{3, 3, 6, 8}, 10))
	fmt.Println(NumSubseq([]int{2, 3, 3, 4, 6, 7}, 12))
}

func NumSubseq(nums []int, target int) int {
	// Time: O(N log N), Space: O(N)
	const mod = 1_000_000_007

	// Sort nums
  // Alokasi slice integer
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	quickSort(sorted, 0, len(sorted)-1)

	// Precompute powers of 2
  // Alokasi slice integer
	pow := make([]int, len(sorted))
	pow[0] = 1
	for i := 1; i < len(sorted); i++ {
		pow[i] = (pow[i-1] * 2) % mod
	}

	count := 0
	left, right := 0, len(sorted)-1

	for left <= right {
		if sorted[left]+sorted[right] <= target {
			// All subsequences with sorted[left] as min and any subset of elements between left+1..right
			count = (count + pow[right-left]) % mod
			left++
		} else {
			right--
		}
	}

	return count
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
```
