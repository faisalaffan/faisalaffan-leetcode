# 2838 — Maximum Coins Heroes Can Collect

## Deskripsi

**Soal:** [2838. Maximum Coins Heroes Can Collect](https://leetcode.com/problems/maximum-coins-heroes-can-collect/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n + m log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func MaximumCoinsHeroesCanCollect(heroes []int, monsters []int, coins []int) []int64`

## Solusi Go

```go
package main

// LeetCode #2838: Maximum Coins Heroes Can Collect
// https://leetcode.com/problems/maximum-coins-heroes-can-collect/
// Difficulty: Medium [Paid]
// Time: O(n log n + m log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func MaximumCoinsHeroesCanCollect(heroes []int, monsters []int, coins []int) []int64 {
	n := len(monsters)
	type monster struct {
		power int
		coin  int
	}
  // Membuat slice untuk menyimpan hasil
	monsterList := make([]monster, n)
	for i := 0; i < n; i++ {
		monsterList[i] = monster{monsters[i], coins[i]}
	}
	sort.Slice(monsterList, func(i, j int) bool {
		return monsterList[i].power < monsterList[j].power
	})

  // Membuat slice untuk menyimpan hasil
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(monsterList[i].coin)
	}

  // Membuat slice untuk menyimpan hasil
	heroSorted := make([]int, len(heroes))
	copy(heroSorted, heroes)
	sort.Ints(heroSorted)

  // Membuat map untuk pencarian O(1): key → value
	heroMap := make(map[int]int64)
	for _, h := range heroSorted {
		if _, ok := heroMap[h]; ok {
			continue
		}
		idx := sort.Search(n, func(i int) bool {
			return monsterList[i].power > h
		})
		heroMap[h] = prefix[idx]
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int64, len(heroes))
	for i, h := range heroes {
		result[i] = heroMap[h]
	}
	return result
}

func main() {
	fmt.Println(MaximumCoinsHeroesCanCollect([]int{1, 4, 2}, []int{1, 1, 3, 5}, []int{2, 3, 4, 5}))
	fmt.Println(MaximumCoinsHeroesCanCollect([]int{5}, []int{1, 2, 3}, []int{10, 20, 30}))
}
```
