# 3859 — Count Subarrays With K Distinct Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countSubarrays(nums []int, k int, m int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, Sliding Window

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3859: Count Subarrays With K Distinct Integers
// https://leetcode.com/problems/count-subarrays-with-k-distinct-integers/
// Difficulty: Hard
//
// Count subarrays that contain exactly k distinct integers, where
// the maximum element in the subarray appears at least m times.
//
// Approach: Sliding window with frequency map. Track distinct count
// and max frequency. Expand right, shrink left when conditions
// violated.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countSubarrays([]int{1, 2, 1, 2, 3}, 2, 1))
	// Example 2
	fmt.Println(countSubarrays([]int{1, 1, 2, 2, 3}, 2, 2))
	// Edge: single element
	fmt.Println(countSubarrays([]int{5}, 1, 1))
	// Edge: k = 1, m = 2
	fmt.Println(countSubarrays([]int{1, 1, 1, 2}, 1, 2))
}

func countSubarrays(nums []int, k int, m int) int64 {
	n := len(nums)
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	var ans int64
	left := 0
	distinct := 0
	maxFreq := 0

	for right := 0; right < n; right++ {
		val := nums[right]
		freq[val]++
		if freq[val] == 1 {
			distinct++
		}
		if freq[val] > maxFreq {
			maxFreq = freq[val]
		}

		for distinct > k || (distinct == k && maxFreq < m) {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				distinct--
			}
			left++
			maxFreq = 0
			for _, c := range freq {
				if c > maxFreq {
					maxFreq = c
				}
			}
		}

		if distinct == k && maxFreq >= m {
			ans += int64(right - left + 1)
		}
	}

	return ans
}
```
