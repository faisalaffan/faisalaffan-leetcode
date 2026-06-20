# 2491 — Divide Players Into Teams Of Equal Skill

## Deskripsi

**Soal:** [2491. Divide Players Into Teams Of Equal Skill](https://leetcode.com/problems/divide-players-into-teams-of-equal-skill/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2491: Divide Players Into Teams of Equal Skill
// https://leetcode.com/problems/divide-players-into-teams-of-equal-skill/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)
// Sort, pair smallest with largest. Each pair sum must equal same target.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(dividePlayers([]int{3, 2, 5, 1, 3, 4})) // 22
	fmt.Println(dividePlayers([]int{1, 1, 2, 3}))       // -1
}

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	n := len(skill)
	target := skill[0] + skill[n-1]
	var sum int64
	for i := 0; i < n/2; i++ {
		if skill[i]+skill[n-1-i] != target {
			return -1
		}
		sum += int64(skill[i]) * int64(skill[n-1-i])
	}
	return sum
}
```
