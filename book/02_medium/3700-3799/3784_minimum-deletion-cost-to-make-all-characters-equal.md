# 3784 — Minimum Deletion Cost To Make All Characters Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumDeletionCostToMakeAllCharactersEqual(s string, cost []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3784: Minimum Deletion Cost to Make All Characters Equal
// https://leetcode.com/problems/minimum-deletion-cost-to-make-all-characters-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumDeletionCostToMakeAllCharactersEqual(s string, cost []int) int64 {
	total := int64(0)
	charCost := [26]int64{}
	for i, ch := range s {
		v := int64(cost[i])
		total += v
		idx := ch - 'a'
		charCost[idx] += v
	}

	maxCost := int64(0)
	for _, v := range charCost {
		if v > maxCost {
			maxCost = v
		}
	}
	return total - maxCost
}

func main() {
	fmt.Println(minimumDeletionCostToMakeAllCharactersEqual("aabaac", []int{1, 2, 3, 4, 1, 10}))
	fmt.Println(minimumDeletionCostToMakeAllCharactersEqual("abc", []int{10, 5, 8}))
	fmt.Println(minimumDeletionCostToMakeAllCharactersEqual("zzzzz", []int{67, 67, 67, 67, 67}))
}
```
