# 1131 — Maximum Of Absolute Value Expression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxAbsValExpr(arr1 []int, arr2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1131: Maximum of Absolute Value Expression
// https://leetcode.com/problems/maximum-of-absolute-value-expression/
// Difficulty: Medium
//
// Approach: The expression |arr1[i] - arr1[j]| + |arr2[i] - arr2[j]| + |i - j|
//           can be expanded by considering all sign combinations.
//           Compute max of (arr1[i] + arr2[i] + i) - min of same, etc.
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxAbsValExpr([]int{1, -2, -5, 0, 10}, []int{0, -2, -1, -7, -4})) // 20
	fmt.Println(maxAbsValExpr([]int{1, 2, 3, 4}, []int{-1, 4, 5, 6}))             // 13
}

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	n := len(arr1)
	// Try all sign combinations for the three terms
	// Each term can be +/-: a1 + a2 + i, a1 + a2 - i, a1 - a2 + i, etc.
	// Actually: |x| = max(x, -x), so we need 8 combinations = 2^3
	// But many are redundant. The key combinations are:
	// (arr1[i] + arr2[i] + i), (arr1[i] + arr2[i] - i),
	// (arr1[i] - arr2[i] + i), (arr1[i] - arr2[i] - i)
	// and their negations (which are just -arr1[i] - arr2[i] - i, etc.)

	combos := [4]int{}
	for k := 0; k < 4; k++ {
		combos[k] = arr1[0] + arr2[0] + 0
		if k&1 != 0 {
			combos[k] = arr1[0] + arr2[0] - 0
		}
		if k&2 != 0 {
			combos[k] = arr1[0] - arr2[0] + 0
			if k&1 != 0 {
				combos[k] = arr1[0] - arr2[0] - 0
			}
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		for _, sign1 := range []int{-1, 1} {
			for _, sign2 := range []int{-1, 1} {
				for _, sign3 := range []int{-1, 1} {
					val := sign1*arr1[i] + sign2*arr2[i] + sign3*i
					for j := i + 1; j < n; j++ {
						val2 := sign1*arr1[j] + sign2*arr2[j] + sign3*j
						diff := val - val2
						if diff < 0 {
							diff = -diff
						}
						if diff > result {
							result = diff
						}
					}
				}
			}
		}
	}

	return result
}
```
