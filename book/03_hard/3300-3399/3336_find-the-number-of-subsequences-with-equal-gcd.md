# 3336 — Find The Number Of Subsequences With Equal Gcd

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func subsequencePairCount(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3336: Find the Number of Subsequences With Equal GCD
// https://leetcode.com/problems/find-the-number-of-subsequences-with-equal-gcd/
// Difficulty: Hard
//
// Use inclusion-exclusion with MObius-like counting:
// 1. For each g, count subsequence pairs where both have GCD that is a multiple of g.
// 2. Use MObius inversion to get exact GCD counts.
// 3. Sum over g where both subsequences have exact GCD = g.

import "fmt"

func main() {
	// Example: [1,2,3,4] -> 2
	fmt.Println(subsequencePairCount([]int{1, 2, 3, 4}))

	// All same: [1,1,1,1] -> 50
	fmt.Println(subsequencePairCount([]int{1, 1, 1, 1}))

	// [2,4,8] -> 3
	fmt.Println(subsequencePairCount([]int{2, 4, 8}))

	// [5,5] -> 1
	fmt.Println(subsequencePairCount([]int{5, 5}))

	// [1,1,1] -> 6
	fmt.Println(subsequencePairCount([]int{1, 1, 1}))
}

const MOD = 1000000007

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func subsequencePairCount(nums []int) int {
	n := len(nums)
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	// Count frequency of each value
  // Alokasi slice
	freq := make([]int, maxVal+1)
	for _, v := range nums {
		freq[v]++
	}

	// Precompute combination nCk for n up to n, k up to 5
  // Matriks 2D
	C := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]int, 6)
		C[i][0] = 1
		for j := 1; j <= i && j <= 5; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % MOD
		}
	}

	// cntMult[g] = number of elements divisible by g
  // Alokasi slice
	cntMult := make([]int, maxVal+1)
	for g := 1; g <= maxVal; g++ {
		for m := g; m <= maxVal; m += g {
			cntMult[g] += freq[m]
		}
	}

	// f[g] = number of ways to pick 2 non-empty disjoint subsequences
	// where each element in each subsequence is divisible by g
	// (i.e., both GCDs are multiples of g)
  // Alokasi slice
	f := make([]int, maxVal+1)
	for g := 1; g <= maxVal; g++ {
		c := cntMult[g]
		// Total ways with 2 non-empty subsequences from c elements:
		// For each element, 3 choices: skip, seq1, seq2
		// Total = 3^c
		// Subtract: seq1 empty => 2^c, seq2 empty => 2^c
		// Add back: both empty => 1
		total := powMod(3, c)
		empty1 := powMod(2, c)
		empty2 := empty1
		f[g] = (total - empty1 - empty2 + 1) % MOD
		if f[g] < 0 {
			f[g] += MOD
		}
	}

	// Use MObius-like inclusion-exclusion to get exact GCD = g
	// gExact[g] = exact pairs with GCD = g
  // Alokasi slice
	gExact := make([]int, maxVal+1)
	for g := maxVal; g >= 1; g-- {
		gExact[g] = f[g]
		for m := 2 * g; m <= maxVal; m += g {
			gExact[g] = (gExact[g] - gExact[m] + MOD) % MOD
		}
	}

	// Sum over g where both subsequences have GCD exactly g
	ans := 0
	for g := 1; g <= maxVal; g++ {
		ans = (ans + gExact[g]) % MOD
	}
	return ans
}

func powMod(a, e int) int {
	res := 1
	for e > 0 {
		if e&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		e >>= 1
	}
	return res
}
```
