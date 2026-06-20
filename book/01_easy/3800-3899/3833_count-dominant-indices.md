# 3833 — Count Dominant Indices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountDominantIndices(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3833: Count Dominant Indices
// https://leetcode.com/problems/count-dominant-indices/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountDominantIndices([]int{5, 4, 3}))
	fmt.Println(CountDominantIndices([]int{4, 1, 2}))
	fmt.Println(CountDominantIndices([]int{1}))
}

// Time: O(n)
// Space: O(n)
func CountDominantIndices(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
  // Alokasi slice
	suffixSum := make([]int, n)
	suffixSum[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + nums[i]
	}
	count := 0
	for i := 0; i < n-1; i++ {
		rightSum := suffixSum[i+1]
		rightCount := n - 1 - i
		if nums[i]*rightCount > rightSum {
			count++
		}
	}
	return count
}
```
