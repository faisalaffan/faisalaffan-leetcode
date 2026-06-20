# 0894 — All Possible Full Binary Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func AllPossibleFullBinaryTrees(n int) []*TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(2^n)  |  **Ruang:** O(2^n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #894: All Possible Full Binary Trees
// https://leetcode.com/problems/all-possible-full-binary-trees/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(len(AllPossibleFullBinaryTrees(7)))
	fmt.Println(len(AllPossibleFullBinaryTrees(3)))
	fmt.Println(len(AllPossibleFullBinaryTrees(1)))
}

// Time: O(2^n) | Space: O(2^n)
func AllPossibleFullBinaryTrees(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}

  // HashMap: O(1) lookup
	memo := make(map[int][]*TreeNode)

	var dfs func(int) []*TreeNode
	dfs = func(count int) []*TreeNode {
		if trees, ok := memo[count]; ok {
			return trees
		}

		if count == 1 {
			return []*TreeNode{{Val: 0}}
		}

		var res []*TreeNode
		for left := 1; left < count; left += 2 {
			right := count - 1 - left
			for _, l := range dfs(left) {
				for _, r := range dfs(right) {
					res = append(res, &TreeNode{Val: 0, Left: l, Right: r})
				}
			}
		}

		memo[count] = res
		return res
	}

	return dfs(n)
}
```
