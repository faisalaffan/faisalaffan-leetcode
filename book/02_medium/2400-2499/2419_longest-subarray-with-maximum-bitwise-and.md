# 2419 — Longest Subarray With Maximum Bitwise And

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestSubarray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2419: Longest Subarray With Maximum Bitwise AND
// https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Max AND in any subarray is just the max element. Find longest consecutive
// subarray where all elements equal the max element.

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 2, 3, 3, 2, 2})) // 2
	fmt.Println(longestSubarray([]int{1, 2, 3, 4}))        // 1
}

func longestSubarray(nums []int) int {
	mx := 0
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}
	ans, cur := 0, 0
	for _, v := range nums {
		if v == mx {
			cur++
		} else {
			cur = 0
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```
