# 3737 — Count Subarrays With Majority Element I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countSubarraysWithMajorityElementI(nums []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3737: Count Subarrays With Majority Element I
// https://leetcode.com/problems/count-subarrays-with-majority-element-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func countSubarraysWithMajorityElementI(nums []int, target int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := i; j < n; j++ {
			if nums[j] == target {
				cnt++
			}
			if cnt*2 > j-i+1 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countSubarraysWithMajorityElementI([]int{1, 2, 1, 2, 1}, 1))
	fmt.Println(countSubarraysWithMajorityElementI([]int{3, 1, 2, 3}, 3))
	fmt.Println(countSubarraysWithMajorityElementI([]int{1, 2, 3, 4}, 1))
}
```
