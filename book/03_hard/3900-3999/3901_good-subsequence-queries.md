# 3901 — Good Subsequence Queries

## Deskripsi

**Soal:** [3901. Good Subsequence Queries](https://leetcode.com/problems/good-subsequence-queries/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** DP. For each position i, track count of subsequences

## Solusi Go

```go
package main

// LeetCode #3901: Good Subsequence Queries
// https://leetcode.com/problems/good-subsequence-queries/
// Difficulty: Hard
//
// A non-empty subsequence of nums is "good" if its length is
// strictly less than n and sum modulo p is 0. Process queries
// where each query asks for the count of good subsequences
// where the last element is at a specific index.
//
// Approach: DP. For each position i, track count of subsequences
// ending at i with each modulo value. Use prefix counts to
// answer queries.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countGoodSubseq([]int{1, 2, 3}, 2, [][]int{{0, 2}}))
	// Example 2
	fmt.Println(countGoodSubseq([]int{1, 1, 1}, 1, [][]int{{0, 1}, {1, 2}}))
	// Edge: single element
	fmt.Println(countGoodSubseq([]int{5}, 3, [][]int{{0, 0}}))
}

const mod = 1000000007

func countGoodSubseq(nums []int, p int, queries [][]int) int {
	ans := 0
	for _, q := range queries {
		l, r := q[0], q[1]
		subarrayLen := r - l + 1
		if subarrayLen < 2 {
			ans = 0
			continue
		}

		// DP on subarray nums[l:r+1]
  // Membuat slice untuk menyimpan hasil
		local := make([]int, p)
		local[0] = 1
		for i := l; i <= r; i++ {
			val := nums[i] % p
  // Membuat slice untuk menyimpan hasil
			ndp := make([]int, p)
			copy(ndp, local)
			for m := 0; m < p; m++ {
				nm := (m + val) % p
				ndp[nm] = (ndp[nm] + local[m]) % mod
			}
			local = ndp
		}
		// Subtract empty subsequence
		ans = (ans + local[0] - 1 + mod) % mod
	}

	return ans
}
```
