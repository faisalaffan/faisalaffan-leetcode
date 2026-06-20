# 3109 — Find The Index Of Permutation

## Deskripsi

**Soal:** [3109. Find The Index Of Permutation](https://leetcode.com/problems/find-the-index-of-permutation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func getPermutationIndex(perm []int) int`

## Solusi Go

```go
package main

// LeetCode #3109: Find the Index of Permutation
// https://leetcode.com/problems/find-the-index-of-permutation/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n)

import "fmt"

func getPermutationIndex(perm []int) int {
	n := len(perm)
	mod := int64(1000000007)

  // Membuat slice untuk menyimpan hasil
	fact := make([]int64, n)
	fact[0] = 1
	for i := 1; i < n; i++ {
		fact[i] = fact[i-1] * int64(i) % mod
	}

  // Membuat slice untuk menyimpan hasil
	bit := make([]int, n+1)

	update := func(idx, val int) {
		for idx <= n {
			bit[idx] += val
			idx += idx & -idx
		}
	}

	query := func(idx int) int {
		sum := 0
		for idx > 0 {
			sum += bit[idx]
			idx -= idx & -idx
		}
		return sum
	}

	for i := 1; i <= n; i++ {
		update(i, 1)
	}

	ans := int64(0)
	for i := 0; i < n; i++ {
		smaller := query(perm[i]) - 1
		ans = (ans + int64(smaller)*fact[n-1-i]) % mod
		update(perm[i], -1)
	}

	return int((ans + 1) % mod)
}

func main() {
	fmt.Println(getPermutationIndex([]int{1, 2, 3})) // Expected: 1
	fmt.Println(getPermutationIndex([]int{3, 2, 1})) // Expected: 6
	fmt.Println(getPermutationIndex([]int{2, 1, 3})) // Expected: 3
}
```
