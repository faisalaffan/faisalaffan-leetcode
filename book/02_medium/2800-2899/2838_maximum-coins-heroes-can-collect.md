# 2838 — Maximum Coins Heroes Can Collect

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumCoinsHeroesCanCollect(heroes []int, monsters []int, coins []int) []int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting, Prefix Sum

**Waktu:** O(n log n + m log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
	monsterList := make([]monster, n)
	for i := 0; i < n; i++ {
		monsterList[i] = monster{monsters[i], coins[i]}
	}
  // Custom sort
	sort.Slice(monsterList, func(i, j int) bool {
		return monsterList[i].power < monsterList[j].power
	})

  // Alokasi slice
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(monsterList[i].coin)
	}

  // Alokasi slice
	heroSorted := make([]int, len(heroes))
	copy(heroSorted, heroes)
  // Sort O(n log n)
	sort.Ints(heroSorted)

  // HashMap: O(1) lookup
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

  // Alokasi slice
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
