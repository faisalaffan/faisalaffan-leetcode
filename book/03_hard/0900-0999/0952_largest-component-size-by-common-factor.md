# 0952 — Largest Component Size By Common Factor

## Deskripsi

**Soal:** [0952. Largest Component Size By Common Factor](https://leetcode.com/problems/largest-component-size-by-common-factor/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func largestComponentSize(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #952: Largest Component Size by Common Factor
// https://leetcode.com/problems/largest-component-size-by-common-factor/
// Difficulty: Hard

import "fmt"

func largestComponentSize(nums []int) int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return 0
	}

	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

  // Membuat slice untuk menyimpan hasil
	parent := make([]int, maxVal+1)
  // Membuat slice untuk menyimpan hasil
	size := make([]int, maxVal+1)
  // Iterasi seluruh elemen
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
	}

	// For each number, union it with its prime factors
	for _, num := range nums {
		x := num
		for p := 2; p*p <= x; p++ {
			if x%p == 0 {
				union(num, p)
				for x%p == 0 {
					x /= p
				}
			}
		}
		if x > 1 {
			union(num, x)
		}
	}

	// Count component sizes among the given numbers
  // Membuat map untuk pencarian O(1): key → value
	compCount := make(map[int]int)
	ans := 0
	for _, num := range nums {
		root := find(num)
		compCount[root]++
		if compCount[root] > ans {
			ans = compCount[root]
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println("Example 1:")
	fmt.Println(largestComponentSize([]int{4, 6, 15, 35}))
	// Expected: 4

	// Example 2
	fmt.Println("Example 2:")
	fmt.Println(largestComponentSize([]int{20, 50, 9, 63}))
	// Expected: 2

	// Example 3
	fmt.Println("Example 3:")
	fmt.Println(largestComponentSize([]int{2, 3, 6, 7, 4, 12, 21, 39}))
	// Expected: 8
}
```
