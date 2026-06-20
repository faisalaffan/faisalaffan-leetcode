# 0548 — Split Array With Equal Sum

## Deskripsi

**Soal:** [0548. Split Array With Equal Sum](https://leetcode.com/problems/split-array-with-equal-sum/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #548: Split Array with Equal Sum
// https://leetcode.com/problems/split-array-with-equal-sum/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(splitArray([]int{1, 2, 1, 2, 1, 2, 1})) // Expected: true
}

func splitArray(nums []int) bool {
	n := len(nums)
	if n < 7 {
		return false
	}

  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	// Try position j (second split, 0-indexed)
	for j := 3; j <= n-4; j++ {
  // Membuat map untuk pencarian O(1): key → value
		set := make(map[int]bool)
		// Try position i (first split)
		for i := 1; i < j-1; i++ {
			sum1 := prefix[i-1]
			sum2 := prefix[j-1] - prefix[i]
			if sum1 == sum2 {
				set[sum1] = true
			}
		}
		// Try position k (third split)
		for k := j + 2; k < n-1; k++ {
			sum3 := prefix[k-1] - prefix[j]
			sum4 := prefix[n-1] - prefix[k]
			if sum3 == sum4 && set[sum3] {
				return true
			}
		}
	}
	return false
}
```
