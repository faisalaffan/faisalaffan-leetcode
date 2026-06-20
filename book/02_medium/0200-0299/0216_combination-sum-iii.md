# 0216 — Combination Sum Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func combinationSum3(k int, n int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** O(C(9,k)), Space: O(k)  |  **Ruang:** O(k)

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

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
  // Alokasi slice
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
