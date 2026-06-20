# 2962 — Count Subarrays Where Max Element Appears At Least K Times

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countSubarrays2962(nums []int, k int) (ans int64)`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2962: Count Subarrays Where Max Element Appears at Least K Times
// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countSubarrays2962([]int{1, 3, 2, 3, 3}, 2))
	fmt.Println(countSubarrays2962([]int{1, 4, 2, 1}, 3))
}

func countSubarrays2962(nums []int, k int) (ans int64) {
	mx := 0
	for _, x := range nums {
		if x > mx {
			mx = x
		}
	}
	n := len(nums)
	cnt, j := 0, 0
	for _, x := range nums {
		for ; j < n && cnt < k; j++ {
			if nums[j] == mx {
				cnt++
			}
		}
		if cnt < k {
			break
		}
		ans += int64(n - j + 1)
		if x == mx {
			cnt--
		}
	}
	return
}
```
