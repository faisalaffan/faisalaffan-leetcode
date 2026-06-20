# 2980 — Check If Bitwise Or Has Trailing Zeros

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CheckIfBitwiseOrHasTrailingZeros(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2980: Check if Bitwise OR Has Trailing Zeros
// https://leetcode.com/problems/check-if-bitwise-or-has-trailing-zeros/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: hasTrailingZeros
	fmt.Println(CheckIfBitwiseOrHasTrailingZeros([]int{1, 2, 3, 4, 5})) // true
	fmt.Println(CheckIfBitwiseOrHasTrailingZeros([]int{2, 4, 8, 16}))   // true
	fmt.Println(CheckIfBitwiseOrHasTrailingZeros([]int{1, 3, 5, 7, 9})) // false
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: hasTrailingZeros
func CheckIfBitwiseOrHasTrailingZeros(nums []int) bool {
	evenCount := 0
	for _, num := range nums {
		if num%2 == 0 {
			evenCount++
			if evenCount >= 2 {
				return true
			}
		}
	}
	return false
}
```
