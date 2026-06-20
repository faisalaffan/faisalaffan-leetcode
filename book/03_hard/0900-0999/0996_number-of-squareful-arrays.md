# 0996 — Number Of Squareful Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numSquarefulPerms(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DFS, DP, Backtracking, Sorting, Bitmask

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #996: Number of Squareful Arrays
// https://leetcode.com/problems/number-of-squareful-arrays/
// Difficulty: Hard
//
// Approach: DFS + backtracking + bitmask.
//   A squareful array is one where every adjacent pair sums to a perfect square.
//   We sort the input and use a visited bitmask to count unique permutations
//   where each adjacent pair is squareful.

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(numSquarefulPerms([]int{1, 17, 8}))       // 2
	fmt.Println(numSquarefulPerms([]int{2, 2, 2}))        // 1
}

func numSquarefulPerms(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	n := len(nums)
	mask := 1<<n - 1
  // HashMap: O(1) lookup
	memo := make(map[int]int)
	return dfs(nums, 0, mask, -1, memo)
}

func dfs(nums []int, used int, all int, last int, memo map[int]int) int {
	if used == all {
		return 1
	}
	key := (used << 5) | (last + 1)
	if val, ok := memo[key]; ok {
		return val
	}
	total := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		if used&(1<<i) != 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && used&(1<<(i-1)) == 0 {
			continue
		}
		if last == -1 || isPerfectSquare(nums[last]+nums[i]) {
			total += dfs(nums, used|(1<<i), all, i, memo)
		}
	}
	memo[key] = total
	return total
}

func isPerfectSquare(n int) bool {
	r := int(math.Sqrt(float64(n)))
	return r*r == n
}
```
