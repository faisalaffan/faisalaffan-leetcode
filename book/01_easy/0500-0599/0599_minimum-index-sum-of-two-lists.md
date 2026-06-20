# 0599 — Minimum Index Sum Of Two Lists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func MinimumIndexSumOfTwoLists(list1, list2 []string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n+m), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #599: Minimum Index Sum of Two Lists
// https://leetcode.com/problems/minimum-index-sum-of-two-lists/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

// Time: O(n+m), Space: O(n)
func MinimumIndexSumOfTwoLists(list1, list2 []string) []string {
  // HashMap: O(1) lookup
	index := make(map[string]int)
	for i, s := range list1 {
		index[s] = i
	}
	minSum := math.MaxInt32
	var result []string
	for j, s := range list2 {
		if i, ok := index[s]; ok {
			sum := i + j
			if sum < minSum {
				minSum = sum
				result = []string{s}
			} else if sum == minSum {
				result = append(result, s)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(MinimumIndexSumOfTwoLists(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
	))
	fmt.Println(MinimumIndexSumOfTwoLists(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"},
	))
}
```
