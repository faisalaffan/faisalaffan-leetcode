# 2722 — Join Two Arrays By Id

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func JoinTwoArraysById(arr1, arr2 []MapItem) []MapItem`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n + m)  |  **Ruang:** O(n + m)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2722: Join Two Arrays by ID
// https://leetcode.com/problems/join-two-arrays-by-id/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)

import (
	"fmt"
	"sort"
)

type MapItem struct {
	ID     int
	Values map[string]int
}

func JoinTwoArraysById(arr1, arr2 []MapItem) []MapItem {
  // HashMap: O(1) lookup
	merged := make(map[int]map[string]int)

	for _, item := range arr1 {
		if _, ok := merged[item.ID]; !ok {
			merged[item.ID] = make(map[string]int)
		}
		for k, v := range item.Values {
			merged[item.ID][k] = v
		}
	}
	for _, item := range arr2 {
		if _, ok := merged[item.ID]; !ok {
			merged[item.ID] = make(map[string]int)
		}
		for k, v := range item.Values {
			merged[item.ID][k] = v
		}
	}

	result := make([]MapItem, 0, len(merged))
	for id, vals := range merged {
		result = append(result, MapItem{ID: id, Values: vals})
	}
  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func main() {
	arr1 := []MapItem{{ID: 1, Values: map[string]int{"a": 1}}, {ID: 2, Values: map[string]int{"b": 2}}}
	arr2 := []MapItem{{ID: 1, Values: map[string]int{"c": 3}}, {ID: 3, Values: map[string]int{"d": 4}}}
	fmt.Println(JoinTwoArraysById(arr1, arr2))

	arr3 := []MapItem{{ID: 1, Values: map[string]int{"x": 10}}}
	arr4 := []MapItem{}
	fmt.Println(JoinTwoArraysById(arr3, arr4))
}
```
