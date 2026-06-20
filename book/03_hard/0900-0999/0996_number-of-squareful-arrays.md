# 0996 — Number Of Squareful Arrays

## Deskripsi

**Soal:** [0996. Number Of Squareful Arrays](https://leetcode.com/problems/number-of-squareful-arrays/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), Backtracking (pelacakan mundur), Bitmask (representasi himpunan dengan bit)

> **Ide Kunci:** DFS + backtracking + bitmask.

## Solusi Go

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
	sort.Ints(nums)
	n := len(nums)
	mask := 1<<n - 1
  // Membuat map untuk pencarian O(1): key → value
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
  // Loop standar: indeks 0 sampai n-1
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
