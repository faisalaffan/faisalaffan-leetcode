# 2932 — Maximum Strong Pair Xor I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumStrongPairXorI(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2932: Maximum Strong Pair XOR I
// https://leetcode.com/problems/maximum-strong-pair-xor-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maximumStrongPairXor
	fmt.Println(MaximumStrongPairXorI([]int{1, 2, 3, 4, 5})) // 7
	fmt.Println(MaximumStrongPairXorI([]int{10, 100}))        // 0
	fmt.Println(MaximumStrongPairXorI([]int{5, 6, 25, 30}))   // 7
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: maximumStrongPairXor
func MaximumStrongPairXorI(nums []int) int {
	n := len(nums)
	maxXor := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			x, y := nums[i], nums[j]
			if abs(x-y) <= min(x, y) {
				if x^y > maxXor {
					maxXor = x ^ y
				}
			}
		}
	}
	return maxXor
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
