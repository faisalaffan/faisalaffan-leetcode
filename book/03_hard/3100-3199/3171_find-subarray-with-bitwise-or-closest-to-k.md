# 3171 — Find Subarray With Bitwise Or Closest To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumDifference(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3171: Find Subarray With Bitwise OR Closest to K
// https://leetcode.com/problems/find-subarray-with-bitwise-or-closest-to-k/
// Difficulty: Hard
//
// Find the minimum absolute difference between the bitwise OR of any subarray
// and k. Since OR only sets bits (never clears), for each ending position there
// are at most O(log MAX) distinct OR values.
//
// Approach: maintain a set of distinct OR values for subarrays ending at each
// position; track the minimum |OR - k|.

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumDifference(nums []int, k int) int {
	ans := abs(nums[0] - k)
	// cur holds distinct OR values of subarrays ending at the current position.
  // HashMap: O(1) lookup
	cur := make(map[int]bool)
	cur[nums[0]] = true

	for _, x := range nums[1:] {
  // HashMap: O(1) lookup
		nxt := make(map[int]bool)
		nxt[x] = true
		if abs(x-k) < ans {
			ans = abs(x - k)
		}
		for v := range cur {
			orVal := v | x
			nxt[orVal] = true
			if abs(orVal-k) < ans {
				ans = abs(orVal - k)
			}
		}
		cur = nxt
	}
	return ans
}

func main() {
	fmt.Println(minimumDifference([]int{1, 2, 4}, 5)) // expect 0 (1|4 = 5)
}
```
