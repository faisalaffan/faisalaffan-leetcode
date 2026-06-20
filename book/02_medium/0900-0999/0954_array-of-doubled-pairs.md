# 0954 — Array Of Doubled Pairs

## Deskripsi

**Soal:** [0954. Array Of Doubled Pairs](https://leetcode.com/problems/array-of-doubled-pairs/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func canReorderDoubled(arr []int) bool`

## Solusi Go

```go
package main

// LeetCode #954: Array of Doubled Pairs
// https://leetcode.com/problems/array-of-doubled-pairs/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(n)
func canReorderDoubled(arr []int) bool {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

  // Membuat slice untuk menyimpan hasil
	keys := make([]int, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return abs(keys[i]) < abs(keys[j])
	})

	for _, v := range keys {
		if freq[v] > freq[2*v] {
			return false
		}
		freq[2*v] -= freq[v]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(canReorderDoubled([]int{3, 1, 3, 6}))
	fmt.Println(canReorderDoubled([]int{2, 1, 2, 6}))
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{1, 2, 4, 16, 8, 4}))
}
```
