# 2712 — Minimum Cost To Make All Characters Equal

## Deskripsi

**Soal:** [2712. Minimum Cost To Make All Characters Equal](https://leetcode.com/problems/minimum-cost-to-make-all-characters-equal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MinimumCostToMakeAllCharactersEqual(s string) int64`

## Solusi Go

```go
package main

// LeetCode #2712: Minimum Cost to Make All Characters Equal
// https://leetcode.com/problems/minimum-cost-to-make-all-characters-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MinimumCostToMakeAllCharactersEqual(s string) int64 {
	n := len(s)

	calc := func(target byte) int64 {
		var cost int64
		flip := 0
		for i := 0; i < n; i++ {
			cur := s[i]
			if flip%2 == 1 {
				if cur == '0' {
					cur = '1'
				} else {
					cur = '0'
				}
			}
			if cur != target {
				flip++
				cost += int64(i + 1)
			}
		}
		return cost
	}

	cost0 := calc('0')
	cost1 := calc('1')
	if cost0 < cost1 {
		return cost0
	}
	return cost1
}

func main() {
	fmt.Println(MinimumCostToMakeAllCharactersEqual("0011"))
	fmt.Println(MinimumCostToMakeAllCharactersEqual("010101"))
}
```
