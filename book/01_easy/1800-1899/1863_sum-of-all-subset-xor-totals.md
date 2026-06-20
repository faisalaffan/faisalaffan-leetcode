# 1863 — Sum Of All Subset Xor Totals

## Deskripsi

**Soal:** [1863. Sum Of All Subset Xor Totals](https://leetcode.com/problems/sum-of-all-subset-xor-totals/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(2^n), Space: O(n) (recursion stack)  
**Kompleksitas Ruang:** O(n) (recursion stack)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), Stack (tumpukan LIFO)

**Fungsi Solusi:** `func SubsetXORSum(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #1863: Sum of All Subset XOR Totals
// https://leetcode.com/problems/sum-of-all-subset-xor-totals/
// Difficulty: Easy

import "fmt"

// Time: O(2^n), Space: O(n) (recursion stack)
func SubsetXORSum(nums []int) int {
	return dfs(nums, 0, 0)
}

func dfs(nums []int, idx int, currentXor int) int {
	if idx == len(nums) {
		return currentXor
	}
	// Include nums[idx] or skip it
	return dfs(nums, idx+1, currentXor^nums[idx]) + dfs(nums, idx+1, currentXor)
}

func main() {
	fmt.Println(SubsetXORSum([]int{1, 3}))
	fmt.Println(SubsetXORSum([]int{5, 1, 6}))
	fmt.Println(SubsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}
```
