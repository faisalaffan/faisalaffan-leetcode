# 3347 — Maximum Frequency Of An Element After Performing Operations Ii

## Deskripsi

**Soal:** [3347. Maximum Frequency Of An Element After Performing Operations Ii](https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser), Prefix Sum (jumlah kumulatif)

> **Ide Kunci:** For each unique value, consider it as the final value.

## Solusi Go

```go
package main

// LeetCode #3347: Maximum Frequency of an Element After Performing Operations II
// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-ii/
// Difficulty: Hard
//
// In one operation, add k to any element in nums. Perform at most
// numOperations operations. Maximize the frequency of any single
// value in the resulting array.
//
// Approach: For each unique value, consider it as the final value.
// Count how many existing elements can reach it within allowed
// operations. Use prefix sums over sorted values.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maxFrequency([]int{1, 2, 4}, 2, 2))
	// Example 2
	fmt.Println(maxFrequency([]int{5, 5, 5, 10}, 5, 1))
	// Edge: single element
	fmt.Println(maxFrequency([]int{7}, 3, 0))
}

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	n := len(nums)

	// Count frequency of each value
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Collect unique values
  // Membuat slice untuk menyimpan hasil
	unique := make([]int, 0, len(freq))
	for v := range freq {
		unique = append(unique, v)
	}
	sort.Ints(unique)

	// Sliding window: count elements in range [val - k, val + k]
	// But only numOperations of them can be changed to val
	ans := 0
	left := 0
	for right := 0; right < n; right++ {
		// Shrink window to [target - k, target + k]
		for nums[right]-nums[left] > 2*k {
			left++
		}
		total := right - left + 1
		originalCnt := freq[nums[right]]
		// We can change at most numOperations elements to this value
		possible := originalCnt + min(numOperations, total-originalCnt)
		if possible > ans {
			ans = possible
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
