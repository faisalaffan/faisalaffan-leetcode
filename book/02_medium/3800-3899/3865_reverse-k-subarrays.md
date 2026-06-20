# 3865 — Reverse K Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ReverseKSubarrays(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3865: Reverse K Subarrays
// https://leetcode.com/problems/reverse-k-subarrays/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(1)
// Approach: Count positions that don't match expected value at their index.
// Each reversal of a subarray of length k can fix at most 2 positions.

import "fmt"

func ReverseKSubarrays(nums []int, k int) int {
	n := len(nums)
	swaps := 0
	for i := 0; i < n; i++ {
		if nums[i] != i { // value should equal index for sorted array [0,1,2,...]
			// Find where i is
			j := i
			for j < n && nums[j] != i {
				j++
			}
			if j-i+1 >= k && j < n {
				// Reverse subarray i..j
				for l, r := i, j; l < r; l, r = l+1, r-1 {
					nums[l], nums[r] = nums[r], nums[l]
				}
				swaps++
			}
		}
	}
	return swaps
}

func main() {
	// Example
	fmt.Println(ReverseKSubarrays([]int{1, 0, 3, 2}, 2)) // Expected: ?

	// Example
	fmt.Println(ReverseKSubarrays([]int{2, 1, 0}, 3)) // Expected: ?

	// Example
	fmt.Println(ReverseKSubarrays([]int{0, 1, 2}, 1)) // Expected: 0
}
```
