# 3939 — Count Non Adjacent Subsets In A Rooted Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countNonAdjacentSubsets(parent []int, nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3939: Count Non Adjacent Subsets in a Rooted Tree
// https://leetcode.com/problems/count-non-adjacent-subsets-in-a-rooted-tree/
// Difficulty: Hard
//
// Count subsets of nodes in a rooted tree where no two selected
// nodes are adjacent (no parent-child pair both selected).
// Each subset must have size exactly k.
//
// Approach: Tree DP. dp[u][j][0/1] = count of subsets of size j
// in subtree of u where u is not selected / selected.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countNonAdjacentSubsets([]int{-1, 0, 0, 1, 1}, []int{1, 2, 3, 4, 5}, 2))
	// Example 2
	fmt.Println(countNonAdjacentSubsets([]int{-1, 0, 0}, []int{1, 1, 1}, 1))
	// Edge: k = 0
	fmt.Println(countNonAdjacentSubsets([]int{-1, 0, 1}, []int{1, 2, 3}, 0))
}

const TREE_MOD = 1000000007

func countNonAdjacentSubsets(parent []int, nums []int, k int) int {
	n := len(parent)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

  // Matriks 2D
	children := make([][]int, n)
	root := -1
	for i, p := range parent {
		if p == -1 {
			root = i
		} else {
			children[p] = append(children[p], i)
		}
	}

	var dfs func(u int) ([]int, []int)
	dfs = func(u int) ([]int, []int) {
  // Alokasi slice
		dp0 := make([]int, k+1)
  // Alokasi slice
		dp1 := make([]int, k+1)
		dp0[0] = 1

		for _, v := range children[u] {
			cdp0, cdp1 := dfs(v)
  // Alokasi slice
			ndp0 := make([]int, k+1)
  // Alokasi slice
			ndp1 := make([]int, k+1)

			for j := 0; j <= k; j++ {
				if dp0[j] == 0 {
					continue
				}
				for cj := 0; cj <= k-j; cj++ {
					if cdp0[cj] == 0 {
						continue
					}
					ndp0[j+cj] = (ndp0[j+cj] + dp0[j]*cdp0[cj]%TREE_MOD) % TREE_MOD
				}
				for cj := 0; cj <= k-j; cj++ {
					if cdp1[cj] == 0 {
						continue
					}
					ndp0[j+cj] = (ndp0[j+cj] + dp0[j]*cdp1[cj]%TREE_MOD) % TREE_MOD
				}
			}

			for j := 0; j <= k; j++ {
				if dp1[j] == 0 {
					continue
				}
				for cj := 0; cj <= k-j; cj++ {
					if cdp0[cj] == 0 {
						continue
					}
					ndp1[j+cj] = (ndp1[j+cj] + dp1[j]*cdp0[cj]%TREE_MOD) % TREE_MOD
				}
			}

			dp0 = ndp0
			dp1 = ndp1
		}

		dp1[1] = (dp1[1] + 1) % TREE_MOD

		return dp0, dp1
	}

	dp0, dp1 := dfs(root)
	return (dp0[k] + dp1[k]) % TREE_MOD
}
```
