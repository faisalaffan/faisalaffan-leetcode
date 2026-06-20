# 3428 — Maximum And Minimum Sums Of At Most Size K Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func init() `

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log n) Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3428: Maximum and Minimum Sums of at Most Size K Subsequences
// https://leetcode.com/problems/maximum-and-minimum-sums-of-at-most-size-k-subsequences/
// Difficulty: Medium
// Time: O(n log n) Space: O(1)

import (
	"fmt"
	"slices"
)

const mod3428 = 1_000_000_007
const mx = 100000

var fac [mx]int
var invFac [mx]int

func init() {
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod3428
	}
	invFac[mx-1] = pow3428(fac[mx-1], mod3428-2)
	for i := mx - 1; i > 0; i-- {
		invFac[i-1] = invFac[i] * i % mod3428
	}
}

func pow3428(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod3428
		}
		x = x * x % mod3428
	}
	return res
}

func comb(n, k int) int {
	if k > n || k < 0 {
		return 0
	}
	return fac[n] * invFac[k] % mod3428 * invFac[n-k] % mod3428
}

func minMaxSums(nums []int, k int) int {
	slices.Sort(nums)
	ans := 0
	s := 1
	n := len(nums)
	for i, x := range nums {
		ans = (ans + s*(x+nums[n-1-i])) % mod3428
		s = (s*2 - comb(i, k-1) + mod3428) % mod3428
	}
	return ans
}

func main() {
	fmt.Println(minMaxSums([]int{5, 0, 6}, 1)) // 22
	fmt.Println(minMaxSums([]int{1, 2, 3}, 2)) // 24
}
```
