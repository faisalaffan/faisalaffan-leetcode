# 0216 — Combination Sum Iii

## Deskripsi

**Soal:** [0216. Combination Sum Iii](https://leetcode.com/problems/combination-sum-iii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(C(9,k)), Space: O(k)  
**Kompleksitas Ruang:** O(k)

**Algoritma:** —

**Fungsi Solusi:** `func combinationSum3(k int, n int) [][]int`

## Solusi Go

```go
package main

// LeetCode #216: Combination Sum III
// https://leetcode.com/problems/combination-sum-iii/
// Difficulty: Medium
// Time: O(C(9,k)), Space: O(k)

import "fmt"

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	var backtrack func(start, remaining int, combo []int)
	backtrack = func(start, remaining int, combo []int) {
		if len(combo) == k && remaining == 0 {
  // Membuat slice untuk menyimpan hasil
			comboCopy := make([]int, len(combo))
			copy(comboCopy, combo)
			result = append(result, comboCopy)
			return
		}
		if len(combo) > k || remaining < 0 {
			return
		}

		for i := start; i <= 9; i++ {
			combo = append(combo, i)
			backtrack(i+1, remaining-i, combo)
			combo = combo[:len(combo)-1]
		}
	}

	backtrack(1, n, []int{})
	return result
}

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
}
```
