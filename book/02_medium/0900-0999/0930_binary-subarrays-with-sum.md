# 0930 — Binary Subarrays With Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func numSubarraysWithSum(nums []int, goal int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #930: Binary Subarrays With Sum
// https://leetcode.com/problems/binary-subarrays-with-sum/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(1)
func numSubarraysWithSum(nums []int, goal int) int {
	// sliding window for sum <= goal, then subtract sum < goal
	atMost := func(g int) int {
		if g < 0 {
			return 0
		}
		sum, cnt, left := 0, 0, 0
		for right, v := range nums {
			sum += v
			for sum > g {
				sum -= nums[left]
				left++
			}
			cnt += right - left + 1
		}
		return cnt
	}
	return atMost(goal) - atMost(goal-1)
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}
```
