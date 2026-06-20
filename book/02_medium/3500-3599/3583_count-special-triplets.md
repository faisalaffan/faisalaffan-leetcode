# 3583 — Count Special Triplets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountSpecialTriplets(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3583: Count Special Triplets
// https://leetcode.com/problems/count-special-triplets/
// Difficulty: Medium
// Complexity: O(n^3) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", CountSpecialTriplets([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", CountSpecialTriplets([]int{1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", CountSpecialTriplets([]int{1, 2, 3}))
}

func CountSpecialTriplets(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j] == nums[k] || nums[i]+nums[k] == nums[j] || nums[j]+nums[k] == nums[i] {
					count++
				}
			}
		}
	}
	return count
}
```
