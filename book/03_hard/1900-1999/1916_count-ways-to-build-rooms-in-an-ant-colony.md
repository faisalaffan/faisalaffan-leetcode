# 1916 — Count Ways To Build Rooms In An Ant Colony

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func waysToBuildRooms(prevRoom []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1916: Count Ways to Build Rooms in an Ant Colony
// https://leetcode.com/problems/count-ways-to-build-rooms-in-an-ant-colony/
// Difficulty: Hard

import "fmt"

const mod = 1_000_000_007

func waysToBuildRooms(prevRoom []int) int {
	n := len(prevRoom)
  // Matriks 2D
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := prevRoom[i]
		children[p] = append(children[p], i)
	}

	// Precompute factorials and inverse factorials
  // Alokasi slice
	fact := make([]int, n+1)
  // Alokasi slice
	invFact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i % mod
	}
	invFact[n] = modPow(fact[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * (i + 1) % mod
	}

	nCr := func(nn, r int) int {
		if r < 0 || r > nn {
			return 0
		}
		return fact[nn] * invFact[r] % mod * invFact[nn-r] % mod
	}

	// Post-order DP: returns (subtreeSize, ways) for each node
	var dfs func(u int) (int, int)
	dfs = func(u int) (int, int) {
		size := 1
		ways := 1
		for _, v := range children[u] {
			subSize, subWays := dfs(v)
			// Interleave this child's subtree with what we've accumulated so far.
			// C((size-1) + subSize, subSize) = ways to merge subSize elements
			// into sequence of (size-1) elements, preserving relative order,
			// since node u must be first.
			interleave := nCr(size-1+subSize, subSize)
			ways = ways * subWays % mod * interleave % mod
			size += subSize
		}
		return size, ways
	}

	_, ans := dfs(0)
	return ans
}

func modPow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

func main() {
	fmt.Println(waysToBuildRooms([]int{-1, 0, 1}))
	fmt.Println(waysToBuildRooms([]int{-1, 0, 0}))
}
```
