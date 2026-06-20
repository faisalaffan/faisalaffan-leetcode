# 3780 — Maximum Sum Of Three Numbers Divisible By Three

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumSumOfThreeNumbersDivisibleByThree(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3780: Maximum Sum of Three Numbers Divisible by Three
// https://leetcode.com/problems/maximum-sum-of-three-numbers-divisible-by-three/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumSumOfThreeNumbersDivisibleByThree(nums []int) int {
	groups := [3][]int{}
	for _, v := range nums {
		r := v % 3
		groups[r] = append(groups[r], v)
	}

	for r := 0; r < 3; r++ {
  // Custom sort dengan comparator
		sort.Slice(groups[r], func(i, j int) bool {
			return groups[r][i] > groups[r][j]
		})
	}

	ans := 0

	// (0,0,0)
	if len(groups[0]) >= 3 {
		sum := groups[0][0] + groups[0][1] + groups[0][2]
		if sum > ans {
			ans = sum
		}
	}

	// (1,1,1)
	if len(groups[1]) >= 3 {
		sum := groups[1][0] + groups[1][1] + groups[1][2]
		if sum > ans {
			ans = sum
		}
	}

	// (2,2,2)
	if len(groups[2]) >= 3 {
		sum := groups[2][0] + groups[2][1] + groups[2][2]
		if sum > ans {
			ans = sum
		}
	}

	// (0,1,2)
	if len(groups[0]) >= 1 && len(groups[1]) >= 1 && len(groups[2]) >= 1 {
		sum := groups[0][0] + groups[1][0] + groups[2][0]
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

func main() {
	fmt.Println(maximumSumOfThreeNumbersDivisibleByThree([]int{4, 2, 3, 1}))
	fmt.Println(maximumSumOfThreeNumbersDivisibleByThree([]int{1, 2, 3, 4, 5}))
	fmt.Println(maximumSumOfThreeNumbersDivisibleByThree([]int{1, 1, 1}))
}
```
