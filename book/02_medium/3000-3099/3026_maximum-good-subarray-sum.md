# 3026 — Maximum Good Subarray Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumSum(nums []int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3026: Maximum Good Subarray Sum
// https://leetcode.com/problems/maximum-good-subarray-sum/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maximumSum([]int{1, 2, 3, 4, 5, 6}, 1))
	fmt.Println(maximumSum([]int{-1, 3, 2, 4, 5}, 3))
	fmt.Println(maximumSum([]int{-1, -2, -3, -4}, 2))
}

func maximumSum(nums []int, k int) int64 {
	const negInf int64 = -(1 << 60)
	ans := negInf
	pref := int64(0)
	first := map[int]int64{}
	for _, x := range nums {
		if v, ok := first[x]; ok {
			if pref+int64(x)-v > ans {
				ans = pref + int64(x) - v
			}
			if pref < v {
				first[x] = pref
			}
		} else {
			first[x] = pref
		}
		pref += int64(x)
		need := x - k
		if v, ok := first[need]; ok {
			if pref-v > ans {
				ans = pref - v
			}
		}
		need2 := x + k
		if v, ok := first[need2]; ok {
			if pref-v > ans {
				ans = pref - v
			}
		}
	}
	if ans == negInf {
		return 0
	}
	return ans
}
```
