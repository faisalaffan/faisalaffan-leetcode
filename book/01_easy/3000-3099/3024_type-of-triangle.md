# 3024 — Type Of Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func TypeOfTriangle(nums []int) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3024: Type of Triangle
// https://leetcode.com/problems/type-of-triangle/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: triangleType
	fmt.Println(TypeOfTriangle([]int{3, 3, 3})) // equilateral
	fmt.Println(TypeOfTriangle([]int{3, 4, 5})) // scalene
	fmt.Println(TypeOfTriangle([]int{3, 3, 5})) // isosceles
	fmt.Println(TypeOfTriangle([]int{1, 2, 3})) // none
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: triangleType
func TypeOfTriangle(nums []int) string {
  // Sort O(n log n)
	sort.Ints(nums)
	a, b, c := nums[0], nums[1], nums[2]

	// Check if valid triangle
	if a+b <= c {
		return "none"
	}

	if a == b && b == c {
		return "equilateral"
	}
	if a == b || b == c || a == c {
		return "isosceles"
	}
	return "scalene"
}
```
