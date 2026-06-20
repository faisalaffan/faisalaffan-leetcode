# 2176 — Count Equal And Divisible Pairs In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountEqualAndDivisiblePairsInAnArray(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2176: Count Equal and Divisible Pairs in an Array
// https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountEqualAndDivisiblePairsInAnArray([]int{3, 1, 2, 2, 2, 1, 3}, 2)) // 4
	fmt.Println(CountEqualAndDivisiblePairsInAnArray([]int{1, 2, 3, 4}, 1))           // 0
}

// Time: O(n^2), Space: O(1)
func CountEqualAndDivisiblePairsInAnArray(nums []int, k int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				count++
			}
		}
	}
	return count
}
```
