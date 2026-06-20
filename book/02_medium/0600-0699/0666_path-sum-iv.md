# 0666 — Path Sum Iv

## Deskripsi

**Soal:** [0666. Path Sum Iv](https://leetcode.com/problems/path-sum-iv/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

## Solusi Go

```go
package main

// LeetCode #666: Path Sum IV
// https://leetcode.com/problems/path-sum-iv/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(pathSumIV([]int{113, 215, 221}))
	fmt.Println(pathSumIV([]int{113, 221}))
}

func pathSumIV(nums []int) int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return 0
	}

  // Membuat map untuk pencarian O(1): key → value
	tree := make(map[int]int)
	for _, num := range nums {
		tree[num/10] = num % 10
	}

	total := 0
	var dfs func(key int, sum int)
	dfs = func(key int, sum int) {
		depth := key / 10
		pos := key % 10
		leftKey := (depth+1)*10 + pos*2 - 1
		rightKey := (depth+1)*10 + pos*2

		curSum := sum + tree[key]

		_, hasLeft := tree[leftKey]
		_, hasRight := tree[rightKey]

		if !hasLeft && !hasRight {
			total += curSum
			return
		}

		if hasLeft {
			dfs(leftKey, curSum)
		}
		if hasRight {
			dfs(rightKey, curSum)
		}
	}

	dfs(nums[0]/10, 0)
	return total
}
```
