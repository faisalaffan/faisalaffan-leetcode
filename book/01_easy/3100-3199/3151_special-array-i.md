# 3151 — Special Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SpecialArrayI(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3151: Special Array I
// https://leetcode.com/problems/special-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isArraySpecial
	fmt.Println(SpecialArrayI([]int{1}))       // true
	fmt.Println(SpecialArrayI([]int{2, 1, 4})) // true
	fmt.Println(SpecialArrayI([]int{4, 3, 1, 6})) // false
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: isArraySpecial
func SpecialArrayI(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if (nums[i]%2 == 0) == (nums[i-1]%2 == 0) {
			return false
		}
	}
	return true
}
```
