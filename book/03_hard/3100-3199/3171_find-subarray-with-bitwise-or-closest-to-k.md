# 3171 — Find Subarray With Bitwise Or Closest To K

## Deskripsi

**Soal:** [3171. Find Subarray With Bitwise Or Closest To K](https://leetcode.com/problems/find-subarray-with-bitwise-or-closest-to-k/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func abs(x int) int`

> **Ide Kunci:** maintain a set of distinct OR values for subarrays ending at each

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	cur := make(map[int]bool)
	cur[nums[0]] = true

	for _, x := range nums[1:] {
  // Membuat map untuk pencarian O(1): key → value
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
