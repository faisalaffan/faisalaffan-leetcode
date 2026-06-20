# 3011 — Find If Array Can Be Sorted

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func canSortArray(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3011: Find if Array Can Be Sorted
// https://leetcode.com/problems/find-if-array-can-be-sorted/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(canSortArray([]int{8, 4, 2, 30, 15}))
	fmt.Println(canSortArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(canSortArray([]int{3, 16, 8, 4, 2}))
}

func canSortArray(nums []int) bool {
	prevMax, curMax, curMin := 0, 0, 0
	prevBits := -1
	for _, x := range nums {
		b := bitsCount(x)
		if b != prevBits {
			if prevBits != -1 {
				prevMax = curMax
			}
			curMax, curMin = x, x
			prevBits = b
		} else {
			if x < curMin {
				curMin = x
			}
			if x > curMax {
				curMax = x
			}
		}
		if curMin < prevMax {
			return false
		}
	}
	return true
}

func bitsCount(x int) int {
	cnt := 0
	for x > 0 {
		cnt += x & 1
		x >>= 1
	}
	return cnt
}
```
