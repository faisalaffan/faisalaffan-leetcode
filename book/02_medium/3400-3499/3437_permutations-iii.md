# 3437 — Permutations Iii

## Deskripsi

**Soal:** [3437. Permutations Iii](https://leetcode.com/problems/permutations-iii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n!) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func permute(n int) [][]int`

## Solusi Go

```go
package main

// LeetCode #3437: Permutations III
// https://leetcode.com/problems/permutations-iii/
// Difficulty: Medium [Paid]
// Time: O(n!) Space: O(n)

import "fmt"

func permute(n int) [][]int {
	var ans [][]int
  // Membuat slice untuk menyimpan hasil
	used := make([]bool, n+1)
  // Membuat slice untuk menyimpan hasil
	cur := make([]int, 0, n)

	var dfs func()
	dfs = func() {
		if len(cur) == n {
  // Membuat slice untuk menyimpan hasil
			tmp := make([]int, n)
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}
		start := len(cur)%2 + 1
		for i := start; i <= n; i += 2 {
			if !used[i] {
				used[i] = true
				cur = append(cur, i)
				dfs()
				cur = cur[:len(cur)-1]
				used[i] = false
			}
		}
	}
	dfs()
	return ans
}

func main() {
	fmt.Println(len(permute(3))) // 2
	fmt.Println(len(permute(4))) // 4
	for _, p := range permute(3) {
		fmt.Println(p)
	}
}
```
