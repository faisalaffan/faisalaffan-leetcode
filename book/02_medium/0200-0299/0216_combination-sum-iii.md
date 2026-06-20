# 0216 — Combination Sum Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func combinationSum3(k int, n int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(C(9,k)), Space: O(k)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #216: Combination Sum III
// https://leetcode.com/problems/combination-sum-iii/
// Difficulty: Medium
// Time: O(C(9,k)), Space: O(k)

import "fmt"

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	var backtrack func(start, remaining int, combo []int)
	backtrack = func(start, remaining int, combo []int) {
		if len(combo) == k && remaining == 0 {
  // Alokasi slice integer
			comboCopy := make([]int, len(combo))
			copy(comboCopy, combo)
			result = append(result, comboCopy)
			return
		}
		if len(combo) > k || remaining < 0 {
			return
		}

		for i := start; i <= 9; i++ {
			combo = append(combo, i)
			backtrack(i+1, remaining-i, combo)
			combo = combo[:len(combo)-1]
		}
	}

	backtrack(1, n, []int{})
	return result
}

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
}
```
