# 2439 — Minimize Maximum Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimizeArrayValue(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2439: Minimize Maximum of Array
// https://leetcode.com/problems/minimize-maximum-of-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Prefix average approach: we can distribute value to the left.
// The min possible max is the max prefix average (ceil).

import "fmt"

func main() {
	fmt.Println(minimizeArrayValue([]int{3, 7, 1, 6})) // 5
	fmt.Println(minimizeArrayValue([]int{10, 1}))      // 10
}

func minimizeArrayValue(nums []int) int {
	var sum int64
	ans := 0
	for i, v := range nums {
		sum += int64(v)
		avg := int((sum + int64(i)) / int64(i+1)) // ceil division
		if avg > ans {
			ans = avg
		}
	}
	return ans
}
```
