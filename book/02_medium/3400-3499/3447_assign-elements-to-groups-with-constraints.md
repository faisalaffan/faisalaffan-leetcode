# 3447 — Assign Elements To Groups With Constraints

## Deskripsi

**Soal:** [3447. Assign Elements To Groups With Constraints](https://leetcode.com/problems/assign-elements-to-groups-with-constraints/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(mx log mx + n + m) Space: O(mx)  
**Kompleksitas Ruang:** O(mx)

**Algoritma:** —

**Fungsi Solusi:** `func assignElements(groups []int, elements []int) []int`

## Solusi Go

```go
package main

// LeetCode #3447: Assign Elements to Groups with Constraints
// https://leetcode.com/problems/assign-elements-to-groups-with-constraints/
// Difficulty: Medium
// Time: O(mx log mx + n + m) Space: O(mx)

import (
	"fmt"
	"slices"
)

func assignElements(groups []int, elements []int) []int {
	mx := slices.Max(groups)
  // Membuat slice untuk menyimpan hasil
	target := make([]int, mx+1)
  // Iterasi seluruh elemen
	for i := range target {
		target[i] = -1
	}

	for i, x := range elements {
		if x > mx || target[x] >= 0 {
			continue
		}
		for y := x; y <= mx; y += x {
			if target[y] < 0 {
				target[y] = i
			}
		}
	}

	for i, x := range groups {
		groups[i] = target[x]
	}
	return groups
}

func main() {
	fmt.Println(assignElements([]int{8, 4, 3, 2, 4}, []int{4, 2})) // [0 0 -1 1 0]
	fmt.Println(assignElements([]int{10, 5, 7}, []int{2, 5})) // [1 1 -1]
}
```
