# 3470 — Permutations Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func permutationsIV(n, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3470: Permutations IV
// https://leetcode.com/problems/permutations-iv/
// Difficulty: Hard
//
// Count permutations of [1..n] with exactly k inversions.
// Standard Mahonian DP: dp[i][j] = sum_{t=0}^{i-1} dp[i-1][j-t]
// where insertion of i at position t adds (i-1-t) inversions.
//
// Optimised with prefix sums for O(n*k) time.

import "fmt"

const MOD = 1_000_000_007

func permutationsIV(n, k int) int {
  // Matriks 2D
	dp := make([][]int, n+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
  // Alokasi slice
		pref := make([]int, k+2)
		for j := 0; j <= k; j++ {
			pref[j+1] = (pref[j] + dp[i-1][j]) % MOD
		}
		for j := 0; j <= k; j++ {
			// dp[i][j] = sum_{t=0}^{i-1} dp[i-1][j-t], where j-t >= 0
			// t = 0 to min(i-1, j)
			// = sum of dp[i-1][j-min(i-1,j) .. j]
			left := j - (i - 1)
			if left < 0 {
				left = 0
			}
			sum := pref[j+1] - pref[left]
			if sum < 0 {
				sum += MOD
			}
			dp[i][j] = sum
		}
	}

	return dp[n][k]
}

func main() {
	fmt.Printf("n=3,k=1 -> %d (expected 2: [1,3,2], [2,1,3])\n", permutationsIV(3, 1))
	fmt.Printf("n=3,k=0 -> %d (expected 1: [1,2,3])\n", permutationsIV(3, 0))
	fmt.Printf("n=3,k=2 -> %d (expected 2: [2,3,1], [3,1,2])\n", permutationsIV(3, 2))
	fmt.Printf("n=4,k=3 -> %d\n", permutationsIV(4, 3))
	fmt.Printf("n=2,k=1 -> %d (expected 1: [2,1])\n", permutationsIV(2, 1))
	fmt.Printf("n=4,k=1 -> %d\n", permutationsIV(4, 1))
	fmt.Printf("n=5,k=4 -> %d\n", permutationsIV(5, 4))
}
```
