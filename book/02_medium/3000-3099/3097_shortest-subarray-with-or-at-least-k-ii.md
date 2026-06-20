# 3097 — Shortest Subarray With Or At Least K Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumSubarrayLength(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Bitmask

**Waktu:** O(n * 32) = O(n)  |  **Ruang:** O(32) = O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3097: Shortest Subarray With OR at Least K II
// https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/
// Difficulty: Medium
// Time: O(n * 32) = O(n) | Space: O(32) = O(1)

import "fmt"

func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}

	n := len(nums)
	ans := n + 1
  // Alokasi slice
	bits := make([]int, 32)
	left := 0
	cur := 0

	for right := 0; right < n; right++ {
		cur |= nums[right]
		for b := 0; b < 32; b++ {
			if nums[right]&(1<<b) != 0 {
				bits[b]++
			}
		}

  // Binary search loop
		for left <= right && cur >= k {
			if right-left+1 < ans {
				ans = right - left + 1
			}

			for b := 0; b < 32; b++ {
				if nums[left]&(1<<b) != 0 {
					bits[b]--
					if bits[b] == 0 {
						cur &^= (1 << b)
					}
				}
			}
			left++
		}
	}

	if ans > n {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minimumSubarrayLength([]int{1, 2, 3}, 3))       // Expected: 1
	fmt.Println(minimumSubarrayLength([]int{1, 2, 3}, 5))       // Expected: 2
	fmt.Println(minimumSubarrayLength([]int{2, 1, 8}, 10))      // Expected: 3
}
```
