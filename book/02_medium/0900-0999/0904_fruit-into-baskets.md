# 0904 — Fruit Into Baskets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FruitIntoBaskets(fruits []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #904: Fruit Into Baskets
// https://leetcode.com/problems/fruit-into-baskets/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FruitIntoBaskets([]int{1, 2, 1}))
	fmt.Println(FruitIntoBaskets([]int{0, 1, 2, 2}))
	fmt.Println(FruitIntoBaskets([]int{1, 2, 3, 2, 2}))
}

// Time: O(n) | Space: O(1)
func FruitIntoBaskets(fruits []int) int {
  // HashMap: O(1) lookup
	cnt := make(map[int]int)
	left, ans := 0, 0

	for right, fruit := range fruits {
		cnt[fruit]++
		for len(cnt) > 2 {
			cnt[fruits[left]]--
			if cnt[fruits[left]] == 0 {
				delete(cnt, fruits[left])
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}

	return ans
}
```
