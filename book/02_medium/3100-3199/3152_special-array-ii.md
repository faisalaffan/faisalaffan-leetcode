# 3152 — Special Array Ii

## Deskripsi

**Soal:** [3152. Special Array Ii](https://leetcode.com/problems/special-array-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func isArraySpecial(nums []int, queries [][]int) []bool`

## Solusi Go

```go
package main

// LeetCode #3152: Special Array II
// https://leetcode.com/problems/special-array-ii/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, n)
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1]
		if nums[i]%2 == nums[i-1]%2 {
			prefix[i]++
		}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]bool, len(queries))
	for i, q := range queries {
		from, to := q[0], q[1]
		ans[i] = prefix[from] == prefix[to]
	}
	return ans
}

func main() {
	fmt.Println(isArraySpecial([]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}))        // Expected: [false]
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}))   // Expected: [false, true]
}
```
