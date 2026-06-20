# 3284 — Sum Of Consecutive Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func getSum(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3284: Sum of Consecutive Subarrays
// https://leetcode.com/problems/sum-of-consecutive-subarrays/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(getSum([]int{1, 2, 3}))    // 20
	fmt.Println(getSum([]int{1, 2, 3, 5})) // 25
	fmt.Println(getSum([]int{1, 2, 3, 4})) // 50
}

func getSum(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	f, g := 1, 1
	s, t := nums[0], nums[0]
	ans := nums[0]

	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]

		if diff == 1 {
			f++
			s += f * nums[i]
			ans = (ans + s) % mod
		} else {
			f = 1
			s = nums[i]
		}

		if diff == -1 {
			g++
			t += g * nums[i]
			ans = (ans + t) % mod
		} else {
			g = 1
			t = nums[i]
		}

		if diff != 1 && diff != -1 {
			ans = (ans + nums[i]) % mod
		}
	}

	return ans
}
```
