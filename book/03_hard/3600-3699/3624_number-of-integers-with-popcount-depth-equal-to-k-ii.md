# 3624 — Number Of Integers With Popcount Depth Equal To K Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func popcountDepthII(nums []int64, queries [][]int64) []int
```

> **💡 Hint:** Precompute depth for all numbers, use Fenwick tree per depth level.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Fenwick Tree (BIT)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3624: Number of Integers With Popcount-Depth Equal to K II
// https://leetcode.com/problems/number-of-integers-with-popcount-depth-equal-to-k-ii/
// Difficulty: Hard
//
// Given array nums and queries, handle range queries to count numbers with
// popcount-depth == k and point updates.
//
// Approach: Precompute depth for all numbers, use Fenwick tree per depth level.

import "fmt"
import "math/bits"

func main() {
	// Example 1
	fmt.Println(popcountDepthII([]int64{1, 2, 3, 4, 5}, [][]int64{{1, 0, 4, 2}, {1, 0, 2, 1}}))
	// Example 2
	fmt.Println(popcountDepthII([]int64{7, 8, 9}, [][]int64{{1, 0, 2, 2}, {2, 1, 10}, {1, 0, 2, 2}}))
	// Edge: single element
	fmt.Println(popcountDepthII([]int64{1}, [][]int64{{1, 0, 0, 1}}))
}

func popcountDepthII(nums []int64, queries [][]int64) []int {
	n := len(nums)
	// Precompute depth for numbers up to 60 (max bits for 10^15)
	// Depth = number of popcount steps until reaching 1
  // Alokasi slice integer
	depth := make([]int, 61)
	depth[0] = 0
	depth[1] = 1
	for i := 2; i <= 60; i++ {
		d := 1
		x := i
		for x > 1 {
			x = bits.OnesCount(uint(x))
			d++
		}
		depth[i] = d
	}

	// Current depths for each element
  // Alokasi slice integer
	curDepth := make([]int, n)
	for i := 0; i < n; i++ {
		pop := bits.OnesCount64(uint64(nums[i]))
		if pop <= 60 {
			curDepth[i] = depth[pop]
		}
	}

	// Fenwick trees for each depth (0..5, since max depth is ~5 for 10^15)
	const MAX_DEPTH = 6
  // Membuat matriks/slice 2D untuk DP
	fenwick := make([][]int, MAX_DEPTH)
	for d := 0; d < MAX_DEPTH; d++ {
		fenwick[d] = make([]int, n+1)
	}
	add := func(tree []int, idx int, val int) {
		idx++
		for idx <= n {
			tree[idx] += val
			idx += idx & -idx
		}
	}
	sum := func(tree []int, idx int) int {
		s := 0
		idx++
		for idx > 0 {
			s += tree[idx]
			idx -= idx & -idx
		}
		return s
	}

	for i, d := range curDepth {
		if d < MAX_DEPTH {
			add(fenwick[d], i, 1)
		}
	}

  // Alokasi slice integer
	result := make([]int, 0)
	for _, q := range queries {
		if q[0] == 1 {
			l, r, k := int(q[1]), int(q[2]), int(q[3])
			if k < MAX_DEPTH {
				cnt := sum(fenwick[k], r) - sum(fenwick[k], l-1)
				result = append(result, cnt)
			} else {
				result = append(result, 0)
			}
		} else {
			idx, val := int(q[1]), int(q[2])
			oldD := curDepth[idx]
			pop := bits.OnesCount64(uint64(val))
			newD := 0
			if pop <= 60 {
				newD = depth[pop]
			}
			if oldD < MAX_DEPTH {
				add(fenwick[oldD], idx, -1)
			}
			curDepth[idx] = newD
			if newD < MAX_DEPTH {
				add(fenwick[newD], idx, 1)
			}
		}
	}
	return result
}
```
