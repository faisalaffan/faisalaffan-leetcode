# 3960 — Frequency Balance Subarray

## Deskripsi

**Soal:** [3960. Frequency Balance Subarray](https://leetcode.com/problems/frequency-balance-subarray/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N^2)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

**Fungsi Solusi:** `func FrequencyBalanceSubarray(nums []int) int`

> **Ide Kunci:** Enumerate all subarrays. Track element frequencies and

## Solusi Go

```go
package main

// LeetCode #3960: Frequency Balance Subarray
// https://leetcode.com/problems/frequency-balance-subarray/
// Difficulty: Medium
// Time: O(N^2) | Space: O(N)
// Approach: Enumerate all subarrays. Track element frequencies and
// frequency-of-frequencies. Valid if: 1 distinct value, or exactly 2
// distinct freq values where one is double the other.

import "fmt"

func FrequencyBalanceSubarray(nums []int) int {
	n := len(nums)
	ans := 0

	for l := 0; l < n; l++ {
  // Membuat map untuk pencarian O(1): key → value
		cnt := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
		freq := make(map[int]int) // frequency-of-frequencies

		for r := l; r < n; r++ {
			val := nums[r]
			oldF := cnt[val]
			if oldF > 0 {
				freq[oldF]--
				if freq[oldF] == 0 {
					delete(freq, oldF)
				}
			}
			newF := oldF + 1
			cnt[val] = newF
			freq[newF]++

			// Check validity
			valid := false
			if len(cnt) == 1 {
				valid = true
			} else if len(freq) == 2 {
				// Exactly two distinct frequency values
  // Membuat slice untuk menyimpan hasil
				vals := make([]int, 0, 2)
				for f := range freq {
					vals = append(vals, f)
				}
				a, b := vals[0], vals[1]
				if a > b {
					a, b = b, a
				}
				// One must be double the other
				if b == 2*a {
					valid = true
				}
			}

			if valid && (r-l+1) > ans {
				ans = r - l + 1
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(FrequencyBalanceSubarray([]int{1, 2, 2, 1, 2, 3, 3, 3})) // Expected: 5

	// Example 2
	fmt.Println(FrequencyBalanceSubarray([]int{5, 5, 5, 5})) // Expected: 4

	// Example 3
	fmt.Println(FrequencyBalanceSubarray([]int{1, 2, 3, 4})) // Expected: 1
}
```
