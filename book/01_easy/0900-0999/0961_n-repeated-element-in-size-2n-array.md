# 0961 — N Repeated Element In Size 2N Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func repeatedNTimes(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #961: N-Repeated Element in Size 2N Array
// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3})) // 3
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2})) // 2
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4})) // 5
}

// repeatedNTimes finds the element repeated n times in a 2n size array.
// Time: O(n). Space: O(1).
func repeatedNTimes(nums []int) int {
	// Since the element appears n times in 2n, any two consecutive elements
	// must contain the repeated element (in most cases).
  // Linear scan O(n)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] || nums[i] == nums[i+2] {
			return nums[i]
		}
	}
	// If not found yet, the repeated element is in the last 3 positions
	return nums[len(nums)-1]
}
```
