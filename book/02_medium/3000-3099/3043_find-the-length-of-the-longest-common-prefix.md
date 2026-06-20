# 3043 — Find The Length Of The Longest Common Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestCommonPrefix(arr1 []int, arr2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** O(n*logM + m*logM)  |  **Ruang:** O(n*logM)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3043: Find the Length of the Longest Common Prefix
// https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/
// Difficulty: Medium
// Time: O(n*logM + m*logM) | Space: O(n*logM)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
	fmt.Println(longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4}))
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixes := map[int]bool{}
	for _, x := range arr1 {
		for x > 0 {
			prefixes[x] = true
			x /= 10
		}
	}
	ans := 0
	for _, x := range arr2 {
		for x > 0 {
			if prefixes[x] {
				if len(strconv.Itoa(x)) > ans {
					ans = len(strconv.Itoa(x))
				}
				break
			}
			x /= 10
		}
	}
	return ans
}
```
