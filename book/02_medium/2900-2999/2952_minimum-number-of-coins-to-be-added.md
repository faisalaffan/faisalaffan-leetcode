# 2952 — Minimum Number Of Coins To Be Added

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumAddedCoins(coins []int, target int) (ans int)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2952: Minimum Number of Coins to be Added
// https://leetcode.com/problems/minimum-number-of-coins-to-be-added/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAddedCoins([]int{1, 4, 10}, 19))
	fmt.Println(minimumAddedCoins([]int{1, 4, 10, 5, 7, 19}, 19))
	fmt.Println(minimumAddedCoins([]int{1, 1, 1}, 20))
}

func minimumAddedCoins(coins []int, target int) (ans int) {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(coins)
	for i, s := 0, 1; s <= target; {
		if i < len(coins) && coins[i] <= s {
			s += coins[i]
			i++
		} else {
			s <<= 1
			ans++
		}
	}
	return
}
```
