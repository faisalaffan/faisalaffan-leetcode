# 1998 — Gcd Sort Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NewDSU1998(n int) *DSU1998`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting, Union-Find

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1998: GCD Sort of an Array
// https://leetcode.com/problems/gcd-sort-of-an-array/
// Difficulty: Hard
// Approach: Union-Find over numbers and their prime factors.
// Two numbers can be swapped if they share a GCD > 1,
// which is equivalent to being connected through prime factors.
// After building DSU, check if each nums[i] and sorted[i] are in the same set.

import (
	"fmt"
	"sort"
)

type DSU1998 struct {
	parent []int
	rank   []int
}

func NewDSU1998(n int) *DSU1998 {
  // Alokasi slice
	p := make([]int, n)
  // Alokasi slice
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DSU1998{parent: p, rank: r}
}

func (d *DSU1998) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU1998) Union(x, y int) {
	xr, yr := d.Find(x), d.Find(y)
	if xr == yr {
		return
	}
	if d.rank[xr] < d.rank[yr] {
		xr, yr = yr, xr
	}
	d.parent[yr] = xr
	if d.rank[xr] == d.rank[yr] {
		d.rank[xr]++
	}
}

// smallestPrimeFactor using sieve
func spfSieve(limit int) []int {
  // Alokasi slice
	spf := make([]int, limit+1)
	for i := 2; i <= limit; i++ {
		if spf[i] == 0 {
			spf[i] = i
			if i*i <= limit {
				for j := i * i; j <= limit; j += i {
					if spf[j] == 0 {
						spf[j] = i
					}
				}
			}
		}
	}
	return spf
}

func gcdSort(nums []int) bool {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	// Sieve for smallest prime factors
	spf := spfSieve(maxVal + 1)
	// DSU for numbers 1..maxVal plus indices for nums
	// We union numbers with their prime factors
	dsu := NewDSU1998(maxVal + 1)

	// For each number, union with its prime factors
	for _, v := range nums {
		x := v
		for x > 1 {
			p := spf[x]
			if p == 0 {
				p = x
			}
			dsu.Union(v, p)
			for x%p == 0 {
				x /= p
			}
		}
	}

	// Sort a copy and check
  // Alokasi slice
	sorted := make([]int, len(nums))
	copy(sorted, nums)
  // Sort O(n log n)
	sort.Ints(sorted)

  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		if dsu.Find(nums[i]) != dsu.Find(sorted[i]) {
			return false
		}
	}
	return true
}

func main() {
	// Example: [7,21,3] -> true
	fmt.Println(gcdSort([]int{7, 21, 3}))

	// Additional tests
	fmt.Println(gcdSort([]int{5, 2, 6, 2}))
	fmt.Println(gcdSort([]int{1, 2, 3, 4}))
	fmt.Println(gcdSort([]int{10, 5, 9, 3, 15}))
}
```
