# 2772 — Apply Operations To Make All Array Elements Equal To Zero

## Deskripsi

**Soal:** [2772. Apply Operations To Make All Array Elements Equal To Zero](https://leetcode.com/problems/apply-operations-to-make-all-array-elements-equal-to-zero/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func ApplyOperationsToMakeAllArrayElementsEqualToZero(nums []int, k int) bool`

## Solusi Go

```go
package main

// LeetCode #2772: Apply Operations to Make All Array Elements Equal to Zero
// https://leetcode.com/problems/apply-operations-to-make-all-array-elements-equal-to-zero/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func ApplyOperationsToMakeAllArrayElementsEqualToZero(nums []int, k int) bool {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	diff := make([]int, n+1)
	cur := 0

	for i := 0; i < n; i++ {
		cur += diff[i]
		val := nums[i] + cur
		if val < 0 {
			return false
		}
		val %= 2
		if val != 0 {
			if i+k > n {
				return false
			}
			diff[i] -= 1
			diff[i+k] += 1
			cur -= 1
		}
	}

	return true
}

func main() {
	fmt.Println(ApplyOperationsToMakeAllArrayElementsEqualToZero([]int{2, 0, 2}, 2))
	fmt.Println(ApplyOperationsToMakeAllArrayElementsEqualToZero([]int{1, 0, 1}, 2))
}
```
