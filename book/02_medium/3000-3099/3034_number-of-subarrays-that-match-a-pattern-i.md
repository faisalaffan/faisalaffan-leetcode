# 3034 — Number Of Subarrays That Match A Pattern I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countMatchingSubarrays(nums []int, pattern []int) (ans int)`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n*m)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3034: Number of Subarrays That Match a Pattern I
// https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-i/
// Difficulty: Medium
// Time: O(n*m) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countMatchingSubarrays([]int{1, 2, 3, 4, 5, 6}, []int{1, 1}))
	fmt.Println(countMatchingSubarrays([]int{1, 4, 4, 1, 3, 5, 5, 3}, []int{1, 0, -1}))
}

func countMatchingSubarrays(nums []int, pattern []int) (ans int) {
	n, m := len(nums), len(pattern)
outer:
	for i := 0; i+m < n; i++ {
		for k := 0; k < m; k++ {
			diff := 0
			if nums[i+k+1] > nums[i+k] {
				diff = 1
			} else if nums[i+k+1] < nums[i+k] {
				diff = -1
			}
			if diff != pattern[k] {
				continue outer
			}
		}
		ans++
	}
	return
}
```
