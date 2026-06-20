# 2216 — Minimum Deletions To Make Array Beautiful

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minDeletion(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2216: Minimum Deletions to Make Array Beautiful
// https://leetcode.com/problems/minimum-deletions-to-make-array-beautiful/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minDeletion(nums []int) int {
	n := len(nums)
	deletions := 0
	i := 0

	for i < n-1 {
		idx := i - deletions
		if idx%2 == 0 && nums[i] == nums[i+1] {
			deletions++
			i++
		} else {
			i++
		}
	}

	// If after deletions the array length is odd, delete last element
	if (n-deletions)%2 == 1 {
		deletions++
	}
	return deletions
}

func main() {
	// Test case 1
	fmt.Println(minDeletion([]int{1, 1, 2, 3, 5}))
	// Expected: 1

	// Test case 2
	fmt.Println(minDeletion([]int{1, 1, 2, 2, 3, 3}))
	// Expected: 2
}
```
